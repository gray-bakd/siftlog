package filter

import (
	"testing"
)

// alwaysAllow is a test stub that allows every line.
type alwaysAllow struct{}

func (a *alwaysAllow) Allow(_ []byte) bool { return true }

// alwaysDeny is a test stub that denies every line.
type alwaysDeny struct{}

func (a *alwaysDeny) Allow(_ []byte) bool { return false }

func TestNewCompositeFilter_NilFiltersIgnored(t *testing.T) {
	cf := NewCompositeFilter(nil, nil)
	if cf.Len() != 0 {
		t.Errorf("expected 0 active filters, got %d", cf.Len())
	}
}

func TestCompositeFilter_NoFilters_AllowsAll(t *testing.T) {
	cf := NewCompositeFilter()
	if !cf.Allow([]byte(`{"level":"info"}`)) {
		t.Error("expected empty composite to allow all lines")
	}
}

func TestCompositeFilter_AllAllow(t *testing.T) {
	cf := NewCompositeFilter(&alwaysAllow{}, &alwaysAllow{})
	if !cf.Allow([]byte(`{"msg":"hello"}`)) {
		t.Error("expected all-allow composite to allow line")
	}
}

func TestCompositeFilter_OneDenies(t *testing.T) {
	cf := NewCompositeFilter(&alwaysAllow{}, &alwaysDeny{})
	if cf.Allow([]byte(`{"msg":"hello"}`)) {
		t.Error("expected composite to deny line when one filter denies")
	}
}

func TestCompositeFilter_AllDeny(t *testing.T) {
	cf := NewCompositeFilter(&alwaysDeny{}, &alwaysDeny{})
	if cf.Allow([]byte(`{"msg":"hello"}`)) {
		t.Error("expected composite to deny line when all filters deny")
	}
}

func TestCompositeFilter_WithRealFilters(t *testing.T) {
	lf, err := NewLevelFilter("warn")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ff, err := NewFieldFilter("service=auth")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cf := NewCompositeFilter(lf, ff)

	allowed := []byte(`{"level":"error","service":"auth","msg":"failed"}`)
	if !cf.Allow(allowed) {
		t.Error("expected composite to allow matching line")
	}

	wrongLevel := []byte(`{"level":"debug","service":"auth","msg":"ok"}`)
	if cf.Allow(wrongLevel) {
		t.Error("expected composite to deny line with level below warn")
	}

	wrongService := []byte(`{"level":"error","service":"web","msg":"failed"}`)
	if cf.Allow(wrongService) {
		t.Error("expected composite to deny line with wrong service")
	}
}

func TestCompositeFilter_Len(t *testing.T) {
	cf := NewCompositeFilter(&alwaysAllow{}, nil, &alwaysDeny{})
	if cf.Len() != 2 {
		t.Errorf("expected Len 2, got %d", cf.Len())
	}
}
