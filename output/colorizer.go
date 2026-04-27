package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[97m"
	Gray   = "\033[90m"
)

// LevelColors maps common log level values to ANSI colors.
var LevelColors = map[string]string{
	"error": Red,
	"fatal": Red,
	"warn":  Yellow,
	"warning": Yellow,
	"info":  Green,
	"debug": Cyan,
	"trace": Gray,
}

// Colorize formats a raw JSON log line with ANSI colors.
// If noColor is true, it returns the line as-is (pretty-printed).
func Colorize(line []byte, noColor bool) (string, error) {
	var record map[string]interface{}
	if err := json.Unmarshal(line, &record); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	if noColor {
		pretty, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			return "", err
		}
		return string(pretty), nil
	}

	var sb strings.Builder

	levelColor := White
	for _, key := range []string{"level", "severity", "lvl"} {
		if val, ok := record[key]; ok {
			if lvl, ok := val.(string); ok {
				if c, found := LevelColors[strings.ToLower(lvl)]; found {
					levelColor = c
				}
			}
			break
		}
	}

	sb.WriteString(levelColor)
	for k, v := range record {
		sb.WriteString(fmt.Sprintf("%s%s=%s%v%s ", Blue, k, White, v, levelColor))
	}
	sb.WriteString(Reset)

	return strings.TrimSpace(sb.String()), nil
}
