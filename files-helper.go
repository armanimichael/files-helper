package main

import (
	"flag"
	"fmt"
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/cmd/find"
	"github.com/armanimichael/files-helper/cmd/replace"
	"strings"
)

func noBKalert() bool {
	var resp string
	fmt.Print("This operation could alter files, do you want to create a backup? [Y/n]: ")
	fmt.Scanf("%s", &resp)
	return resp != "n"
}

func main() {
	command := flag.String("cmd", "", "Util to run")
	rootDir := flag.String("root", "./", "Root path")
	searchPattern := flag.String("pattern", "", "Search pattern")
	replaceStr := flag.String("replace", "", "What to repalace the search pattern with")
	extensionsStr := flag.String("extensions", "", "Lookable file extensions separated by comma (ex. txt,html,go)")
	backup := flag.Bool("backup", false, "Backup matching file before")
	flag.Parse()

	extensions := strings.Split(*extensionsStr, ",")
	opts := cmd.Opts{
		Root:          *rootDir,
		Extensions:    extensions,
		SearchPattern: *searchPattern,
		Replace:       *replaceStr,
		LogFile:       true,
		Backup:        *backup,
	}

	switch *command {
	case "find":
		find.SearchInFiles(opts)
	case "replace":
		*backup = noBKalert()
		replace.SubstituteInFiles(opts)
	}
}
