package examples

import "strings"

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// CSVFlag implements flag.Value for custom command-line flag parsing.
type CSVFlag []string

func (c *CSVFlag) String() string { return strings.Join(*c, ",") }
func (c *CSVFlag) Set(v string) error {
	*c = append(*c, strings.Split(v, ",")...)
	return nil
}
