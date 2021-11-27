package replace

import (
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/internal/data"
	"github.com/armanimichael/files-helper/internal/template"
	"log"
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
		if err := os.RemoveAll(template.ContentFolder); err != nil {
			log.Fatal(err)
		}
	}
}

func Test_SubstituteInFiles(t *testing.T) {
	defer createAndCleanTestFiles()()

	opts := cmd.Opts{
		Root:          template.ContentFolder,
		Extensions:    []string{"html"},
		SearchPattern: "<li>",
		Replace:       "<pizza>",
		LogFile:       false,
	}
	SubstituteInFiles(opts)
}
