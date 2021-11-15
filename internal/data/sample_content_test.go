package data

import (
	"os"
	"path"
	"testing"
)

func Test_GetSamplePlaintext(t *testing.T) {
	result := GetSamplePlaintext(2)
	if len(result) <= 0 {
		t.Fatal("no content found in api")
	}

	folder := "samples"
	upsertSampleFolder(folder)
	file, err := os.Create(path.Join(folder, "plaintext.txt"))
	if err != nil {
		t.Fatal(err)
	}

	_, err = file.WriteAt(result, 0)
	if err != nil {
		t.Fatal(err)
	}

	file.Close()
	if err := os.RemoveAll(folder); err != nil {
		t.Fatal(err)
	}
}
