package examples

import "fmt"

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// FormatLabel accepts fmt.Stringer to work with any type that has a string representation.
func FormatLabel(s fmt.Stringer) string {
	return fmt.Sprintf("[%s]", s.String())
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// UserID implements fmt.Stringer to customize how the type is printed.
type UserID int

func (u UserID) String() string {
	return fmt.Sprintf("user-%d", u)
}

// Config implements fmt.GoStringer to control debug formatting with %#v.
type Config struct {
	Host     string
	Password string
}

func (c Config) GoString() string {
	return fmt.Sprintf("Config{Host:%q, Password:\"***\"}", c.Host)
}
