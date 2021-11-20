package bytesHelpers

import (
	"bytes"
	"testing"
)

func testIsInReader(content string, searchpattern string) (err error, found bool) {
	contentBytes := []byte(content)
	reader := bytes.NewReader(contentBytes)
	err, found = IsInReader(reader, searchpattern)
	if err != nil {
		return err, found
	}
	return nil, found
}

func Test_IsInReader_Find(t *testing.T) {
	content := "First line\nSecond line\nThird line"
	err, found := testIsInReader(content, "Second")
	if err != nil {
		t.Error(err)
	}
	if !found {
		t.Error("it should find the substring")
	}
}

func Test_IsInReader_NoFind(t *testing.T) {
	content := "First line\nSecond line\nThird line"
	err, found := testIsInReader(content, "test")
	if err != nil {
		t.Error(err)
	}
	if found {
		t.Error("it shouldn't find the substring")
	}
}

func Test_IsInReader_CaseSensitive(t *testing.T) {
	content := "First line\nSecond line\nThird line"
	err, found := testIsInReader(content, "second")
	if err != nil {
		t.Error(err)
	}
	if found {
		t.Error("it shouldn't find the substring")
	}
}
