package main

import (
	"fmt"
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/cmd/find"
	"github.com/armanimichael/files-helper/cmd/replace"
	"os"
	"strings"
)

func main() {
	if !areMandatoryArgsPresent() {
		fmt.Println("You must specify a command to run [find | replace]")
		os.Exit(1)
	}

	command := os.Args[1]

	opts := cmd.Opts{
		Root:          getTargetDir(),
		Extensions:    getFilteredExtensions(),
		SearchPattern: getSearchPattern(),
		LogFile:       getLoggingChoice(),
		Replace:       getReplaceString(command),
		Backup:        getBackupChoice(command),
	}

	switch command {
	case "find":
		find.SearchInFiles(opts)
	case "replace":
		replace.SubstituteInFiles(opts)
	}
}

func getTargetDir() string {
	targetDir := "./"
	fmt.Printf("Select target directory [./]: ")
	fmt.Scanln(&targetDir)

	return targetDir
}

func getSearchPattern() string {
	searchPattern := ""
	fmt.Printf("Search pattern: ")
	fmt.Scanln(&searchPattern)

	if searchPattern == "" {
		fmt.Printf("You must specify a search pattern!")
		os.Exit(1)
	}

	return searchPattern
}

func getFilteredExtensions() []string {
	extensionsStr := "txt"
	fmt.Printf("Filter by extension (comma separated) [txt]: ")
	fmt.Scanln(&extensionsStr)

	return strings.Split(extensionsStr, ",")
}

func getReplaceString(command string) string {
	replaceStr := ""

	if command == "replace" {
		fmt.Printf("Replace with: ")
		fmt.Scanln(&replaceStr)
	}

	return replaceStr
}

func getBackupChoice(command string) bool {
	backup := true

	if command == "replace" {
		fmt.Printf("Perform backup? [true]: ")
		fmt.Scanf("%t\n", &backup)
	}

	return backup
}

func getLoggingChoice() bool {
	allowLogging := true
	fmt.Printf("Verbose? [true]: ")
	fmt.Scanf("%t\n", &allowLogging)

	return allowLogging
}

func areMandatoryArgsPresent() bool {
	allowedCommands := map[string]bool{
		"replace": true,
		"find":    true,
	}

	_, ok := allowedCommands[os.Args[1]]
	return len(os.Args) >= 2 && ok
}
