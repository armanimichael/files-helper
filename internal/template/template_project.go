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
const testFolder string = "./_test_proj"

var wg sync.WaitGroup

func getSampleFileExtensions() []string {
	return []string{"txt", "html", "css", ""}
}

func getRandomExtension() string {
	extensions := getSampleFileExtensions()
	count := len(extensions)
	i := rand.Intn(count)
	return extensions[i]
}

func GenerateTestProject() {
	os.RemoveAll(testFolder)
	for i := 0; i < rootFolerCount; i++ {
		wg.Add(1)
		go generateSubDirectoryWithFiles(i)
	}
	wg.Wait()
}

func generateSubDirectoryWithFiles(i int) {
	defer wg.Done()
	dir := path.Join(testFolder, fmt.Sprintf("%02d", i))
	generateTestDirectories(dir)
	generateTestFiles(dir)
}

func generateTestDirectories(dir string) {
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func generateTestFiles(dir string) {
	filesCount := rand.Intn(11)
	for i := 0; i < filesCount; i++ {
		extension := getRandomExtension()
		filename := fmt.Sprintf("%02d.%s", i, extension)
		file := path.Join(dir, filename)

		out, err := os.Create(file)
		if err != nil {
			continue
		}

		paragraphs := 2
		out.Write(data.GetSamplePlaintext(paragraphs))
		out.Close()
	}
}
