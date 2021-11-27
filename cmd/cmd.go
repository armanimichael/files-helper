package cmd

import "log"

func FilterExtension(extension string, allowedExtensions []string) bool {
	for _, ext := range allowedExtensions {
		if extension == ("." + ext) {
			return true
		}
	}
	return false
}

func PathFatal(currentPath string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", currentPath, err)
	}
}

type Opts struct {
	Root          string
	Extensions    []string
	SearchPattern string
	LogFile       bool
}
