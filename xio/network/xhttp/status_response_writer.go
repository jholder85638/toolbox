package xhttp

import "net/http"

// StatusResponseWriter wraps an http.ResponseWriter and provides methods to
// retrieve the status code and number of bytes written.
type StatusResponseWriter struct {
	Original http.ResponseWriter
	status   int
	written  int
}

// Status returns the status that was set, or http.StatusOK if no call to
// WriteHeader() was made.
func (w *StatusResponseWriter) Status() int {
	if w.status != 0 {
		return w.status
	}
	return http.StatusOK
}

// BytesWritten returns the number of bytes written.
func (w *StatusResponseWriter) BytesWritten() int {
	return w.written
}

// Header implements http.ResponseWriter.
func (w *StatusResponseWriter) Header() http.Header {
	return w.Original.Header()
}

// Write implements http.ResponseWriter.
func (w *StatusResponseWriter) Write(data []byte) (int, error) {
	n, err := w.Original.Write(data)
	w.written += n
	return n, err
}

// WriteHeader implements http.ResponseWriter.
func (w *StatusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.Original.WriteHeader(status)
}

// Flush implements http.Flusher.
func (w *StatusResponseWriter) Flush() {
	f, ok := w.Original.(http.Flusher)
	if ok {
		f.Flush()
	}
}