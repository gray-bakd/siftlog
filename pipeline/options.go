package pipeline

import (
	"errors"
	"io"

	"github.com/user/siftlog/output"
)

// Options holds configuration for a pipeline run.
type Options struct {
	Input      io.Reader
	Output     io.Writer
	Queries    []string
	MinLevel   string
	FormatMode string
	Color      bool
	Highlight  []string
}

// Validate checks that required fields are set and values are valid.
func (o *Options) Validate() error {
	if o.Input == nil {
		return errors.New("input reader is required")
	}
	if o.Output == nil {
		return errors.New("output writer is required")
	}
	if o.FormatMode == "" {
		o.FormatMode = "raw"
	}
	if _, err := output.ParseFormatMode(o.FormatMode); err != nil {
		return err
	}
	if o.MinLevel != "" {
		if _, err := output.ParseLevel(o.MinLevel); err != nil {
			return err
		}
	}
	return nil
}
