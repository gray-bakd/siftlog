package pipeline

import (
	"fmt"
	"io"

	"github.com/user/siftlog/filter"
	"github.com/user/siftlog/output"
)

// Options holds all configuration for a pipeline run.
type Options struct {
	Input       io.Reader
	Output      io.Writer
	Queries     []string
	MinLevel    string
	FormatMode  output.FormatMode
	NoColor     bool
}

// Validate checks that the Options are complete and consistent.
// It returns an error if required fields are missing or values are invalid.
func (o *Options) Validate() error {
	if o.Input == nil {
		return fmt.Errorf("input reader must not be nil")
	}
	if o.Output == nil {
		return fmt.Errorf("output writer must not be nil")
	}
	if o.MinLevel != "" {
		if _, err := filter.NewLevelFilter(o.MinLevel); err != nil {
			return fmt.Errorf("invalid min-level %q: %w", o.MinLevel, err)
		}
	}
	if _, err := filter.ParseQueries(o.Queries); err != nil {
		return fmt.Errorf("invalid query: %w", err)
	}
	return nil
}
