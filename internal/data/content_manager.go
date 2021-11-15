package data

import (
	"fmt"
	"log"
	"os"
	"path"
)

const contentFolder = "./_sample-data"

func upsertSampleFolder(folder string) {
	_, err := os.Stat(folder)
	if os.IsNotExist(err) {
		if err := os.Mkdir(folder, os.ModeDir); err != nil {
			log.Fatal(err)
		}
	}
}

func CreateSamplePlaintextFiles(maxParagraphs int) {
	upsertSampleFolder(contentFolder)
	for i := 1; i <= maxParagraphs; i++ {
		content := GetSamplePlaintext(i)
		filename := fmt.Sprintf("plain-%2d.txt", i)
		file := CreateSampleFile(filename)
		PopulateSampleFile(file, content)
	}
}

func CreateSampleFile(filename string) *os.File {
	filePath := path.Join(contentFolder, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func PopulateSampleFile(file *os.File, content []byte) {
	if _, err := file.Write(content); err != nil {
		log.Fatal(err)
	}
	file.Close()
}
