package filter

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// RegexFilter allows log lines where a specified field matches a regular expression.
type RegexFilter struct {
	field  string
	regexp *regexp.Regexp
}

// NewRegexFilter creates a RegexFilter for the given field and pattern.
// Returns an error if the pattern is not a valid regular expression.
func NewRegexFilter(field, pattern string) (*RegexFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("regex filter: field name must not be empty")
	}
	if pattern == "" {
		return nil, fmt.Errorf("regex filter: pattern must not be empty")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("regex filter: invalid pattern %q: %w", pattern, err)
	}
	return &RegexFilter{field: field, regexp: re}, nil
}

// Allow returns true if the log line's specified field value matches the regex.
// Lines with missing fields or non-string values are rejected.
func (f *RegexFilter) Allow(line string) bool {
	if line == "" {
		return false
	}
	var record map[string]interface{}
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return false
	}
	val, ok := record[f.field]
	if !ok {
		return false
	}
	str, ok := val.(string)
	if !ok {
		return false
	}
	return f.regexp.MatchString(str)
}
