package handlers

import (
	"net/http"
)

type ErrorData struct {
	StatusCode int
	Message    string
}

func ErrorHandler(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	data := ErrorData{
		StatusCode: status,
		Message:    message,
	}
	RenderTemplate(w, "error.html", data)
}