package git

import (
	"sort"
	"strings"

	"github.com/metux/go-metabuild/util"
	"github.com/metux/go-metabuild/util/cmd"
)

const delim = " ||| "

func Authors() ([]string, error) {
	c := append(gitCmd, "log", `--pretty=format:%an`+delim+`%ae`)

	if out, err := cmd.RunOutLines(c, true); err == nil {
		m := make(map[string]string)
		for _, l := range out {
			if name, email := util.StrSplitTwo(l, delim); email == "" {
				if _, ok := m[name]; !ok {
					m[name] = ""
				}
			} else {
				m[name] = strings.ToLower(email)
			}
		}
		lines := []string{}
		for k, v := range m {
			if v == "" {
				lines = append(lines, k)
			} else {
				lines = append(lines, k+" <"+v+">")
			}
		}
		sort.Strings(lines)
		return lines, nil
	} else {
		return []string{}, err
	}
}
