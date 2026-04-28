package output

import (
	"testing"
)

func TestFormat_Raw(t *testing.T) {
	input := `{"level":"info","msg":"hello"}`
	got := Format(input, FormatRaw)
	if got != input {
		t.Errorf("FormatRaw: expected %q, got %q", input, got)
	}
}

func TestFormat_Compact(t *testing.T) {
	// Compact should normalise spacing.
	input := `{  "level" :  "info",  "msg":  "hello"  }`
	got := Format(input, FormatCompact)
	want := `{"level":"info","msg":"hello"}`
	if got != want {
		t.Errorf("FormatCompact: expected %q, got %q", want, got)
	}
}

func TestFormat_Pretty(t *testing.T) {
	input := `{"level":"info","msg":"hello"}`
	got := Format(input, FormatPretty)
	// Pretty output must contain a newline and indentation.
	if len(got) <= len(input) {
		t.Errorf("FormatPretty: expected longer output than %d chars, got %d", len(input), len(got))
	}
	if got[0] != '{' {
		t.Errorf("FormatPretty: expected output to start with '{', got %q", string(got[0]))
	}
}

func TestFormat_InvalidJSON_ReturnsOriginal(t *testing.T) {
	input := `not valid json`
	for _, mode := range []FormatMode{FormatRaw, FormatPretty, FormatCompact} {
		got := Format(input, mode)
		if got != input {
			t.Errorf("mode %d: expected original line for invalid JSON, got %q", mode, got)
		}
	}
}

func TestFormat_EmptyLine(t *testing.T) {
	for _, mode := range []FormatMode{FormatRaw, FormatPretty, FormatCompact} {
		got := Format("", mode)
		if got != "" {
			t.Errorf("mode %d: expected empty string for empty input, got %q", mode, got)
		}
	}
}

func TestParseFormatMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  FormatMode
	}{
		{"raw", FormatRaw},
		{"", FormatRaw},
		{"pretty", FormatPretty},
		{"PRETTY", FormatPretty},
		{"compact", FormatCompact},
		{"Compact", FormatCompact},
	}
	for _, c := range cases {
		got, err := ParseFormatMode(c.input)
		if err != nil {
			t.Errorf("ParseFormatMode(%q): unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseFormatMode(%q): expected %d, got %d", c.input, c.want, got)
		}
	}
}

func TestParseFormatMode_Invalid(t *testing.T) {
	_, err := ParseFormatMode("xml")
	if err == nil {
		t.Error("ParseFormatMode(\"xml\"): expected error, got nil")
	}
}
