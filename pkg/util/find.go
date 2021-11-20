package bytesHelpers

import (
	"bufio"
	"io"
	"strings"
)

// IsInReader takes a reader and a search pattern
// returns true and breaks the reading once it finds the given pattern.
func IsInReader(r io.Reader, searchPattern string) (err error, found bool) {
	reader := bufio.NewReader(r)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchPattern) && searchPattern != "" {
			return nil, true
		}
	}

	if err = scanner.Err(); err != nil {
		return err, false
	}
	return nil, false
}
