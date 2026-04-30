package pipeline_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/siftlog/output"
	"github.com/siftlog/pipeline"
)

func TestRun_WithTruncate_ShortensLongFields(t *testing.T) {
	long := strings.Repeat("x", 200)
	input := `{"level":"info","msg":"` + long + `"}` + "\n"

	reader := strings.NewReader(input)
	var buf bytes.Buffer

	opts := pipeline.Options{
		Input:      reader,
		Output:     &buf,
		TruncateOn: true,
		Truncate:   output.TruncateOptions{MaxLength: 20, Suffix: "…"},
		FormatMode: output.FormatRaw,
	}

	if err := pipeline.Run(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &obj); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
	msg, _ := obj["msg"].(string)
	if len([]rune(msg)) > 21 {
		t.Errorf("msg not truncated, length=%d", len(msg))
	}
}

func TestRun_WithTruncate_Disabled_LeavesLongFields(t *testing.T) {
	long := strings.Repeat("y", 200)
	input := `{"level":"info","msg":"` + long + `"}` + "\n"

	reader := strings.NewReader(input)
	var buf bytes.Buffer

	opts := pipeline.Options{
		Input:      reader,
		Output:     &buf,
		TruncateOn: false,
		FormatMode: output.FormatRaw,
	}

	if err := pipeline.Run(opts); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), long) {
		t.Error("expected long field to be present when truncation disabled")
	}
}
