package output

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTruncateFields_EmptyLine(t *testing.T) {
	result := TruncateFields("", DefaultTruncateOptions())
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestTruncateFields_InvalidJSON(t *testing.T) {
	line := "not json at all"
	result := TruncateFields(line, DefaultTruncateOptions())
	if result != line {
		t.Errorf("expected original line, got %q", result)
	}
}

func TestTruncateFields_NoTruncationNeeded(t *testing.T) {
	line := `{"msg":"hello","level":"info"}`
	result := TruncateFields(line, DefaultTruncateOptions())
	var orig, got map[string]interface{}
	json.Unmarshal([]byte(line), &orig)
	json.Unmarshal([]byte(result), &got)
	if got["msg"] != orig["msg"] {
		t.Errorf("msg should be unchanged")
	}
}

func TestTruncateFields_TruncatesLongValue(t *testing.T) {
	long := strings.Repeat("a", 200)
	line := `{"msg":"` + long + `"}`
	opts := TruncateOptions{MaxLength: 10, Suffix: "..."}
	result := TruncateFields(line, opts)

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(result), &obj); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	msg, _ := obj["msg"].(string)
	if !strings.HasSuffix(msg, "...") {
		t.Errorf("expected suffix '...', got %q", msg)
	}
	if len([]rune(msg)) != 13 { // 10 + len("...")
		t.Errorf("unexpected length %d", len(msg))
	}
}

func TestTruncateFields_ZeroMaxLengthDisabled(t *testing.T) {
	long := strings.Repeat("b", 300)
	line := `{"msg":"` + long + `"}`
	opts := TruncateOptions{MaxLength: 0, Suffix: "..."}
	result := TruncateFields(line, opts)
	if result != line {
		t.Errorf("expected unchanged line when MaxLength=0")
	}
}

func TestTruncateFields_NonStringValuesUntouched(t *testing.T) {
	line := `{"count":42,"active":true,"ratio":3.14}`
	opts := TruncateOptions{MaxLength: 2, Suffix: "~"}
	result := TruncateFields(line, opts)
	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(result), &obj); err != nil {
		t.Fatalf("result is not valid JSON: %v", err)
	}
	if obj["count"] != float64(42) {
		t.Errorf("numeric field should be unchanged")
	}
}
