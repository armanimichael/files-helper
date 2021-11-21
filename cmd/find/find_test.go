package find

import (
	"github.com/armanimichael/files-helper/internal/data"
	"github.com/armanimichael/files-helper/internal/template"
	"os"
	"testing"
)

func createAndCleanTestFiles() func() {
	paragraphs := 4
	data.CreateSamplePlaintextFiles(paragraphs)
	data.CreateSampleHtmlFiles(paragraphs)
	template.GenerateTestProject(paragraphs)

	return func() {
		os.RemoveAll(data.ContentFolder)
		os.RemoveAll(template.ContentFolder)
	}
}

func Test_FindInFiles(t *testing.T) {
	defer createAndCleanTestFiles()()

	extensions := []string{"html"}
	SearchInFiles(template.ContentFolder, extensions, "<li>")
}
