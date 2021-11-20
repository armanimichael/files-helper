package template

import (
	"github.com/armanimichael/files-helper/internal/data"
	"os"
	"testing"
)

func cleanTestFolder(folder string, t *testing.T) {
	if err := os.RemoveAll(folder); err != nil {
		t.Fatal(err)
	}
}

func Test_GenerateTestProject(t *testing.T) {
	data.CreateSamplePlaintextFiles(10)
	GenerateTestProject(testFolder)
	cleanTestFolder(testFolder, t)
	cleanTestFolder(data.ContentFolder, t)
}
