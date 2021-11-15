package data

import (
	"os"
	"testing"
)

func Test_GetSamplePlaintext(t *testing.T) {
	result := GetSamplePlaintext(2)
	if len(result) <= 0 {
		t.Fatal("no content found in api")
	}

	file, err := os.Create("./samples/plaintext.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteAt(result, 0)
	if err != nil {
		t.Fatal(err)
	}
}
