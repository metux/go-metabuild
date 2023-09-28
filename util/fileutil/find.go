package fileutil

import (
	"github.com/metux/go-metabuild/util/cmd"
)

func FindFile(dir string, name string) []string {
	files, _ := cmd.RunOutLines([]string{"find", dir, "-name", name}, true)
	return files
}
