package output

import (
	"fmt"
	"io"
)

// Writer wraps an io.Writer and provides structured log line output
// with optional formatting and colorization.
type Writer struct {
	out      io.Writer
	format   FormatMode
	noColor  bool
}

// NewWriter creates a new Writer with the given output destination,
// format mode, and color preference.
func NewWriter(out io.Writer, format FormatMode, noColor bool) *Writer {
	return &Writer{
		out:     out,
		format:  format,
		noColor: noColor,
	}
}

// WriteLine formats and writes a single log line to the underlying writer.
// It applies formatting first, then colorization, and appends a newline.
func (w *Writer) WriteLine(line string) error {
	formatted := Format(line, w.format)
	colored := Colorize(formatted, !w.noColor)
	_, err := fmt.Fprintln(w.out, colored)
	return err
}
