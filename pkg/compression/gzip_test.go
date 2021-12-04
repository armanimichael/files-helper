package compression

import (
	"bytes"
	"compress/gzip"
	"reflect"
	"strings"
	"testing"
)

const testContent = "Test content"

var testHeader = gzip.Header{
	Name:    "Test",
	Comment: "Test",
}

var testSource = strings.NewReader(testContent)

func Test_Gzip(t *testing.T) {
	var testTarget bytes.Buffer

	if err := Gzip(testHeader, &testTarget, testSource); err != nil {
		t.Fatal(err)
	}
}

func Test_UnGzip(t *testing.T) {
	var testTarget bytes.Buffer
	var unzipContent bytes.Buffer

	Gzip(testHeader, &testTarget, testSource)
	if _, err := UnGzip(&unzipContent, &testTarget); err != nil {
		t.Fatal(err)
	}
}

func Test_Gzip_Content(t *testing.T) {
	testSource.Reset(testContent)
	var testTarget bytes.Buffer
	var unzipContent bytes.Buffer

	Gzip(testHeader, &testTarget, testSource)
	UnGzip(&unzipContent, &testTarget)

	content := string(unzipContent.Bytes())
	if content != testContent {
		t.Errorf("unzipped content (%s) is different from the original content (%s)", content, testContent)
	}
}

func Test_Gzip_Header(t *testing.T) {
	var testTarget bytes.Buffer
	var unzipContent bytes.Buffer

	Gzip(testHeader, &testTarget, testSource)
	header, _ := UnGzip(&unzipContent, &testTarget)

	if !reflect.DeepEqual(header, testHeader) {
		t.Errorf("unzipped header (%v) is different from the original header (%v)", header, testHeader)
	}
}
