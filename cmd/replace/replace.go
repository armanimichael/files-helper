package replace

import (
	"github.com/armanimichael/files-helper/cmd"
	"github.com/armanimichael/files-helper/pkg/util"
	"io/ioutil"
	"log"
	"os"
)

func SubstituteInFiles(opts cmd.Opts) {
	cmd.HandleFoundFiles(opts, func(file *os.File, currentPath string, opts cmd.Opts) {
		if opts.LogFile {
			log.Println(currentPath)
		}

		newContent, err := util.Replace(file, opts.SearchPattern, opts.Replace)
		if err != nil {
			cmd.PathFatal(currentPath, err)
		}
		if ioutil.WriteFile(currentPath, newContent, 02) != nil {
			cmd.PathFatal(currentPath, err)
		}
		if err := file.Close(); err != nil {
			cmd.PathFatal(currentPath, err)
		}
	})
}
