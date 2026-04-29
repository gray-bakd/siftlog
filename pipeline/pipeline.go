package pipeline

import (
	"fmt"

	"github.com/user/siftlog/filter"
	"github.com/user/siftlog/input"
	"github.com/user/siftlog/output"
)

// Run executes the full log processing pipeline using the given Options.
func Run(opts Options) error {
	if err := opts.Validate(); err != nil {
		return fmt.Errorf("invalid options: %w", err)
	}

	var filters []filter.Filter

	if len(opts.Queries) > 0 {
		queries, err := filter.ParseQueries(opts.Queries)
		if err != nil {
			return fmt.Errorf("invalid query: %w", err)
		}
		filters = append(filters, filter.NewFieldFilter(queries))
	}

	if opts.MinLevel != "" {
		lvl, err := output.ParseLevel(opts.MinLevel)
		if err != nil {
			return fmt.Errorf("invalid level: %w", err)
		}
		filters = append(filters, filter.NewLevelFilter(lvl))
	}

	fmtMode, _ := output.ParseFormatMode(opts.FormatMode)
	writer := output.NewWriter(opts.Output, fmtMode, opts.Color)
	reader := input.NewLineReader(opts.Input)

	for reader.Next() {
		line := reader.Line()
		if line == "" {
			continue
		}

		allowed := true
		for _, f := range filters {
			if !f.Allow(line) {
				allowed = false
				break
			}
		}
		if !allowed {
			continue
		}

		if len(opts.Highlight) > 0 {
			line = output.HighlightFields(line, opts.Highlight, opts.Color)
			fmt.Fprintln(opts.Output, line)
			continue
		}

		writer.WriteLine(line)
	}

	return reader.Err()
}
