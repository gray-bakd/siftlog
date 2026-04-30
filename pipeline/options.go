package pipeline

import (
	"errors"
	"io"

	"github.com/siftlog/output"
)

// Options holds all configuration for a pipeline run.
type Options struct {
	Input        io.Reader
	Output       io.Writer
	Queries      []string
	MinLevel     string
	ColorEnabled bool
	FormatMode   output.FormatMode
	Highlight    []string
	Truncate     output.TruncateOptions
	TruncateOn   bool
}

// Validate checks that required fields are set and values are coherent.
func (o *Options) Validate() error {
	if o.Input == nil {
		return errors.New("pipeline: Input must not be nil")
	}
	if o.Output == nil {
		return errors.New("pipeline: Output must not be nil")
	}
	if o.MinLevel != "" {
		if _, err := output.ParseLevel(o.MinLevel); err != nil {
			return err
		}
	}
	return nil
}
