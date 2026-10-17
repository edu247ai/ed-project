package handlers

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/mail"
	"strings"
	"time"

	"email-server/smtp/pkg/logger"
)

type Session struct {
	conn   net.Conn
	cfg    ServerConfig
	reader *bufio.Reader
	writer *bufio.Writer
	router *Router
	state  *State
	deps   Dependencies
}

type State struct {
	HeloName   string
	MailFrom   string
	RcptTo     []string
	NullSender bool
}

func NewSession(conn net.Conn, cfg ServerConfig, deps Dependencies) *Session {
	return &Session{
		conn:   conn,
		cfg:    cfg,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		router: NewRouter(),
		state:  &State{},
		deps:   deps,
	}
}

func (s *Session) Run() error {
	if err := s.writeLine("220 " + s.cfg.Hostname + " ESMTP ready"); err != nil {
		return err
	}
	for {
		s.conn.SetReadDeadline(time.Now().Add(s.cfg.ReadTimeout))
		line, err := s.readLine(512)
		if err != nil {
			if errors.Is(err, errLineTooLong) {
				_ = s.writeLine("500 Command line too long")
			}
			return err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			continue
		}
		cmd := ParseCommand(line)
		if cmd.Name == "DATA" {
			resp := handleData(cmd, s.state)
			if err := s.writeLine(resp.Text); err != nil {
				return err
			}
			if !strings.HasPrefix(resp.Text, "354 ") {
				continue
			}
			raw, err := s.readData()
			if err != nil {
				if errors.Is(err, errMessageTooLarge) {
					_ = s.writeLine("552 Message size exceeds fixed maximum message size")
					return nil
				}
				return err
			}
			response := s.processData(context.Background(), raw)
			if err := s.writeLine(response.Text); err != nil {
				return err
			}
			resetTransaction(s.state)
			continue
		}
		resp := s.router.Handle(cmd, s.state)
		if cmd.Name == "EHLO" && s.cfg.MaxMessageSize > 0 {
			resp.Text = strings.TrimSuffix(resp.Text, "250 SIZE") + fmt.Sprintf("250 SIZE %d", s.cfg.MaxMessageSize)
		}
		if resp.Text != "" {
			if err := s.writeLine(resp.Text); err != nil {
				return err
			}
		}
		if resp.Close {
			return nil
		}
	}
}

var (
	errLineTooLong     = errors.New("SMTP line exceeds configured limit")
	errMessageTooLarge = errors.New("SMTP message exceeds configured limit")
)

func (s *Session) writeLine(line string) error {
	s.conn.SetWriteDeadline(time.Now().Add(s.cfg.WriteTimeout))
	if _, err := io.WriteString(s.writer, line); err != nil {
		return err
	}
	if !strings.HasSuffix(line, "\r\n") {
		if _, err := io.WriteString(s.writer, "\r\n"); err != nil {
			return err
		}
	}
	return s.writer.Flush()
}

func (s *Session) readLine(limit int) (string, error) {
	var line strings.Builder
	for {
		part, err := s.reader.ReadSlice('\n')
		if line.Len()+len(part) > limit {
			return "", errLineTooLong
		}
		line.Write(part)
		if err == nil {
			return line.String(), nil
		}
		if !errors.Is(err, bufio.ErrBufferFull) {
			return "", err
		}
	}
}

func (s *Session) readData() ([]byte, error) {
	var message []byte
	for {
		s.conn.SetReadDeadline(time.Now().Add(s.cfg.ReadTimeout))
		line, err := s.readLine(s.cfg.MaxMessageSize + 3)
		if errors.Is(err, errLineTooLong) {
			return nil, errMessageTooLarge
		}
		if err != nil {
			return nil, err
		}
		if line == ".\r\n" || line == ".\n" {
			return message, nil
		}
		if strings.HasPrefix(line, "..") {
			line = line[1:]
		}
		if len(message)+len(line) > s.cfg.MaxMessageSize {
			return nil, errMessageTooLarge
		}
		message = append(message, line...)
	}
}

func (s *Session) ProcessData(raw []byte) Response {
	return s.processData(context.Background(), raw)
}

func (s *Session) processData(ctx context.Context, raw []byte) Response {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if s.deps.Storage == nil {
		return Error(451, "Message storage unavailable")
	}
	domain := ""
	if at := strings.LastIndexByte(s.state.MailFrom, '@'); at >= 0 {
		domain = s.state.MailFrom[at+1:]
	}
	spfResult := "none"
	spfDomain := domain
	if s.cfg.EnableSPF && s.deps.SPF != nil && domain != "" {
		result, err := s.deps.SPF(ctx, s.remoteIP(), domain)
		if err != nil {
			return Error(451, "Temporary SPF evaluation failure")
		}
		spfResult = result
	}
	dkimResult, dkimDomain := "none", ""
	if s.cfg.EnableDKIM && s.deps.DKIM != nil {
		verifiedDomain, verified, err := s.deps.DKIM(ctx, raw)
		if err != nil {
			return Error(451, "Temporary DKIM evaluation failure")
		}
		if verified {
			dkimResult, dkimDomain = "pass", verifiedDomain
		} else if hasDKIMSignature(raw) {
			dkimResult = "fail"
		}
	}
	fromDomain := messageFromDomain(raw)
	if s.cfg.EnableDMARC && s.deps.DMARC != nil && fromDomain != "" {
		result, err := s.deps.DMARC(ctx, spfResult, spfDomain, dkimResult, dkimDomain, fromDomain)
		if err != nil {
			return Error(451, "Temporary DMARC evaluation failure")
		}
		if result == "reject" {
			return Error(550, "DMARC policy rejected message")
		}
	}
	if s.cfg.EnableSpam && s.deps.Spam != nil {
		score, flags, err := s.deps.Spam.Check(ctx, raw)
		if err != nil {
			return Error(451, "Temporary spam check failure")
		}
		logger.Debug("Rspamd score %.2f, symbols: %v", score, flags)
		if score > s.cfg.SpamRejectScore {
			return Error(550, "Message rejected as spam")
		}
	}

	msgID, err := s.deps.Storage.SaveIncomingMessage(ctx, s.state.HeloName, s.state.MailFrom, s.state.RcptTo, raw)
	if err != nil {
		return Error(451, "Temporary storage failure")
	}
	resetTransaction(s.state)
	return Ok("Message accepted, ID=" + msgID)
}

func messageFromDomain(raw []byte) string {
	message, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		return ""
	}
	from, err := mail.ParseAddress(message.Header.Get("From"))
	if err != nil {
		return ""
	}
	at := strings.LastIndexByte(from.Address, '@')
	if at < 0 {
		return ""
	}
	return strings.ToLower(strings.TrimSuffix(from.Address[at+1:], "."))
}

func hasDKIMSignature(raw []byte) bool {
	message, err := mail.ReadMessage(bytes.NewReader(raw))
	if err != nil {
		return false
	}
	return len(message.Header["Dkim-Signature"]) > 0
}

func (s *Session) remoteIP() string {
	host, _, err := net.SplitHostPort(s.conn.RemoteAddr().String())
	if err != nil {
		return ""
	}
	return host
}
