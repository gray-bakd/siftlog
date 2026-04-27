package filter_test

import (
	"testing"

	"github.com/user/siftlog/filter"
)

func TestParseQueries_Valid(t *testing.T) {
	queries, err := filter.ParseQueries([]string{"level=error", "service=api"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(queries) != 2 {
		t.Fatalf("expected 2 queries, got %d", len(queries))
	}
	if queries[0].Field != "level" || queries[0].Value != "error" {
		t.Errorf("unexpected query[0]: %+v", queries[0])
	}
}

func TestParseQueries_Invalid(t *testing.T) {
	_, err := filter.ParseQueries([]string{"noequalssign"})
	if err == nil {
		t.Fatal("expected error for missing '=', got nil")
	}
}

func TestMatch_AllFieldsMatch(t *testing.T) {
	line := `{"level":"error","service":"api","msg":"timeout"}`
	queries := []filter.Query{{Field: "level", Value: "error"}, {Field: "service", Value: "api"}}
	if !filter.Match(line, queries) {
		t.Error("expected match, got no match")
	}
}

func TestMatch_FieldMissing(t *testing.T) {
	line := `{"level":"info","msg":"ok"}`
	queries := []filter.Query{{Field: "service", Value: "api"}}
	if filter.Match(line, queries) {
		t.Error("expected no match when field is missing")
	}
}

func TestMatch_ValueMismatch(t *testing.T) {
	line := `{"level":"info"}`
	queries := []filter.Query{{Field: "level", Value: "error"}}
	if filter.Match(line, queries) {
		t.Error("expected no match on value mismatch")
	}
}

func TestMatch_InvalidJSON(t *testing.T) {
	line := `not json at all`
	queries := []filter.Query{{Field: "level", Value: "error"}}
	if filter.Match(line, queries) {
		t.Error("expected no match for invalid JSON")
	}
}
