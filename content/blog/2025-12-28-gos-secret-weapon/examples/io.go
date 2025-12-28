package examples

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// ProcessLogs processes logs from reader to writer using io.Reader and io.Writer.
func ProcessLogs(w io.Writer, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			if _, err := fmt.Fprintln(w, line); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

// WriteLabel writes a label using io.StringWriter for efficient string writing.
func WriteLabel(w io.StringWriter, label string) error {
	_, err := w.WriteString("[" + label + "]")
	return err
}

// StreamSize gets the size of a seekable stream using io.Seeker.
func StreamSize(s io.Seeker) (int64, error) {
	current, err := s.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	size, err := s.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	_, err = s.Seek(current, io.SeekStart)
	return size, err
}

// ReadChunk reads a specific chunk from ReaderAt for random access reads.
func ReadChunk(r io.ReaderAt, offset int64, size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := r.ReadAt(buf, offset)
	return buf, err
}

// Transfer uses io.WriterTo for efficient copying when source is WriterTo.
func Transfer(dst io.Writer, src io.WriterTo) (int64, error) {
	return src.WriteTo(dst)
}

// ============================================================================
// IMPLEMENTING INTERFACES
// ============================================================================

// AReader is an infinite stream of 'A's that implements io.Reader.
type AReader struct{}

func (AReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
