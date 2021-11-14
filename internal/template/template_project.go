package template

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
)

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
		file := path.Join(dir, fmt.Sprintf("%d.txt", i))
		os.Create(file)
	}
}
