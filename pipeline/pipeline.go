package pipeline

import (
	"fmt"

	"github.com/siftlog/filter"
	"github.com/siftlog/input"
	"github.com/siftlog/output"
)

// Run executes the full log processing pipeline according to opts.
func Run(opts Options) error {
	if err := opts.Validate(); err != nil {
		return err
	}

	// Build composite filter.
	var filters []filter.Filter

	if len(opts.Queries) > 0 {
		queries, err := filter.ParseQueries(opts.Queries)
		if err != nil {
			return fmt.Errorf("pipeline: invalid query: %w", err)
		}
		filters = append(filters, filter.NewFieldFilter(queries))
	}

	if opts.MinLevel != "" {
		lf, err := filter.NewLevelFilter(opts.MinLevel)
		if err != nil {
			return fmt.Errorf("pipeline: invalid level: %w", err)
		}
		filters = append(filters, lf)
	}

	composite := filter.NewCompositeFilter(filters...)

	// Build writer.
	w := output.NewWriter(opts.Output, opts.FormatMode, opts.ColorEnabled)

	// Process lines.
	reader := input.NewLineReader(opts.Input)
	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			continue
		}
		if !composite.Allow(line) {
			continue
		}
		if len(opts.Highlight) > 0 {
			line = output.HighlightFields(line, opts.Highlight, opts.ColorEnabled)
		}
		if opts.TruncateOn {
			line = output.TruncateFields(line, opts.Truncate)
		}
		if err := w.WriteLine(line); err != nil {
			return fmt.Errorf("pipeline: write error: %w", err)
		}
	}
	return reader.Err()
}
