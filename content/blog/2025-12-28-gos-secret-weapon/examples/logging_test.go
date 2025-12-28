package examples

import (
	"fmt"
	"log/slog"
)

// ExampleCustomHandler demonstrates implementing slog.Handler.
func ExampleCustomHandler() {
	h := &CustomHandler{}
	logger := slog.New(h)

	// Add context with WithAttrs
	logger = logger.With("env", "prod")

	logger.Info("Application started", "version", "1.0")
	// Output: [INFO] Application started env=prod version=1.0
}

// ExampleToken demonstrates implementing slog.LogValuer for redaction.
func ExampleToken() {
	token := Token("secret-abc123")
	val := token.LogValue()
	fmt.Println("Token value:", val.String())
	// Output: Token value: REDACTED
}

// ExampleWithRequestID demonstrates wrapping a slog.Handler.
func ExampleWithRequestID() {
	h := &CustomHandler{}
	wrapped := WithRequestID(h, "req-123")
	logger := slog.New(wrapped)

	logger.Info("Request processed")
	// Output: [INFO] Request processed request_id=req-123
}