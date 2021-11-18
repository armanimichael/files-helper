package data

import (
	"fmt"
	"log"
	"os"
	"path"
)

type contentGenerator func(int) []byte

const contentFolder = "./_sample-data"

func upsertSampleFolder(folder string) {
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		if err := os.Mkdir(folder, os.ModeDir); err != nil {
			log.Fatal(err)
		}
	}
}

func createSampleFiles(maxParagraphs int, nameFormat string, content contentGenerator) {
	upsertSampleFolder(contentFolder)
	for i := 1; i <= maxParagraphs; i++ {
		bytes := content(i)
		filename := fmt.Sprintf(nameFormat, i)
		file := createSampleFile(filename)
		populateSampleFile(file, bytes)
	}
}

func CreateSamplePlaintextFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, "plain-%02d.txt", GetSamplePlaintext)
}

func createSampleFile(filename string) *os.File {
	filePath := path.Join(contentFolder, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func populateSampleFile(file *os.File, content []byte) {
	if _, err := file.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
