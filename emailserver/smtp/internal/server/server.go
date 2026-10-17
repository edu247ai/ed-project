package server

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"email-server/smtp/internal/handlers"
	"email-server/smtp/pkg/config"
	"email-server/smtp/pkg/logger"
)

type Server struct {
	cfg  config.SMTPConfig
	deps handlers.Dependencies
}

func New(cfg config.SMTPConfig) *Server {
	return NewWithDependencies(cfg, handlers.Dependencies{})
}

func NewWithDependencies(cfg config.SMTPConfig, deps handlers.Dependencies) *Server {
	if cfg.MaxConnections <= 0 {
		cfg.MaxConnections = 500
	}
	if cfg.MaxMessageSize <= 0 {
		cfg.MaxMessageSize = 20 << 20
	}
	if cfg.ReadTimeout <= 0 {
		cfg.ReadTimeout = 5 * time.Minute
	}
	if cfg.WriteTimeout <= 0 {
		cfg.WriteTimeout = 30 * time.Second
	}
	if cfg.Hostname == "" {
		cfg.Hostname = "localhost"
	}
	return &Server{cfg: cfg, deps: deps}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.cfg.ListenAddress)
	if err != nil {
		return err
	}
	logger.Info("SMTP server listening on %s", listener.Addr())
	return s.Serve(ctx, listener)
}

func (s *Server) Serve(ctx context.Context, listener net.Listener) error {
	var sessions sync.WaitGroup
	var activeMu sync.Mutex
	active := make(map[net.Conn]struct{})
	semaphore := make(chan struct{}, s.cfg.MaxConnections)
	stopListener := make(chan struct{})
	go func() {
		select {
		case <-ctx.Done():
			_ = listener.Close()
			activeMu.Lock()
			for conn := range active {
				_ = conn.Close()
			}
			activeMu.Unlock()
		case <-stopListener:
		}
	}()
	defer close(stopListener)
	defer func() {
		activeMu.Lock()
		for conn := range active {
			_ = conn.Close()
		}
		activeMu.Unlock()
		sessions.Wait()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil || errors.Is(err, net.ErrClosed) {
				return nil
			}
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				logger.Error("temporary SMTP accept error: %v", err)
				continue
			}
			return err
		}
		if ctx.Err() != nil {
			_ = conn.Close()
			continue
		}
		select {
		case semaphore <- struct{}{}:
			activeMu.Lock()
			active[conn] = struct{}{}
			activeMu.Unlock()
			sessions.Add(1)
			go func() {
				defer sessions.Done()
				defer func() { <-semaphore }()
				defer func() {
					activeMu.Lock()
					delete(active, conn)
					activeMu.Unlock()
				}()
				defer conn.Close()
				if err := handlers.NewSession(conn, handlers.ServerConfig{
					Addr:            s.cfg.ListenAddress,
					ReadTimeout:     s.cfg.ReadTimeout,
					WriteTimeout:    s.cfg.WriteTimeout,
					Hostname:        s.cfg.Hostname,
					MaxMessageSize:  s.cfg.MaxMessageSize,
					EnableSPF:       s.cfg.EnableSPF,
					EnableDKIM:      s.cfg.EnableDKIM,
					EnableDMARC:     s.cfg.EnableDMARC,
					EnableSpam:      s.cfg.EnableSpam,
					SpamRejectScore: 8,
				}, s.deps).Run(); err != nil {
					logger.Debug("SMTP session ended: %v", err)
				}
			}()
		default:
			_, _ = conn.Write([]byte("421 Too many connections, try again later\r\n"))
			_ = conn.Close()
		}
	}
}
