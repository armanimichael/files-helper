package util

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

// IsInReader takes a reader and a search pattern
// returns true and breaks the reading once it finds the given pattern.
func IsInReader(r io.Reader, searchPattern string) (found bool, err error) {
	scanner := bufio.NewScanner(r)

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

func Replace(r io.Reader, old string, new string) (newContent []byte, err error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}

	newContent = bytes.Replace(content, []byte(old), []byte(new), -1)
	return newContent, nil
}
