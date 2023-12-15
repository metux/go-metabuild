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

func StripAbs(fn string) string {
	for fn != "" && fn[0] == '/' {
		fn = fn[1:]
	}
	return fn
}
