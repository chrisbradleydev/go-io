package writer

import (
	"fmt"
	"io"
)

func WriteString(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, fmt.Errorf("error occured while writing: %w", err)
	}
	return n, nil
}
