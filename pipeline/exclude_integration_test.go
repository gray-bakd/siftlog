package pipeline

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/siftlog/filter"
	"github.com/user/siftlog/input"
	"github.com/user/siftlog/output"
)

func TestRun_WithExcludeFilter_RemovesMatchingLines(t *testing.T) {
	raw := strings.Join([]string{
		`{"level":"debug","msg":"verbose"}`,
		`{"level":"info","msg":"hello"}`,
		`{"level":"debug","msg":"trace"}`,
		`{"level":"warn","msg":"careful"}`,
	}, "\n")

	exclude, err := filter.NewExcludeFilter("level", "debug")
	if err != nil {
		t.Fatalf("failed to create exclude filter: %v", err)
	}

	var buf bytes.Buffer
	opts := Options{
		Input:  input.NewLineReader(strings.NewReader(raw)),
		Output: output.NewWriter(&buf, false, output.FormatRaw),
		Filter: exclude,
	}

	if err := Run(opts); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	result := buf.String()
	if strings.Contains(result, "verbose") || strings.Contains(result, "trace") {
		t.Error("expected debug lines to be excluded")
	}
	if !strings.Contains(result, "hello") || !strings.Contains(result, "careful") {
		t.Error("expected non-debug lines to be included")
	}
}

func TestRun_WithExcludeAndFieldOutput_CombinesCorrectly(t *testing.T) {
	raw := strings.Join([]string{
		`{"level":"info","msg":"keep","secret":"hidden"}`,
		`{"level":"debug","msg":"drop","secret":"hidden"}`,
	}, "\n")

	exclude, _ := filter.NewExcludeFilter("level", "debug")

	var buf bytes.Buffer
	opts := Options{
		Input:  input.NewLineReader(strings.NewReader(raw)),
		Output: output.NewWriter(&buf, false, output.FormatRaw),
		Filter: exclude,
		ExcludeFields: []string{"secret"},
	}

	if err := Run(opts); err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	result := buf.String()
	if strings.Contains(result, "drop") {
		t.Error("expected debug line to be excluded")
	}
	if strings.Contains(result, "hidden") {
		t.Error("expected secret field to be excluded from output")
	}
	if !strings.Contains(result, "keep") {
		t.Error("expected info line to be included")
	}
}
