package pipeline

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_NoFilter_OutputsAllLines(t *testing.T) {
	input := `{"level":"info","msg":"started"}
{"level":"error","msg":"failed"}
`
	var out bytes.Buffer
	err := Run(strings.NewReader(input), Options{
		Queries: nil,
		NoColor: true,
		Out:     &out,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 output lines, got %d", len(lines))
	}
}

func TestRun_WithFilter_OnlyMatchingLines(t *testing.T) {
	input := `{"level":"info","msg":"started"}
{"level":"error","msg":"failed"}
{"level":"info","msg":"done"}
`
	var out bytes.Buffer
	err := Run(strings.NewReader(input), Options{
		Queries: []string{"level=info"},
		NoColor: true,
		Out:     &out,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 matching lines, got %d", len(lines))
	}
	for _, l := range lines {
		if !strings.Contains(l, `"info"`) {
			t.Errorf("unexpected line in output: %s", l)
		}
	}
}

func TestRun_InvalidQuery_ReturnsError(t *testing.T) {
	var out bytes.Buffer
	err := Run(strings.NewReader(""), Options{
		Queries: []string{"notavalidquery"},
		NoColor: true,
		Out:     &out,
	})
	if err == nil {
		t.Error("expected error for invalid query, got nil")
	}
}

func TestRun_EmptyInput_NoOutput(t *testing.T) {
	var out bytes.Buffer
	err := Run(strings.NewReader(""), Options{
		NoColor: true,
		Out:     &out,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("expected no output, got: %q", out.String())
	}
}
