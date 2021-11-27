package find

import (
	"github.com/armanimichael/files-helper/cmd"
	"log"
	"os"
)

func SearchInFiles(opts cmd.Opts) {
	cmd.HandleFoundFiles(opts, func(file *os.File, currentPath string, opts cmd.Opts) {
		if opts.LogFile {
			log.Println(currentPath)
		}
	})
}
