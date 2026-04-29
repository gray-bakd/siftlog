package pipeline_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/siftlog/pipeline"
)

func TestRun_WithHighlight_MarksFields(t *testing.T) {
	input := strings.NewReader(`{"level":"info","msg":"started","svc":"api"}` + "\n")
	var out bytes.Buffer

	err := pipeline.Run(pipeline.Options{
		Input:     input,
		Output:    &out,
		Highlight: []string{"msg"},
		Color:     false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := out.String()
	if !strings.Contains(result, "[msg=started]") {
		t.Errorf("expected msg to be highlighted, got: %q", result)
	}
}

func TestRun_WithHighlight_FilteredOut(t *testing.T) {
	input := strings.NewReader(`{"level":"debug","msg":"verbose"}` + "\n")
	var out bytes.Buffer

	err := pipeline.Run(pipeline.Options{
		Input:     input,
		Output:    &out,
		MinLevel:  "info",
		Highlight: []string{"msg"},
		Color:     false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out.Len() != 0 {
		t.Errorf("expected no output for filtered line, got: %q", out.String())
	}
}

func TestRun_WithHighlight_ColorEnabled(t *testing.T) {
	input := strings.NewReader(`{"level":"warn","msg":"watch out"}` + "\n")
	var out bytes.Buffer

	err := pipeline.Run(pipeline.Options{
		Input:     input,
		Output:    &out,
		Highlight: []string{"msg"},
		Color:     true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := out.String()
	if !strings.Contains(result, "\033[1m") {
		t.Errorf("expected ANSI bold in color output, got: %q", result)
	}
}
