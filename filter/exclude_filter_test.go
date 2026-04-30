package filter

import "testing"

func TestNewExcludeFilter_Valid(t *testing.T) {
	f, err := NewExcludeFilter("level", "debug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestNewExcludeFilter_EmptyField(t *testing.T) {
	_, err := NewExcludeFilter("", "debug")
	if err == nil {
		t.Fatal("expected error for empty field")
	}
}

func TestNewExcludeFilter_EmptyValue(t *testing.T) {
	_, err := NewExcludeFilter("level", "")
	if err == nil {
		t.Fatal("expected error for empty value")
	}
}

func TestExcludeFilter_Allow_MatchExcludes(t *testing.T) {
	f, _ := NewExcludeFilter("level", "debug")
	line := `{"level":"debug","msg":"verbose"}`
	if f.Allow(line) {
		t.Error("expected line to be excluded")
	}
}

func TestExcludeFilter_Allow_NoMatchAllows(t *testing.T) {
	f, _ := NewExcludeFilter("level", "debug")
	line := `{"level":"info","msg":"hello"}`
	if !f.Allow(line) {
		t.Error("expected line to be allowed")
	}
}

func TestExcludeFilter_Allow_CaseInsensitive(t *testing.T) {
	f, _ := NewExcludeFilter("level", "DEBUG")
	line := `{"level":"debug","msg":"verbose"}`
	if f.Allow(line) {
		t.Error("expected case-insensitive match to exclude line")
	}
}

func TestExcludeFilter_Allow_MissingFieldAllows(t *testing.T) {
	f, _ := NewExcludeFilter("level", "debug")
	line := `{"msg":"no level field"}`
	if !f.Allow(line) {
		t.Error("expected line without field to be allowed")
	}
}

func TestExcludeFilter_Allow_InvalidJSONAllows(t *testing.T) {
	f, _ := NewExcludeFilter("level", "debug")
	if !f.Allow("not json") {
		t.Error("expected invalid JSON to be allowed")
	}
}

func TestExcludeFilter_Allow_EmptyLineAllows(t *testing.T) {
	f, _ := NewExcludeFilter("level", "debug")
	if !f.Allow("") {
		t.Error("expected empty line to be allowed")
	}
}
