package archive

import (
	"log"
	"os"
	"testing"
)

func Test_TarFolder(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err := TarFolder(file.Name()); err != nil {
		t.Fatal(err)
	}

	file.Close()
	os.Remove(file.Name())
	os.Remove(file.Name() + ".tar")
}
