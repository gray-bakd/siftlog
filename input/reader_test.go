package input

import (
	"strings"
	"testing"
)

func TestLineReader_MultipleLines(t *testing.T) {
	input := "line one\nline two\nline three"
	r := NewLineReader(strings.NewReader(input))

	expected := []string{"line one", "line two", "line three"}
	for _, want := range expected {
		got, ok := r.Next()
		if !ok {
			t.Fatalf("expected more lines, got EOF")
		}
		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	}

	_, ok := r.Next()
	if ok {
		t.Error("expected EOF after all lines consumed")
	}

	if err := r.Err(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLineReader_EmptyInput(t *testing.T) {
	r := NewLineReader(strings.NewReader(""))
	_, ok := r.Next()
	if ok {
		t.Error("expected no lines from empty input")
	}
	if err := r.Err(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLineReader_SingleLine(t *testing.T) {
	r := NewLineReader(strings.NewReader(`{"level":"info","msg":"hello"}`))
	line, ok := r.Next()
	if !ok {
		t.Fatal("expected one line")
	}
	if line != `{"level":"info","msg":"hello"}` {
		t.Errorf("unexpected line: %q", line)
	}
	_, ok = r.Next()
	if ok {
		t.Error("expected EOF after single line")
	}
}

func TestLineReader_BlankLines(t *testing.T) {
	r := NewLineReader(strings.NewReader("a\n\nb"))
	lines := []string{}
	for {
		line, ok := r.Next()
		if !ok {
			break
		}
		lines = append(lines, line)
	}
	if len(lines) != 3 {
		t.Errorf("expected 3 lines (including blank), got %d", len(lines))
	}
	if lines[1] != "" {
		t.Errorf("expected blank line at index 1, got %q", lines[1])
	}
}
