package examples

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// DebugSQLValue inspects what a custom type sends to the database using driver.Valuer.
func DebugSQLValue(v driver.Valuer) string {
	val, err := v.Value()
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	if b, ok := val.([]byte); ok {
		return fmt.Sprintf("SQL Value: %s", string(b))
	}
	return fmt.Sprintf("SQL Value: %v", val)
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// JSONB is a custom JSON type for databases that implements both sql.Scanner and driver.Valuer.
type JSONB map[string]any

func (j *JSONB) Scan(v any) error {
	var bytes []byte
	switch x := v.(type) {
	case []byte:
		bytes = x
	case string:
		bytes = []byte(x)
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into JSONB", v)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}
