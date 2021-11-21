package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const ContentFolder = "./_sample_data"
const plainFileNameFormat = "plain-%02d.txt"
const htmlFileNameFormat = "html-%02d.txt"

func readSampleContent(paragraphs int, fileNameFormat string) (content []byte, err error) {
	fileName := path.Join(ContentFolder, fmt.Sprintf(fileNameFormat, paragraphs))
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
