package data

import (
	"io/ioutil"
	"testing"
)

func Test_CreateSamplePlaintextFiles(t *testing.T) {
	paragraphsCount := 2
	CreateSamplePlaintextFiles(paragraphsCount)

	subs, err := ioutil.ReadDir(ContentFolder)
	if err != nil {
		t.Fatal(err)
	}
	if subCount := len(subs); subCount != paragraphsCount {
		t.Fatalf("looking for %d files, found %d instead", paragraphsCount, subCount)
	}
	cleanTestFolder(ContentFolder, t)
}

func Test_ReadSamplePlaintext(t *testing.T) {
	cleanTestFolder(ContentFolder, t)
	paragraphs := 2
	CreateSamplePlaintextFiles(paragraphs)

	if _, err := ReadSamplePlaintext(2); err != nil {
		t.Fatal(err)
	}

	cleanTestFolder(ContentFolder, t)
}
