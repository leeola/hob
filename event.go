package hob

import "io"

type EventRequest struct {
	Body io.ReadCloser
}

type EventResponseWriter interface {
	io.Writer
}

type eventResponseWriter struct {
	Body io.ReadCloser
}
