package output

import (
	"testing"
)

func TestParseLevel_KnownValues(t *testing.T) {
	cases := []struct {
		input    string
		expected Level
	}{
		{"info", LevelInfo},
		{"INFO", LevelInfo},
		{"Info", LevelInfo},
		{"debug", LevelDebug},
		{"DEBUG", LevelDebug},
		{"warn", LevelWarn},
		{"warning", LevelWarn},
		{"WARNING", LevelWarn},
		{"error", LevelError},
		{"ERROR", LevelError},
		{"err", LevelError},
		{"fatal", LevelFatal},
		{"panic", LevelFatal},
		{"FATAL", LevelFatal},
		{"trace", LevelTrace},
		{"TRACE", LevelTrace},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got := ParseLevel(tc.input)
			if got != tc.expected {
				t.Errorf("ParseLevel(%q) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}

func TestParseLevel_UnknownValues(t *testing.T) {
	cases := []string{"", "verbose", "notice", "critical", "42"}
	for _, input := range cases {
		t.Run(input, func(t *testing.T) {
			got := ParseLevel(input)
			if got != LevelUnknown {
				t.Errorf("ParseLevel(%q) = %v, want LevelUnknown", input, got)
			}
		})
	}
}

func TestLevel_String(t *testing.T) {
	cases := []struct {
		level    Level
		expected string
	}{
		{LevelTrace, "TRACE"},
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{LevelFatal, "FATAL"},
		{LevelUnknown, "UNKNOWN"},
	}
	for _, tc := range cases {
		if got := tc.level.String(); got != tc.expected {
			t.Errorf("Level(%d).String() = %q, want %q", tc.level, got, tc.expected)
		}
	}
}
