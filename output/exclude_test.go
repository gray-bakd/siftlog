package output

import (
	"encoding/json"
	"testing"
)

func TestExcludeFields_EmptyLine(t *testing.T) {
	result := ExcludeFields("", ExcludeOptions{Fields: []string{"foo"}})
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestExcludeFields_NoFields(t *testing.T) {
	line := `{"level":"info","msg":"hello"}`
	result := ExcludeFields(line, ExcludeOptions{})
	if result != line {
		t.Errorf("expected unchanged line, got %q", result)
	}
}

func TestExcludeFields_InvalidJSON(t *testing.T) {
	line := "not json"
	result := ExcludeFields(line, ExcludeOptions{Fields: []string{"foo"}})
	if result != line {
		t.Errorf("expected original line, got %q", result)
	}
}

func TestExcludeFields_RemovesSingleField(t *testing.T) {
	line := `{"level":"info","msg":"hello","secret":"password"}`
	result := ExcludeFields(line, ExcludeOptions{Fields: []string{"secret"}})

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(result), &obj); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if _, ok := obj["secret"]; ok {
		t.Error("expected 'secret' field to be removed")
	}
	if obj["level"] != "info" {
		t.Error("expected 'level' field to remain")
	}
}

func TestExcludeFields_RemovesMultipleFields(t *testing.T) {
	line := `{"level":"info","msg":"hello","token":"abc","password":"xyz"}`
	result := ExcludeFields(line, ExcludeOptions{Fields: []string{"token", "password"}})

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(result), &obj); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if _, ok := obj["token"]; ok {
		t.Error("expected 'token' to be removed")
	}
	if _, ok := obj["password"]; ok {
		t.Error("expected 'password' to be removed")
	}
	if obj["msg"] != "hello" {
		t.Error("expected 'msg' to remain")
	}
}

func TestExcludeFields_NonExistentField(t *testing.T) {
	line := `{"level":"info","msg":"hello"}`
	result := ExcludeFields(line, ExcludeOptions{Fields: []string{"nonexistent"}})

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(result), &obj); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if len(obj) != 2 {
		t.Errorf("expected 2 fields, got %d", len(obj))
	}
}
