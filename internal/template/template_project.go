package template

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
)

func getSampleFileExtensions() []string {
	return []string{"txt", "html", "css", ""}
}

func getRandomExtension() string {
	extensions := getSampleFileExtensions()
	count := len(extensions)
	i := rand.Intn(count)
	return extensions[i]
}

func generateTestProject() {
	const testFolder string = "./_test_proj"
	os.RemoveAll(testFolder)
	for i := 0; i < 100; i++ {
		dir := path.Join(testFolder, fmt.Sprintf("%d", i))
		generateTestDirectories(dir)
		generateTestFiles(dir)
	}
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
		filename := fmt.Sprintf("%d.%s", i, extension)
		file := path.Join(dir, filename)
		os.Create(file)
	}
}
