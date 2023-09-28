package util

import (
	"os"
	"strings"
)

type CSource struct {
	Includes []string
	MainBody strings.Builder
	Text     strings.Builder
}

func (cs *CSource) AddInclude(i string) {
	cs.Includes = append(cs.Includes, i)
}

func (cs *CSource) AddIncludes(i []string) {
	for _, x := range i {
		cs.AddInclude(x)
	}
}

func (cs *CSource) Write(fn string) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	// write includes
	for _, hdr := range cs.Includes {
		f.WriteString("#include <")
		f.WriteString(hdr)
		f.WriteString(">\n")
	}

	f.WriteString(cs.Text.String())

	// write main function
	f.WriteString("int main() {\n")
	f.WriteString(cs.MainBody.String())
	f.WriteString("\n}\n")

	return nil
}
