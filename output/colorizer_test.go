package output

import (
	"strings"
	"testing"
)

func TestColorize_ValidJSON_NoColor(t *testing.T) {
	line := []byte(`{"level":"info","msg":"started","port":8080}`)
	out, err := Colorize(line, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "level") {
		t.Errorf("expected output to contain field 'level', got: %s", out)
	}
	if strings.Contains(out, "\033[") {
		t.Errorf("expected no ANSI codes in noColor mode, got: %s", out)
	}
}

func TestColorize_ValidJSON_WithColor(t *testing.T) {
	line := []byte(`{"level":"error","msg":"something failed"}`)
	out, err := Colorize(line, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "\033[") {
		t.Errorf("expected ANSI codes in color mode, got: %s", out)
	}
	if !strings.Contains(out, Red) {
		t.Errorf("expected red color for error level, got: %s", out)
	}
}

func TestColorize_InvalidJSON(t *testing.T) {
	line := []byte(`not json`)
	_, err := Colorize(line, false)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestColorize_LevelColors(t *testing.T) {
	cases := []struct {
		level    string
		expected string
	}{
		{"info", Green},
		{"warn", Yellow},
		{"debug", Cyan},
		{"error", Red},
		{"trace", Gray},
	}

	for _, tc := range cases {
		line := []byte(`{"level":"` + tc.level + `","msg":"test"}`)
		out, err := Colorize(line, false)
		if err != nil {
			t.Fatalf("level %s: unexpected error: %v", tc.level, err)
		}
		if !strings.Contains(out, tc.expected) {
			t.Errorf("level %s: expected color %q in output, got: %s", tc.level, tc.expected, out)
		}
	}
}
