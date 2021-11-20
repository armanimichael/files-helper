package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func CreateSamplePlaintextFiles(maxParagraphs int) {
	createSampleFiles(maxParagraphs, plainFileNameFormat, GetSamplePlaintext)
}

func ReadSamplePlaintext(paragraphs int) (content []byte, err error) {
	fileName := path.Join(ContentFolder, fmt.Sprintf(plainFileNameFormat, paragraphs))
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		return []byte{}, fmt.Errorf("file %v does not exist", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	content, err = ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}
