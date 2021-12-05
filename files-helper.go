package main

import (
	"flag"
	"fmt"
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/cmd/find"
	"github.com/armanimichael/files-helper/cmd/replace"
	"os"
	"strings"
)

func noBKalert() bool {
	var resp string
	fmt.Print("This operation could alter files, do you want to create a backup? [Y/n]: ")
	fmt.Scanf("%s", &resp)
	return resp != "n"
}

func ensureMandatoryFields(command, pattern, replace string) {
	if command == "" {
		fmt.Println("You must specify a command to run (ex. `-cmd find`)")
		os.Exit(1)
	}

	if pattern == "" {
		fmt.Println("You must specify a search pattern (ex. `-pattern test`)")
		os.Exit(1)
	}

	if command == "replace" && replace == "" {
		fmt.Println("You must specify a replace string (ex. `-replace test`)")
		os.Exit(1)
	}
}

func ensureExistingField(command string) {
	if command != "find" && command != "replace" {
		fmt.Printf("The command `%s` does not exist (try with `find` or `replace`)\n", command)
		os.Exit(1)
	}
}

func main() {
	command := flag.String("cmd", "", "Util to run")
	rootDir := flag.String("root", "./", "Root path")
	searchPattern := flag.String("pattern", "", "Search pattern")
	replaceStr := flag.String("replace", "", "What to replace the search pattern with")
	extensionsStr := flag.String("extensions", "txt", "Lookable file extensions separated by comma (ex. txt,html,go)")
	backup := flag.Bool("backup", false, "Backup matching file before")
	flag.Parse()

	ensureMandatoryFields(*command, *searchPattern, *replaceStr)
	ensureExistingField(*command)
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
		if !(*backup) {
			*backup = noBKalert()
		}
		replace.SubstituteInFiles(opts)
	}
}
