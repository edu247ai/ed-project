package handlers

import (
	"net/mail"
	"strings"
)

type HandlerFunc func(cmd Command, st *State) Response

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	r := &Router{handlers: make(map[string]HandlerFunc)}
	r.registerDefaults()
	return r
}

func (r *Router) Handle(cmd Command, st *State) Response {
	if h, ok := r.handlers[cmd.Name]; ok {
		return h(cmd, st)
	}
	return Error(500, "Command unrecognized")
}

func (r *Router) registerDefaults() {
	r.handlers["HELO"] = handleHelo
	r.handlers["EHLO"] = handleEhlo
	r.handlers["MAIL"] = handleMailFrom
	r.handlers["RCPT"] = handleRcptTo
	r.handlers["DATA"] = handleData
	r.handlers["RSET"] = handleRset
	r.handlers["QUIT"] = handleQuit
	r.handlers["NOOP"] = func(Command, *State) Response { return Ok("OK") }
	r.handlers["VRFY"] = func(Command, *State) Response { return Error(252, "Cannot verify user") }
}

func handleHelo(cmd Command, st *State) Response {
	if strings.TrimSpace(cmd.Arg) == "" {
		return Error(501, "HELO requires a hostname")
	}
	st.HeloName = cmd.Arg
	resetTransaction(st)
	return Ok("Hello " + cmd.Arg)
}

func handleEhlo(cmd Command, st *State) Response {
	if strings.TrimSpace(cmd.Arg) == "" {
		return Error(501, "EHLO requires a hostname")
	}
	st.HeloName = cmd.Arg
	resetTransaction(st)
	return Response{
		Text: "250-" + cmd.Arg + "\r\n" +
			"250-PIPELINING\r\n" +
			"250-8BITMIME\r\n" +
			"250 SIZE",
	}
}

func handleMailFrom(cmd Command, st *State) Response {
	if st.HeloName == "" {
		return Error(503, "Send HELO or EHLO first")
	}
	if !strings.HasPrefix(strings.ToUpper(cmd.Arg), "FROM:") {
		return Error(501, "Syntax: MAIL FROM:<address>")
	}
	addr, ok := parsePath(strings.TrimSpace(cmd.Arg[5:]))
	if !ok {
		return Error(501, "Invalid sender address")
	}
	st.MailFrom = addr
	st.NullSender = addr == ""
	st.RcptTo = nil
	return Ok("Sender OK")
}

func handleRcptTo(cmd Command, st *State) Response {
	if st.MailFrom == "" && !st.NullSender {
		return Error(503, "MAIL FROM required before RCPT TO")
	}
	if !strings.HasPrefix(strings.ToUpper(cmd.Arg), "TO:") {
		return Error(501, "Syntax: RCPT TO:<address>")
	}
	addr, ok := parsePath(strings.TrimSpace(cmd.Arg[3:]))
	if !ok || addr == "" {
		return Error(501, "Invalid recipient address")
	}
	st.RcptTo = append(st.RcptTo, addr)
	return Ok("Recipient OK")
}

func handleData(cmd Command, st *State) Response {
	if st.MailFrom == "" && !st.NullSender {
		return Error(503, "MAIL FROM required before DATA")
	}
	if len(st.RcptTo) == 0 {
		return Error(503, "RCPT TO required before DATA")
	}
	return Response{Text: "354 End data with <CR><LF>.<CR><LF>"}
}

func handleRset(cmd Command, st *State) Response {
	helo := st.HeloName
	*st = State{HeloName: helo}
	return Ok("State reset")
}

func handleQuit(cmd Command, st *State) Response {
	return Close("221 Bye")
}

func resetTransaction(st *State) {
	st.MailFrom = ""
	st.NullSender = false
	st.RcptTo = nil
}

func parsePath(value string) (string, bool) {
	if index := strings.IndexAny(value, " \t"); index >= 0 {
		value = value[:index]
	}
	if !strings.HasPrefix(value, "<") || !strings.HasSuffix(value, ">") {
		return "", false
	}
	address := value[1 : len(value)-1]
	if address == "" {
		return "", true
	}
	parsed, err := mail.ParseAddress(address)
	if err != nil || parsed.Address != address {
		return "", false
	}
	return address, true
}
