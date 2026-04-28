package filter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Query represents a single field=value filter condition.
type Query struct {
	Field string
	Value string
}

// ParseQueries parses CLI arguments of the form "field=value" into Query structs.
func ParseQueries(args []string) ([]Query, error) {
	queries := make([]Query, 0, len(args))
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 || parts[0] == "" {
			return nil, fmt.Errorf("expected field=value, got %q", arg)
		}
		queries = append(queries, Query{Field: parts[0], Value: parts[1]})
	}
	return queries, nil
}

// Match returns true if the JSON log line satisfies all provided queries.
func Match(line string, queries []Query) bool {
	var record map[string]interface{}
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		// Not valid JSON — skip line
		return false
	}

	for _, q := range queries {
		val, ok := record[q.Field]
		if !ok {
			return false
		}
		// Compare as string representation
		if fmt.Sprintf("%v", val) != q.Value {
			return false
		}
	}
	return true
}

// MatchAny returns true if the JSON log line satisfies at least one of the
// provided queries. Returns false if queries is empty or the line is not
// valid JSON.
func MatchAny(line string, queries []Query) bool {
	if len(queries) == 0 {
		return false
	}
	var record map[string]interface{}
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return false
	}
	for _, q := range queries {
		val, ok := record[q.Field]
		if ok && fmt.Sprintf("%v", val) == q.Value {
			return true
		}
	}
	return false
}
