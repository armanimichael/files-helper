package backup

import (
	"compress/gzip"
	"fmt"
	"github.com/armanimichael/files-helper/pkg/compression"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

const DirnameSuffix = "_backup"

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

func getBackupFilename(filename string) string {
	const bkExtension = "gz"
	currentTime := fmt.Sprint(time.Now().UnixNano())

	return fmt.Sprintf("%s-%s.%s", currentTime, filepath.Base(filename), bkExtension)
}

func getBackupFile(file *os.File) string {
	bkPath := getRootPath(file.Name())
	bkFilename := getBackupFilename(file.Name())
	bkFile := path.Join(bkPath+DirnameSuffix, bkFilename)

	if err := os.MkdirAll(filepath.Dir(bkFile), os.ModeDir); err != nil {
		log.Fatal(err)
	}
	return bkFile
}

// GzipFile creates a gzipped copy of the file
// with the same folder structure as the original file
func GzipFile(file *os.File) {
	header := createHeader(file)
	bkFile := getBackupFile(file)

	target, err := os.Create(bkFile)
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	if err := compression.Gzip(header, target, file); err != nil {
		log.Fatal(err)
	}
}
