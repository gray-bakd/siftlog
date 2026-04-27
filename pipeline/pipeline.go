package pipeline

import (
	"fmt"
	"io"

	"github.com/user/siftlog/filter"
	"github.com/user/siftlog/input"
	"github.com/user/siftlog/output"
)

// Options holds configuration for a pipeline run.
type Options struct {
	Queries  []string
	NoColor  bool
	Out      io.Writer
	ErrOut   io.Writer
}

// Run reads JSON log lines from r, filters them by queries, colorizes, and writes to opts.Out.
func Run(r io.Reader, opts Options) error {
	queries, err := filter.ParseQueries(opts.Queries)
	if err != nil {
		return fmt.Errorf("invalid query: %w", err)
	}

	reader := input.NewLineReader(r)
	for {
		line, ok := reader.Next()
		if !ok {
			break
		}
		if line == "" {
			continue
		}
		if len(queries) > 0 && !filter.Match(line, queries) {
			continue
		}
		formatted := output.Colorize(line, !opts.NoColor)
		fmt.Fprintln(opts.Out, formatted)
	}

	if err := reader.Err(); err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	return nil
}
