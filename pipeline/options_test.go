package pipeline

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/siftlog/output"
)

func TestOptions_Validate_Valid(t *testing.T) {
	opts := &Options{
		Input:      strings.NewReader(""),
		Output:     &bytes.Buffer{},
		FormatMode: output.FormatRaw,
	}
	if err := opts.Validate(); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}

func TestOptions_Validate_NilInput(t *testing.T) {
	opts := &Options{
		Input:  nil,
		Output: &bytes.Buffer{},
	}
	if err := opts.Validate(); err == nil {
		t.Error("expected error for nil input")
	}
}

func TestOptions_Validate_NilOutput(t *testing.T) {
	opts := &Options{
		Input:  strings.NewReader(""),
		Output: nil,
	}
	if err := opts.Validate(); err == nil {
		t.Error("expected error for nil output")
	}
}

func TestOptions_Validate_InvalidMinLevel(t *testing.T) {
	opts := &Options{
		Input:    strings.NewReader(""),
		Output:   &bytes.Buffer{},
		MinLevel: "superverbose",
	}
	if err := opts.Validate(); err == nil {
		t.Error("expected error for invalid min-level")
	}
}

func TestOptions_Validate_ValidMinLevel(t *testing.T) {
	opts := &Options{
		Input:    strings.NewReader(""),
		Output:   &bytes.Buffer{},
		MinLevel: "warn",
	}
	if err := opts.Validate(); err != nil {
		t.Errorf("expected no error for valid min-level, got: %v", err)
	}
}

func TestOptions_Validate_InvalidQuery(t *testing.T) {
	opts := &Options{
		Input:   strings.NewReader(""),
		Output:  &bytes.Buffer{},
		Queries: []string{"noequalssign"},
	}
	if err := opts.Validate(); err == nil {
		t.Error("expected error for invalid query")
	}
}
