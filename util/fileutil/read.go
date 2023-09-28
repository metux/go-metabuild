package fileutil

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/metux/go-metabuild/util"
)

func ReadLines(fn string) []string {
	bytesRead, _ := ioutil.ReadFile(fn)
	fileContent := string(bytesRead)
	return strings.Split(fileContent, "\n")
}

func ReadLinesGlob(glob string) ([]string, error) {
	lines := []string{}

	files, err := filepath.Glob(glob)
	if err == nil {
		for _, ent := range files {
			lines = append(lines, ReadLines(ent)...)
		}
	}

	return lines, err
}

func ReadLinesGlobUniq(glob string) ([]string, error) {
	lines, err := ReadLinesGlob(glob)
	return util.Uniq(lines), err
}
