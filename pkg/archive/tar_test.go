package archive

import (
	"log"
	"os"
	"testing"
)

func createTarTestFile() *os.File {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func tarTestFile(t *testing.T, file *os.File) {
	if err := TarFolder(file.Name()); err != nil {
		t.Fatal(err)
	}
}

func deleteTestFiles(file *os.File) {
	file.Close()
	os.Remove(file.Name())
	os.Remove(file.Name() + ".tar")
}

func Test_TarFolder(t *testing.T) {
	file := createTarTestFile()
	tarTestFile(t, file)
	deleteTestFiles(file)
}

func Test_UnTarFolder(t *testing.T) {
	file := createTarTestFile()
	tarTestFile(t, file)

	untarredDir := "untarred"
	os.Mkdir(untarredDir, os.ModeDir)
	if err := UntarFolder(file.Name()+".tar", untarredDir); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(untarredDir); err != nil {
		log.Fatal(err)
	}
	deleteTestFiles(file)
}
