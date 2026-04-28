package filter

import (
	"testing"

	"github.com/user/siftlog/output"
)

func TestNewLevelFilter_Valid(t *testing.T) {
	f, err := NewLevelFilter("warn")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.Min != output.LevelWarn {
		t.Errorf("expected LevelWarn, got %v", f.Min)
	}
}

func TestNewLevelFilter_Invalid(t *testing.T) {
	_, err := NewLevelFilter("verbose")
	if err == nil {
		t.Fatal("expected error for unknown level, got nil")
	}
}

func TestLevelFilter_Allow(t *testing.T) {
	f, _ := NewLevelFilter("warn")

	cases := []struct {
		name    string
		line    string
		expect  bool
	}{
		{
			name:   "debug below threshold",
			line:   `{"level":"debug","msg":"verbose"}`,
			expect: false,
		},
		{
			name:   "info below threshold",
			line:   `{"level":"info","msg":"hello"}`,
			expect: false,
		},
		{
			name:   "warn at threshold",
			line:   `{"level":"warn","msg":"careful"}`,
			expect: true,
		},
		{
			name:   "error above threshold",
			line:   `{"level":"error","msg":"boom"}`,
			expect: true,
		},
		{
			name:   "fatal above threshold",
			line:   `{"level":"fatal","msg":"dead"}`,
			expect: true,
		},
		{
			name:   "alternate lvl field",
			line:   `{"lvl":"error","msg":"alt"}`,
			expect: true,
		},
		{
			name:   "no level field passes through",
			line:   `{"msg":"no level here"}`,
			expect: true,
		},
		{
			name:   "non-JSON passes through",
			line:   "plain text log line",
			expect: true,
		},
		{
			name:   "empty line rejected",
			line:   "",
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := f.Allow(tc.line)
			if got != tc.expect {
				t.Errorf("Allow(%q) = %v, want %v", tc.line, got, tc.expect)
			}
		})
	}
}
