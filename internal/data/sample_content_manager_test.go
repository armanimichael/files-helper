package data

import (
	"io/ioutil"
	"os"
	"testing"
)

func cleanTestFolder(folder string, t *testing.T) {
	if err := os.RemoveAll(folder); err != nil {
		t.Fatal(err)
	}
}

func Test_upsertSampleFolder(t *testing.T) {
	const testFolder = "sample_test"
	upsertSampleFolder(testFolder)

	if _, err := os.Stat(testFolder); err != nil {
		t.Fatal(err)
	}
	cleanTestFolder(testFolder, t)
}

func Test_CreateSampleHtmlFiles(t *testing.T) {
	cleanTestFolder(ContentFolder, t)
	paragraphsCount := 2
	CreateSampleHtmlFiles(paragraphsCount)

	subs, err := ioutil.ReadDir(ContentFolder)
	if err != nil {
		t.Fatal(err)
	}
	if subCount := len(subs); subCount != paragraphsCount {
		t.Fatalf("looking for %d files, found %d instead", paragraphsCount, subCount)
	}
	cleanTestFolder(ContentFolder, t)
}
