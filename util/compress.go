package util

import (
	"compress/gzip"
	"os"
)

func GzipCompress(infile string, outfile string, fmode os.FileMode) error {
	data, err := os.ReadFile(infile)
	if err != nil {
		return err
	}

	handle, err := os.OpenFile(outfile, os.O_CREATE|os.O_RDWR, fmode)
	if err != nil {
		return err
	}
	defer handle.Close()

	zipWriter, err := gzip.NewWriterLevel(handle, 9)
	if err != nil {
		return err
	}

	if _, err := zipWriter.Write(data); err != nil {
		return err
	}
	if err = zipWriter.Close(); err != nil {
		return err
	}
	return nil
}
