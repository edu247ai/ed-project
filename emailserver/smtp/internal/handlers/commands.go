package handlers

import (
	"strings"
)

type CommandType int

const (
	CmdUnknown CommandType = iota
	CmdHELO
	CmdEHLO
	CmdMAIL
	CmdRCPT
	CmdDATA
	CmdRSET
	CmdQUIT
	CmdNOOP
	CmdVRFY
)

type Command struct {
	Type CommandType
	Name string
	Arg  string
	Raw  string
}

func ParseCommand(line string) Command {
	raw := strings.TrimSpace(line)
	if raw == "" {
		return Command{Type: CmdUnknown, Raw: raw}
	}

	arg := ""
	name, arg, found := strings.Cut(raw, " ")
	if !found {
		name, arg, found = strings.Cut(raw, "\t")
	}

	return Command{
		Type: mapNameToType(name),
		Name: strings.ToUpper(name),
		Arg:  strings.TrimSpace(arg),
		Raw:  raw,
	}
}

func mapNameToType(name string) CommandType {
	switch name {
	case "HELO":
		return CmdHELO
	case "EHLO":
		return CmdEHLO
	case "MAIL":
		return CmdMAIL
	case "RCPT":
		return CmdRCPT
	case "DATA":
		return CmdDATA
	case "RSET":
		return CmdRSET
	case "QUIT":
		return CmdQUIT
	case "NOOP":
		return CmdNOOP
	case "VRFY":
		return CmdVRFY
	default:
		return CmdUnknown
	}
}
