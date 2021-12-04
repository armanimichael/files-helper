package cmd

import (
	"github.com/armanimichael/files-helper/cmd/backup"
	"github.com/armanimichael/files-helper/pkg/util"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Opts struct {
	Root          string
	Extensions    []string
	SearchPattern string
	Replace       string
	LogFile       bool
	Backup        bool
}

type FileHandler func(file *os.File, currentPath string, opts Opts)

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
func ReadFile(currentPath string) *os.File {
	file, err := os.OpenFile(currentPath, os.O_RDONLY, os.ModeType)
	PathFatal(currentPath, err)
	return file
}

func resetFilePointer(file *os.File) {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}
}

// HandleFoundFiles walks all files under a root directory
// and allow working with the files respecting the filters
// automatically closing the file once done
func HandleFoundFiles(opts Opts, handler FileHandler) {
	filepath.WalkDir(opts.Root, func(currentPath string, d fs.DirEntry, err error) error {
		PathFatal(currentPath, err)
		if !IsSupportedPath(d, currentPath, opts.Extensions) {
			return nil
		}
		file := ReadFile(currentPath)

		// Handle found content
		found, err := util.IsInReader(file, opts.SearchPattern)
		resetFilePointer(file)
		PathFatal(currentPath, err)
		if found {
			if opts.Backup {
				backup.GzipFile(file)
				resetFilePointer(file)
			}

			handler(file, currentPath, opts)
		}

		file.Close()
		return nil
	})
}
