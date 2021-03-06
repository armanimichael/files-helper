package template

import (
	"fmt"
	"github.com/armanimichael/files-helper/internal/data"
	"log"
	"math/rand"
	"os"
	"path"
	"sync"
)

const rootFolerCount = 100
const ContentFolder string = "./_test_proj"

var wg sync.WaitGroup
var maxParagraphs int

func GenerateTestProject(paragraphs int) {
	maxParagraphs = paragraphs
	os.RemoveAll(ContentFolder)
	for i := 0; i < rootFolerCount; i++ {
		wg.Add(1)
		go generateSubDirectoryWithFiles(i)
	}
	wg.Wait()
}

func getSampleContent(extension string, paragraphs int) []byte {
	if extension == "html" {
		content, err := data.ReadSampleHtml(paragraphs)
		if err != nil {
			log.Fatal(err)
		}
		return content
	}

	content, err := data.ReadSamplePlaintext(paragraphs)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func getSampleFileExtensions() []string {
	return []string{"txt", "html", "css", ""}
}

func getRandomExtension() string {
	extensions := getSampleFileExtensions()
	count := len(extensions)
	i := rand.Intn(count)
	return extensions[i]
}

func generateSubDirectoryWithFiles(i int) {
	defer wg.Done()
	dir := path.Join(ContentFolder, fmt.Sprintf("%02d", i))
	generateTestDirectories(dir)
	generateTestFiles(dir)
}

func generateTestDirectories(dir string) {
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func generateTestFiles(dir string) {
	filesCount := rand.Intn(maxParagraphs + 1)
	for i := 0; i < filesCount; i++ {
		extension := getRandomExtension()
		filename := fmt.Sprintf("%02d.%s", i, extension)
		file := path.Join(dir, filename)

		out, err := os.Create(file)
		if err != nil {
			continue
		}

		paragraphs := rand.Intn(maxParagraphs) + 1
		content := getSampleContent(extension, paragraphs)
		out.Write(content)
		out.Close()
	}
}
