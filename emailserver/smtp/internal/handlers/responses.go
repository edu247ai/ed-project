package handlers

import "fmt"

type Response struct {
	Text  string
	Close bool
}

func Ok(text string) Response {
	return Response{Text: "250 " + text}
}

func Error(code int, text string) Response {
	return Response{Text: fmt.Sprintf("%d %s", code, text)}
}

func Close(text string) Response {
	return Response{Text: text, Close: true}
}
