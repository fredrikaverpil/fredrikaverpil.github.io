package examples

import (
	"encoding"
	"encoding/json"
	"time"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// ToText serializes any TextMarshaler to string using the encoding.TextMarshaler interface.
func ToText(m encoding.TextMarshaler) (string, error) {
	b, err := m.MarshalText()
	return string(b), err
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// Status is a custom type with text marshaling that implements encoding.TextMarshaler.
type Status int

const (
	Running Status = iota
	Stopped
)

func (s Status) MarshalText() ([]byte, error) {
	if s == Running {
		return []byte("running"), nil
	}
	return []byte("stopped"), nil
}

// Timestamp implements json.Marshaler to customize JSON serialization.
type Timestamp int64

func (t Timestamp) MarshalJSON() ([]byte, error) {
	s := time.Unix(int64(t), 0).Format(time.RFC3339)
	// Use json.Marshal to ensure the string is correctly quoted and escaped.
	return json.Marshal(s)
}

// LogEntry uses Timestamp for JSON serialization.
type LogEntry struct {
	Msg  string
	Time Timestamp
}
