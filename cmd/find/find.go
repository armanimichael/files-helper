package find

import (
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/pkg/util"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

func SearchInFiles(opts cmd.Opts) {
	filepath.WalkDir(opts.Root, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			cmd.LogPathFatal(currentPath, err)
		}
		// Skip if is dir or no input extension
		if d.IsDir() || !cmd.FilterExtension(path.Ext(currentPath), opts.Extensions) {
			return nil
		}

		file, err := os.OpenFile(currentPath, os.O_RDONLY, os.ModeType)
		if err != nil {
			cmd.LogPathFatal(currentPath, err)
		}
		defer file.Close()

		found, err := util.IsInReader(file, opts.SearchPattern)
		if err != nil {
			cmd.LogPathFatal(currentPath, err)
		}
		if found && opts.LogFile {
			log.Println(currentPath)
		}

		return nil
	})
}
