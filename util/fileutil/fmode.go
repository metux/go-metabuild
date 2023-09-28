package fileutil

import (
	"os"
	"strconv"
)

// FIXME: should also parse textual representation
func FileModeParse(s string) (os.FileMode, error) {
	n, err := strconv.ParseUint(s, 8, 32)
	return os.FileMode(n), err
}
