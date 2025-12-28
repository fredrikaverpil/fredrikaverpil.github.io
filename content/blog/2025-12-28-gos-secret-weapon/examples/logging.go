package examples

import (
	"context"
	"fmt"
	"log/slog"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// WithRequestID wraps a slog.Handler to add a request ID attribute.
// This demonstrates how accepting slog.Handler allows you to decorate any logging backend.
func WithRequestID(h slog.Handler, requestID string) slog.Handler {
	return h.WithAttrs([]slog.Attr{slog.String("request_id", requestID)})
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// CustomHandler implements slog.Handler to demonstrate a simple text logger.
type CustomHandler struct {
	attrs []slog.Attr
}

func (h *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	// Print the level and message
	fmt.Printf("[%s] %s", r.Level, r.Message)

	// Print handler attributes (e.g. from WithAttrs)
	for _, a := range h.attrs {
		fmt.Printf(" %s=%v", a.Key, a.Value.Any())
	}

	// Print record attributes (e.g. from logger.Info args)
	r.Attrs(func(a slog.Attr) bool {
		fmt.Printf(" %s=%v", a.Key, a.Value.Any())
		return true
	})

	fmt.Println()
	return nil
}

// WithAttrs returns a new handler with the added attributes.
func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Copy existing attributes to the new handler
	newAttrs := make([]slog.Attr, len(h.attrs)+len(attrs))
	copy(newAttrs, h.attrs)
	copy(newAttrs[len(h.attrs):], attrs)
	return &CustomHandler{attrs: newAttrs}
}

// WithGroup is a no-op in this simplified example.
func (h *CustomHandler) WithGroup(name string) slog.Handler { return h }

// Token implements slog.LogValuer for redacting sensitive data.
type Token string

func (t Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED")
}
