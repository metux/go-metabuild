package fileutil

import (
	"os"
	"strings"

	"github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/strs"
)

func FindFile(dir string, name string) []string {
	files, _ := cmd.RunOut([]string{"find", dir, "-name", name}, true)
	return strs.SplitNL(files)
}

func ListDir(dir string, suffix string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	names := []string{}
	for _, e := range entries {
		n := e.Name()
		if strings.HasSuffix(n, suffix) {
			names = append(names, dir+"/"+n)
		}
	}
	return names, err
}
