package gnu

import (
	"strings"

	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/cmd"
)

const (
	ErrArchUnknown = util.Error("objdump - unknown architecture")
)

func RunObjdump(c []string, args []string) (string, error) {
	if len(c) == 0 {
		c = []string{"objdump"}
	}
	c = append(c, args...)
	return cmd.RunOut(c, true)
}

func ArchXlate(a string) (string, error) {
	if strings.HasSuffix(a, "i386") {
		return "i386", nil
	}
	return a, ErrArchUnknown
}

func ELFDepends(cmd []string, fn string) ([]string, error) {
	out, err := RunObjdump(cmd, []string{"-x", fn})
	return util.StrLinesFieldX(util.StrLinesGrep(out, "NEEDED "), 1), err
}

func ELFArch(cmd []string, fn string) (string, error) {
	out, err := RunObjdump(cmd, []string{"-a", fn})
	if err != nil {
		return "", err
	}
	return util.StrLinesFieldX(util.StrLinesGrep(out, "file format "), 3)[0], nil
}
