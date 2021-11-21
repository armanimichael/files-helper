package util

import (
	"bufio"
	"io"
	"strings"
)

// IsInReader takes a reader and a search pattern
// returns true and breaks the reading once it finds the given pattern.
func IsInReader(r io.Reader, searchPattern string) (found bool, err error) {
	reader := bufio.NewReader(r)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchPattern) && searchPattern != "" {
			return true, nil
		}
	}

	if err = scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}
