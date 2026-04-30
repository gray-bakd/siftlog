package output

import "encoding/json"

// ExcludeOptions configures which fields to remove from output.
type ExcludeOptions struct {
	Fields []string
}

// ExcludeFields removes specified fields from a JSON log line.
// If the line is not valid JSON or no fields are specified, it is returned unchanged.
func ExcludeFields(line string, opts ExcludeOptions) string {
	if line == "" || len(opts.Fields) == 0 {
		return line
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return line
	}

	excludeSet := toExcludeSet(opts.Fields)
	for field := range excludeSet {
		delete(obj, field)
	}

	out, err := json.Marshal(obj)
	if err != nil {
		return line
	}
	return string(out)
}

func toExcludeSet(fields []string) map[string]struct{} {
	s := make(map[string]struct{}, len(fields))
	for _, f := range fields {
		if f != "" {
			s[f] = struct{}{}
		}
	}
	return s
}
