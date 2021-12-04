package backup

import (
	"os"
	"testing"
)

func Test_GzipFile(t *testing.T) {
	const filename = "test"
	testFile, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile.Name())
	defer os.RemoveAll(filename + "_backup")

	testFile.Write([]byte("test"))
	if _, err := testFile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	GzipFile(testFile)
	testFile.Close()
}
