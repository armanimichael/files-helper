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

func GzipFile(file *os.File) {
	header := gzip.Header{
		Name:    file.Name(),
		ModTime: time.Now(),
	}

	bkPath := getRootPath(file.Name())
	bkFilename := filepath.Base(file.Name()) + ".gz"
	bkFile := path.Join(bkPath+"_backup", bkFilename)
	os.MkdirAll(filepath.Dir(bkFile), os.ModeDir)

	target, err := os.Create(bkFile)
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	if err := compression.Gzip(header, target, file); err != nil {
		log.Fatal(err)
	}
}
