package strs

import (
	"fmt"
	"strings"
)

func StrEscC(in string) string {
	sbEsc := strings.Builder{}
	for idx := 0; idx < len(in); idx++ {
		switch rune := in[idx]; rune {
		case '"':
			sbEsc.WriteString(`\"`)
		case '\'':
			sbEsc.WriteString(`\'`)
		case '\n':
			sbEsc.WriteString(`\n`)
		case '\\':
			sbEsc.WriteString(`\\`)
		default:
			if b := byte(rune); b > 127 {
				sbEsc.WriteString(fmt.Sprintf(`\%o`, b))
			} else {
				sbEsc.WriteString(string(rune))
			}
		}
	}
	return sbEsc.String()
}
