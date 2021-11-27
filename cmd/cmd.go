package cmd

import (
	"io/fs"
	"log"
	"os"
	"path"
)

type Opts struct {
	Root          string
	Extensions    []string
	SearchPattern string
	LogFile       bool
}

func filterExtension(extension string, allowedExtensions []string) bool {
	for _, ext := range allowedExtensions {
		if extension == ("." + ext) {
			return true
		}
	}
	return false
}

// IsSupportedPath returns whether the current path is supported by the used function
func IsSupportedPath(dir fs.DirEntry, currentPath string, extensions []string) bool {
	isSupportedExt := filterExtension(path.Ext(currentPath), extensions)
	return !dir.IsDir() && isSupportedExt
}

// PathFatal logs and exits the program with an error if it finds any path related error
func PathFatal(currentPath string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", currentPath, err)
	}
}

// ReadFile returns a file and handle any error
func ReadFile(currentPath string, err error) *os.File {
	file, err := os.OpenFile(currentPath, os.O_RDONLY, os.ModeType)
	PathFatal(currentPath, err)
	return file
}
