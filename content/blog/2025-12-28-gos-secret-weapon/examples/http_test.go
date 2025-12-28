package examples

import (
	"bytes"
	"fmt"
	"net/http"
)

// ExampleEnforceJSON demonstrates http.Handler middleware
func ExampleEnforceJSON() {
	handler := EnforceJSON(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Success")
	}))

	req, _ := http.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	recorder := &mockResponseWriter{header: make(http.Header)}
	handler.ServeHTTP(recorder, req)

	fmt.Println("Status Code:", recorder.statusCode)
	// Output: Status Code: 200
}

// ExampleLoggingTransport demonstrates http.RoundTripper implementation
func ExampleLoggingTransport() {
	transport := &LoggingTransport{Next: http.DefaultTransport}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	transport.RoundTrip(req)
	// Output: Sending request to: http://example.com
}

// mockResponseWriter is a helper for testing http.ResponseWriter interactions.
type mockResponseWriter struct {
	statusCode int
	header     http.Header
	body       bytes.Buffer
}

func (m *mockResponseWriter) Header() http.Header { return m.header }
func (m *mockResponseWriter) Write(b []byte) (int, error) {
	if m.statusCode == 0 {
		m.statusCode = http.StatusOK
	}
	return m.body.Write(b)
}
func (m *mockResponseWriter) WriteHeader(code int) { m.statusCode = code }

