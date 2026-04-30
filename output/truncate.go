package output

import (
	"encoding/json"
	"fmt"
)

// TruncateOptions controls how field values are truncated.
type TruncateOptions struct {
	MaxLength int
	Suffix    string
}

// DefaultTruncateOptions returns sensible defaults.
func DefaultTruncateOptions() TruncateOptions {
	return TruncateOptions{
		MaxLength: 120,
		Suffix:    "…",
	}
}

// TruncateFields parses the JSON line and truncates string values that exceed
// MaxLength, replacing the overflow with Suffix. Non-JSON lines are returned
// unchanged. A MaxLength of 0 disables truncation.
func TruncateFields(line string, opts TruncateOptions) string {
	if line == "" || opts.MaxLength == 0 {
		return line
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return line
	}

	modified := false
	for k, v := range obj {
		if s, ok := v.(string); ok {
			runes := []rune(s)
			if len(runes) > opts.MaxLength {
				obj[k] = string(runes[:opts.MaxLength]) + opts.Suffix
				modified = true
			}
		}
	}

	if !modified {
		return line
	}

	out, err := json.Marshal(obj)
	if err != nil {
		return line
	}
	return fmt.Sprintf("%s", out)
}
