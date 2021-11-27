package find

import (
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/pkg/util"
	"io/fs"
	"log"
	"path/filepath"
)

func SearchInFiles(opts cmd.Opts) {
	filepath.WalkDir(opts.Root, func(currentPath string, d fs.DirEntry, err error) error {
		cmd.PathFatal(currentPath, err)
		if !cmd.IsSupportedPath(d, currentPath, opts.Extensions) {
			return nil
		}

		file := cmd.ReadFile(currentPath, err)
		defer file.Close()

		// Handle found content
		found, err := util.IsInReader(file, opts.SearchPattern)
		cmd.PathFatal(currentPath, err)
		if found && opts.LogFile {
			log.Println(currentPath)
		}
		return nil
	})

}
