package util

import (
	"bytes"
	"strings"
	"testing"
)

func testIsInReader(content string, searchpattern string) (err error, found bool) {
	contentBytes := []byte(content)
	reader := bytes.NewReader(contentBytes)
	found, err = IsInReader(reader, searchpattern)
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

func Test_Replace(t *testing.T) {
	content := "ABC"
	expected := "AAC"
	match := "B"
	replace := "A"
	r := strings.NewReader(content)
	actual, err := Replace(r, match, replace)
	if err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Fatalf("%s should be equal to %s", string(actual), expected)
	}
}
