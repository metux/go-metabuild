package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

func FileMD5(fn string) (string, error) {
	h := md5.New()

	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MD5Files(dir string, files []string) (string, error) {
	sb := strings.Builder{}
	for _, fn := range files {
		md5, err := FileMD5(dir + "/" + fn)
		if err != nil {
			return "", err
		}
		sb.WriteString(md5)
		sb.WriteString(" ")
		sb.WriteString(fn)
		sb.WriteString("\n")
	}
	return sb.String(), nil
}
