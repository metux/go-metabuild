package base

import (
	"strings"
)

type logWriter struct {
	Prefix  string
	Builder *BaseBuilder
}

func (log logWriter) Write(p []byte) (n int, err error) {
	for _, s1 := range strings.Split(strings.Trim(string(p), " \n"), "\n") {
		log.Builder.Logf("(%s) %s", log.Prefix, s1)
	}
	return len(p), nil
}
