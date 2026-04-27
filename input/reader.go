package input

import (
	"bufio"
	"io"
)

// LineReader reads lines from an io.Reader one at a time.
type LineReader struct {
	scanner *bufio.Scanner
}

// NewLineReader creates a new LineReader wrapping the given reader.
func NewLineReader(r io.Reader) *LineReader {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	return &LineReader{scanner: scanner}
}

// Next advances to the next line and returns it.
// Returns (line, true) if a line was read, or ("", false) on EOF or error.
func (lr *LineReader) Next() (string, bool) {
	if lr.scanner.Scan() {
		return lr.scanner.Text(), true
	}
	return "", false
}

// Err returns any error encountered during scanning (excluding io.EOF).
func (lr *LineReader) Err() error {
	return lr.scanner.Err()
}
