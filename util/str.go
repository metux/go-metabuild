package util

import (
	"strings"
)

func ValIf[K interface{}](b bool, yes K, no K) K {
	if b {
		return yes
	}
	return no
}

func StrListFilter(s []string, f func(s string) string) []string {
	result := make([]string, len(s))
	for x, y := range s {
		result[x] = f(y)
	}
	return result
}

func StrLinesGrep(text string, match string) []string {
	lines := []string{}
	for _, l := range strings.Split(text, "\n") {
		if strings.Contains(l, match) {
			lines = append(lines, l)
		}
	}
	return lines
}

// scan through lines and extract the x'th field
func StrLinesFieldX(lines []string, fieldnum int) []string {
	res := []string{}
	for _, l := range lines {
		a := strings.Fields(l)
		if len(a) > fieldnum {
			res = append(res, a[fieldnum])
		}
	}
	return res
}

func StrPrefix(prefix string, lines []string) []string {
	for idx, x := range lines {
		lines[idx] = prefix + x
	}
	return lines
}

func StrDirPrefix(prefix string, lines []string) []string {
	if prefix != "" {
		if !strings.HasSuffix(prefix, "/") {
			prefix = prefix + "/"
		}
		for idx, v := range lines {
			lines[idx] = prefix + v
		}
	}
	return lines
}

func StrEscLF(str string) string {
	return strings.Replace(
		strings.Trim(str, " \n"), "\n", "\\n", -1)
}

func YesNo(yes bool) string {
	if yes {
		return "yes"
	} else {
		return "no"
	}
}
