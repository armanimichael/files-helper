package backup

import (
	"compress/gzip"
	"github.com/armanimichael/files-helper/pkg/compression"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

func getRootPath(filename string) string {
	upper := filepath.Dir(filename)
	if upper == "." {
		return filename
	}
	return getRootPath(upper)
}

func createHeader(file *os.File) gzip.Header {
	return gzip.Header{
		Name:    file.Name(),
		ModTime: time.Now(),
	}
}

func getBackupFilename(file *os.File) string {
	const bkExtension = ".gz"

	bkPath := getRootPath(file.Name())
	bkFilename := filepath.Base(file.Name()) + bkExtension
	bkFile := path.Join(bkPath+"_backup", bkFilename)

	if err := os.MkdirAll(filepath.Dir(bkFile), os.ModeDir); err != nil {
		log.Fatal(err)
	}
	return bkFile
}

// GzipFile creates a gzipped copy of the file
// with the same folder structure as the original file
func GzipFile(file *os.File) {
	header := createHeader(file)
	bkFile := getBackupFilename(file)

	target, err := os.Create(bkFile)
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	if err := compression.Gzip(header, target, file); err != nil {
		log.Fatal(err)
	}
}
