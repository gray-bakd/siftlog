package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ExcludeFilter rejects log lines where a specific field matches a given value.
type ExcludeFilter struct {
	field string
	value string
}

// NewExcludeFilter creates a filter that rejects lines where field equals value (case-insensitive).
func NewExcludeFilter(field, value string) (*ExcludeFilter, error) {
	if field == "" {
		return nil, fmt.Errorf("exclude filter: field must not be empty")
	}
	if value == "" {
		return nil, fmt.Errorf("exclude filter: value must not be empty")
	}
	return &ExcludeFilter{
		field: field,
		value: strings.ToLower(value),
	}, nil
}

// Allow returns false if the line's field matches the excluded value.
func (f *ExcludeFilter) Allow(line string) bool {
	if line == "" {
		return true
	}
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return true
	}
	val, ok := obj[f.field]
	if !ok {
		return true
	}
	actual := strings.ToLower(fmt.Sprintf("%v", val))
	return actual != f.value
}
