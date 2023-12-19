package fileutil

import (
	"os"
	"path/filepath"
)

func ScanFiles(dirname string) ([]string, error) {
	l := len(dirname) + 1
	res := []string{}

	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if path != dirname && !info.IsDir() {
				res = append(res, path[l:])
			}
			return nil
		})

	return res, err
}
