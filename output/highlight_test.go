package output

import (
	"strings"
	"testing"
)

func TestHighlightFields_EmptyFields(t *testing.T) {
	line := `{"level":"info","msg":"hello"}`
	result := HighlightFields(line, nil, false)
	if result != line {
		t.Errorf("expected unchanged line, got %q", result)
	}
}

func TestHighlightFields_EmptyLine(t *testing.T) {
	result := HighlightFields("", []string{"msg"}, false)
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestHighlightFields_InvalidJSON(t *testing.T) {
	line := "not json"
	result := HighlightFields(line, []string{"msg"}, false)
	if result != line {
		t.Errorf("expected original line returned, got %q", result)
	}
}

func TestHighlightFields_PlainHighlight(t *testing.T) {
	line := `{"level":"info","msg":"hello"}`
	result := HighlightFields(line, []string{"msg"}, false)
	if !strings.Contains(result, "[msg=hello]") {
		t.Errorf("expected msg to be bracketed, got %q", result)
	}
	if strings.Contains(result, "[level=") {
		t.Errorf("expected level to not be bracketed, got %q", result)
	}
}

func TestHighlightFields_ColorHighlight(t *testing.T) {
	line := `{"level":"info","msg":"hello"}`
	result := HighlightFields(line, []string{"msg"}, true)
	if !strings.Contains(result, "\033[1m") {
		t.Errorf("expected bold ANSI code in result, got %q", result)
	}
	if !strings.Contains(result, "\033[0m") {
		t.Errorf("expected reset ANSI code in result, got %q", result)
	}
}

func TestHighlightFields_CaseInsensitiveField(t *testing.T) {
	line := `{"Level":"info","Msg":"hello"}`
	result := HighlightFields(line, []string{"msg"}, false)
	if !strings.Contains(result, "[Msg=hello]") {
		t.Errorf("expected Msg to be highlighted case-insensitively, got %q", result)
	}
}

func TestHighlightFields_MultipleFields(t *testing.T) {
	line := `{"level":"warn","msg":"oops","code":"500"}`
	result := HighlightFields(line, []string{"msg", "code"}, false)
	if !strings.Contains(result, "[msg=oops]") {
		t.Errorf("expected msg highlighted, got %q", result)
	}
	if !strings.Contains(result, "[code=500]") {
		t.Errorf("expected code highlighted, got %q", result)
	}
}
