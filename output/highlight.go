package output

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// HighlightFields returns a formatted string with specific fields highlighted
// using ANSI bold/underline codes when color is enabled.
func HighlightFields(line string, fields []string, color bool) string {
	if len(fields) == 0 || line == "" {
		return line
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return line
	}

	if !color {
		return formatHighlightedPlain(obj, fields)
	}
	return formatHighlightedColor(obj, fields)
}

func formatHighlightedPlain(obj map[string]interface{}, fields []string) string {
	set := toSet(fields)
	var parts []string
	keys := sortedKeys(obj)
	for _, k := range keys {
		v := obj[k]
		if set[strings.ToLower(k)] {
			parts = append(parts, fmt.Sprintf("[%s=%v]", k, v))
		} else {
			parts = append(parts, fmt.Sprintf("%s=%v", k, v))
		}
	}
	return strings.Join(parts, " ")
}

func formatHighlightedColor(obj map[string]interface{}, fields []string) string {
	const bold = "\033[1m"
	const underline = "\033[4m"
	const reset = "\033[0m"

	set := toSet(fields)
	var parts []string
	keys := sortedKeys(obj)
	for _, k := range keys {
		v := obj[k]
		if set[strings.ToLower(k)] {
			parts = append(parts, fmt.Sprintf("%s%s%s=%v%s", bold, underline, k, v, reset))
		} else {
			parts = append(parts, fmt.Sprintf("%s=%v", k, v))
		}
	}
	return strings.Join(parts, " ")
}

func toSet(fields []string) map[string]bool {
	s := make(map[string]bool, len(fields))
	for _, f := range fields {
		s[strings.ToLower(f)] = true
	}
	return s
}

func sortedKeys(obj map[string]interface{}) []string {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
