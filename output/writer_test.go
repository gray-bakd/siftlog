package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewWriter_NotNil(t *testing.T) {
	w := NewWriter(&bytes.Buffer{}, FormatRaw, true)
	if w == nil {
		t.Fatal("expected non-nil Writer")
	}
}

func TestWriteLine_RawNoColor(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf, FormatRaw, true)

	line := `{"level":"info","msg":"hello"}`
	if err := w.WriteLine(line); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := strings.TrimRight(buf.String(), "\n")
	if got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestWriteLine_CompactNoColor(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf, FormatCompact, true)

	line := `{ "level" : "info" , "msg" : "hello" }`
	if err := w.WriteLine(line); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := strings.TrimRight(buf.String(), "\n")
	if strings.Contains(got, "  ") {
		t.Errorf("compact output should not contain extra spaces: %q", got)
	}
}

func TestWriteLine_InvalidJSON_PassesThrough(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf, FormatRaw, true)

	line := "not json at all"
	if err := w.WriteLine(line); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := strings.TrimRight(buf.String(), "\n")
	if got != line {
		t.Errorf("expected original line %q, got %q", line, got)
	}
}

func TestWriteLine_MultipleLines(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf, FormatRaw, true)

	lines := []string{
		`{"level":"info","msg":"first"}`,
		`{"level":"error","msg":"second"}`,
	}
	for _, l := range lines {
		if err := w.WriteLine(l); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	result := buf.String()
	for _, l := range lines {
		if !strings.Contains(result, l) {
			t.Errorf("expected output to contain %q", l)
		}
	}
}
