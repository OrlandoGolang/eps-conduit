package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// customWriter is a thin wrapper around http.ResponseWriter which captures
// the status code and content length in a way that a normal writer cannot
type customWriter struct {
	http.ResponseWriter
	status int
	length int
}

// StatusCode returns the response's status code
func (w *customWriter) StatusCode() int {
	return w.status
}

// ContentLength returns the response's body length
func (w *customWriter) ContentLength() int {
	return w.length
}

// WriteHeader satisfies the http.ResponseWriter interface
func (w *customWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Write satisfies the http.ResponseWriter interface
func (w *customWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

// LoggingMiddleware intercepts a handler writing the relevant information of
// the request to the access log file
func LoggingMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	file, err := os.OpenFile(config.AccessLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}
	logger := log.New(file, "", 0)

	return func(w http.ResponseWriter, r *http.Request) {
		writer := customWriter{w, 0, 0}
		start := time.Now()
		fn(&writer, r)
		end := time.Now()

		logger.Printf(
			"%s %d %s %s %s %d %d %s %d \"%s %s%s %s\" \"%s\"",
			end.Format(time.RFC3339),
			config.NextHost,
			config.Backends[config.NextHost],
			r.RemoteAddr,
			end.Sub(start),
			r.ContentLength,
			writer.ContentLength(),
			r.Proto,
			writer.StatusCode(),
			r.Method,
			r.Host,
			r.RequestURI,
			r.Proto,
			r.UserAgent(),
		)
	}
}
