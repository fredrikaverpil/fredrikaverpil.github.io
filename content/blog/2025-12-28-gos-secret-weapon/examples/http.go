package examples

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// EnforceJSON is middleware that requires JSON content type using http.Handler.
func EnforceJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// LogStatus is middleware that logs the response status using statusRecorder.
func LogStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)
		slog.Info("request", "path", r.URL.Path, "status", rec.status)
	})
}

// StreamEvents sends server-sent events using http.Flusher interface.
func StreamEvents(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, "data: Event %d\n\n", i)
		flusher.Flush()
		time.Sleep(time.Second)
	}
}

// UpgradeConnection handles protocol upgrades using http.Hijacker.
func UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	conn.Write([]byte("HTTP/1.1 101 Switching Protocols\r\n\r\n"))
}

// HandleConn handles any network connection using net.Conn interface.
func HandleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Hello from %v\n", conn.LocalAddr())
}

// RunServer accepts any listener implementing net.Listener interface.
func RunServer(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go HandleConn(conn)
	}
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// statusRecorder wraps http.ResponseWriter to capture status code.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

// LoggingTransport implements http.RoundTripper to log outbound HTTP requests.
type LoggingTransport struct {
	Next http.RoundTripper
}

func (l *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("Sending request to:", req.URL)
	return l.Next.RoundTrip(req)
}
