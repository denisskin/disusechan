package disusechan

import (
	"bufio"
	"bytes"
	"io"
)

func FetchLines(f io.Reader, fn func(line []byte) error) error {
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadSlice('\n')
		eof := err == io.EOF
		if eof || err == nil {
			err = fn(bytes.TrimSpace(line))
		}
		if eof || err != nil {
			return err
		}
	}
}

func FetchCSV(f io.Reader, fn func(values [][]byte) error) error {
	return FetchLines(f, func(line []byte) error {
		return fn(csvValues(line))
	})
}
