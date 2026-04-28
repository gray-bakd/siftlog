package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FieldFilter matches log lines where a specific field contains a substring.
type FieldFilter struct {
	Field   string
	Pattern string
}

// NewFieldFilter creates a FieldFilter from a "field=pattern" string.
// Returns an error if the expression is not in the expected format.
func NewFieldFilter(expr string) (*FieldFilter, error) {
	parts := strings.SplitN(expr, "=", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("invalid field filter expression %q: expected field=pattern", expr)
	}
	return &FieldFilter{
		Field:   parts[0],
		Pattern: parts[1],
	}, nil
}

// Allow returns true if the given JSON log line contains the field and
// its string value contains the filter's pattern (case-insensitive).
func (f *FieldFilter) Allow(line string) bool {
	if line == "" {
		return false
	}

	var record map[string]interface{}
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return false
	}

	val, ok := record[f.Field]
	if !ok {
		return false
	}

	strVal := fmt.Sprintf("%v", val)
	return strings.Contains(
		strings.ToLower(strVal),
		strings.ToLower(f.Pattern),
	)
}
