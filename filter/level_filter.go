package filter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/user/siftlog/output"
)

// levelFieldCandidates lists JSON field names commonly used for log level.
var levelFieldCandidates = []string{"level", "lvl", "severity", "log_level"}

// LevelFilter discards log lines whose level is below the minimum threshold.
type LevelFilter struct {
	Min output.Level
}

// NewLevelFilter parses a minimum level string and returns a LevelFilter.
// Returns an error if the level string is not recognised.
func NewLevelFilter(minLevel string) (*LevelFilter, error) {
	l := output.ParseLevel(minLevel)
	if l == output.LevelUnknown {
		return nil, fmt.Errorf("unknown log level %q", minLevel)
	}
	return &LevelFilter{Min: l}, nil
}

// Allow returns true when the line's level is >= the minimum threshold.
// Lines that are not valid JSON or have no recognised level field are allowed
// through so non-structured lines are never silently dropped.
func (lf *LevelFilter) Allow(line string) bool {
	line = strings.TrimSpace(line)
	if line == "" {
		return false
	}
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(line), &obj); err != nil {
		return true // non-JSON passes through
	}
	for _, key := range levelFieldCandidates {
		if val, ok := obj[key]; ok {
			raw := fmt.Sprintf("%v", val)
			parsed := output.ParseLevel(raw)
			if parsed == output.LevelUnknown {
				return true
			}
			return parsed >= lf.Min
		}
	}
	return true // no level field — pass through
}
