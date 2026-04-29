package filter_test

import (
	"testing"

	"github.com/user/siftlog/filter"
)

func TestNewRegexFilter_Valid(t *testing.T) {
	f, err := filter.NewRegexFilter("message", `error.*timeout`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestNewRegexFilter_EmptyField(t *testing.T) {
	_, err := filter.NewRegexFilter("", `.*`)
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewRegexFilter_EmptyPattern(t *testing.T) {
	_, err := filter.NewRegexFilter("message", "")
	if err == nil {
		t.Fatal("expected error for empty pattern")
	}
}

func TestNewRegexFilter_InvalidPattern(t *testing.T) {
	_, err := filter.NewRegexFilter("message", `[invalid`)
	if err == nil {
		t.Fatal("expected error for invalid regex pattern")
	}
}

func TestRegexFilter_Allow_Match(t *testing.T) {
	f, _ := filter.NewRegexFilter("message", `timeout`)
	line := `{"message": "connection timeout occurred"}`
	if !f.Allow(line) {
		t.Error("expected line to be allowed")
	}
}

func TestRegexFilter_Allow_NoMatch(t *testing.T) {
	f, _ := filter.NewRegexFilter("message", `^fatal`)
	line := `{"message": "info: all systems nominal"}`
	if f.Allow(line) {
		t.Error("expected line to be rejected")
	}
}

func TestRegexFilter_Allow_FieldMissing(t *testing.T) {
	f, _ := filter.NewRegexFilter("message", `.*`)
	line := `{"level": "error"}`
	if f.Allow(line) {
		t.Error("expected line with missing field to be rejected")
	}
}

func TestRegexFilter_Allow_NonStringField(t *testing.T) {
	f, _ := filter.NewRegexFilter("code", `404`)
	line := `{"code": 404}`
	if f.Allow(line) {
		t.Error("expected non-string field to be rejected")
	}
}

func TestRegexFilter_Allow_InvalidJSON(t *testing.T) {
	f, _ := filter.NewRegexFilter("message", `.*`)
	if f.Allow("not-json") {
		t.Error("expected invalid JSON to be rejected")
	}
}

func TestRegexFilter_Allow_EmptyLine(t *testing.T) {
	f, _ := filter.NewRegexFilter("message", `.*`)
	if f.Allow("") {
		t.Error("expected empty line to be rejected")
	}
}
