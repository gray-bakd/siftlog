package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FormatMode controls how output lines are rendered.
type FormatMode int

const (
	FormatRaw    FormatMode = iota // original JSON line as-is
	FormatPretty                   // indented multi-line JSON
	FormatCompact                  // single-line minified JSON
)

// Format takes a raw JSON log line and returns it formatted according to mode.
// Non-JSON lines are returned unchanged regardless of mode.
func Format(line string, mode FormatMode) string {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return line
	}

	switch mode {
	case FormatPretty:
		var obj interface{}
		if err := json.Unmarshal([]byte(trimmed), &obj); err != nil {
			return line
		}
		b, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return line
		}
		return string(b)

	case FormatCompact:
		var obj interface{}
		if err := json.Unmarshal([]byte(trimmed), &obj); err != nil {
			return line
		}
		b, err := json.Marshal(obj)
		if err != nil {
			return line
		}
		return string(b)

	default: // FormatRaw
		return line
	}
}

// ParseFormatMode converts a string flag value to a FormatMode.
func ParseFormatMode(s string) (FormatMode, error) {
	switch strings.ToLower(s) {
	case "raw", "":
		return FormatRaw, nil
	case "pretty":
		return FormatPretty, nil
	case "compact":
		return FormatCompact, nil
	default:
		return FormatRaw, fmt.Errorf("unknown format mode %q: must be raw, pretty, or compact", s)
	}
}
