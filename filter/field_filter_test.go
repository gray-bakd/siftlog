package filter

import (
	"testing"
)

func TestNewFieldFilter_Valid(t *testing.T) {
	f, err := NewFieldFilter("service=auth")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Field != "service" {
		t.Errorf("expected field %q, got %q", "service", f.Field)
	}
	if f.Pattern != "auth" {
		t.Errorf("expected pattern %q, got %q", "auth", f.Pattern)
	}
}

func TestNewFieldFilter_Invalid(t *testing.T) {
	cases := []string{
		"",
		"noequals",
		"=missingfield",
		"missingvalue=",
	}
	for _, c := range cases {
		_, err := NewFieldFilter(c)
		if err == nil {
			t.Errorf("expected error for input %q, got nil", c)
		}
	}
}

func TestFieldFilter_Allow_Match(t *testing.T) {
	f, _ := NewFieldFilter("service=auth")
	line := `{"service":"auth-service","level":"info"}`
	if !f.Allow(line) {
		t.Errorf("expected Allow to return true for matching line")
	}
}

func TestFieldFilter_Allow_CaseInsensitive(t *testing.T) {
	f, _ := NewFieldFilter("service=AUTH")
	line := `{"service":"auth-service"}`
	if !f.Allow(line) {
		t.Errorf("expected case-insensitive match")
	}
}

func TestFieldFilter_Allow_NoMatch(t *testing.T) {
	f, _ := NewFieldFilter("service=payments")
	line := `{"service":"auth-service","level":"info"}`
	if f.Allow(line) {
		t.Errorf("expected Allow to return false for non-matching line")
	}
}

func TestFieldFilter_Allow_MissingField(t *testing.T) {
	f, _ := NewFieldFilter("host=localhost")
	line := `{"service":"auth","level":"info"}`
	if f.Allow(line) {
		t.Errorf("expected Allow to return false when field is missing")
	}
}

func TestFieldFilter_Allow_InvalidJSON(t *testing.T) {
	f, _ := NewFieldFilter("service=auth")
	if f.Allow("not-json") {
		t.Errorf("expected Allow to return false for invalid JSON")
	}
}

func TestFieldFilter_Allow_EmptyLine(t *testing.T) {
	f, _ := NewFieldFilter("service=auth")
	if f.Allow("") {
		t.Errorf("expected Allow to return false for empty line")
	}
}
