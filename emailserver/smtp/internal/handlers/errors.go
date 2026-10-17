package handlers

import "fmt"

type SMTPError struct {
	Code int
	Msg  string
}

func (e SMTPError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

// --- Common SMTP Errors ---

func ErrSyntax(msg string) error {
	return SMTPError{Code: 501, Msg: msg}
}

func ErrBadSequence(msg string) error {
	return SMTPError{Code: 503, Msg: msg}
}

func ErrPolicy(msg string) error {
	return SMTPError{Code: 550, Msg: msg}
}

func ErrTemp(msg string) error {
	return SMTPError{Code: 451, Msg: msg}
}

func ErrInternal(msg string) error {
	return SMTPError{Code: 451, Msg: "Internal error: " + msg}
}

// Convert error → Response
func ErrorToResponse(err error) Response {
	if smtpErr, ok := err.(SMTPError); ok {
		return Response{Text: smtpErr.Error()}
	}
	return Response{Text: "451 Internal server error"}
}
