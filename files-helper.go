package main

import (
	"flag"
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/cmd/find"
	"strings"
)

func main() {
	command := flag.String("cmd", "", "Util to run")
	rootDir := flag.String("root", "./", "Root path")
	searchPattern := flag.String("pattern", "", "Search pattern")
	extensionsStr := flag.String("extensions", "", "Lookable file extensions separated by comma (ex. txt,html,go)")
	flag.Parse()

	extensions := strings.Split(*extensionsStr, ",")
	opts := cmd.Opts{
		Root:          *rootDir,
		Extensions:    extensions,
		SearchPattern: *searchPattern,
		LogFile:       true,
	}

	switch *command {
	case "find":
		find.SearchInFiles(opts)
	}
}
