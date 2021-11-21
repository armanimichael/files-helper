package find

import (
	"fmt"
	"github.com/armanimichael/files-helper/pkg/util"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func filterExtension(extension string, allowedExtensions []string) bool {
	for _, ext := range allowedExtensions {
		if extension == ("." + ext) {
			return true
		}
	}
	return false
}

func SearchInFiles(root string, extensions []string, searchPattern string) {
	filepath.WalkDir(root, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf("%s: %s", currentPath, err)
		}
		if d.IsDir() || !filterExtension(path.Ext(currentPath), extensions) {
			return nil
		}

		file, err := os.OpenFile(currentPath, os.O_RDONLY, os.ModeType)
		if err != nil {
			log.Fatalf("%s: %s", currentPath, err)
		}
		defer file.Close()

		found, err := util.IsInReader(file, searchPattern)
		if err != nil {
			log.Fatalf("%s: %s", currentPath, err)
		}
		if found {
			fmt.Println(currentPath)
		}

		return nil
	})
}
