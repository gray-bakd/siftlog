package output

import "strings"

// Level represents a log severity level.
type Level int

const (
	LevelUnknown Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// levelNames maps common string representations to Level values.
var levelNames = map[string]Level{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"warning": LevelWarn,
	"error": LevelError,
	"err":   LevelError,
	"fatal": LevelFatal,
	"panic": LevelFatal,
}

// ParseLevel converts a raw string value from a log field into a Level.
func ParseLevel(raw string) Level {
	norm := strings.ToLower(strings.TrimSpace(raw))
	if l, ok := levelNames[norm]; ok {
		return l
	}
	return LevelUnknown
}

// String returns the canonical name of the level.
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "TRACE"
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
