package fileutil

import (
	"path/filepath"
)

// Note: this *must* produce a relative path, *must not* do Abs()
func MkPath(dir string, fn string) string {
	if dir == "" {
		return fn
	}
	return filepath.Clean(dir + "/" + fn)
}
