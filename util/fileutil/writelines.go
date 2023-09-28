package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func WriteFileLines(fn string, lines []string) error {
	os.MkdirAll(filepath.Dir(fn), 0755)

	f, err := os.Create(fn)
	defer f.Close()
	if err != nil {
		return err
	}

	for _, l := range lines {
		f.WriteString(l)
		f.WriteString("\n")
	}

	return nil
}

func WriteText(fn string, t string) error {
	f, err := os.Create(fn)
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(t)
	return nil
}

func WriteTemplate(input string, output string, marker string, text string) error {
	// read template file
	tmpl, err := os.ReadFile(input) // just pass the file name
	if err != nil {
		return err
	}

	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(strings.Replace(string(tmpl), marker, text, 1))
	return err
}

type ExtWriter struct {
	os.File
}

func (e ExtWriter) WriteFolded(name string, val string) {
	// prepare the description field
	if val := strings.TrimSpace(val); val != "" {
		e.WriteString(name)
		e.WriteString(":")
		for _, l := range strings.Split(val, "\n") {
			e.WriteString(" ")
			if l == "" {
				e.WriteString(".")
			} else {
				e.WriteString(l)
			}
			e.WriteString("\n")
		}
	}
}

func (e ExtWriter) WriteField(name string, val string) {
	if val != "" {
		e.WriteString(name)
		e.WriteString(": ")
		e.WriteString(val)
		e.WriteString("\n")
	}
}

func (e ExtWriter) WriteVar(name string, val string) {
	if val != "" {
		e.WriteString(name)
		e.WriteString("=")
		e.WriteString(val)
		e.WriteString("\n")
	}
}

func CreateWriter(fn string) (ExtWriter, error) {
	if fn == "" {
		panic("CreateWriter: empty file name")
	}
	f, err := os.Create(fn)
	if err != nil {
		return ExtWriter{}, err
	}
	return ExtWriter{*f}, err
}
