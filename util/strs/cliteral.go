package strs

import (
	"strings"
)

func CLiteral(str string, indent string, maxWidth int) string {
	escaped := StrEscC(str)

	sb := strings.Builder{}

	l1 := ""
	for len(escaped) > maxWidth {
		l1 = escaped[:maxWidth]
		escaped = escaped[maxWidth:]
		if l1[len(l1)-1] == '\\' {
			l1 = l1 + escaped[:1]
			escaped = escaped[1:]
		}
		sb.WriteString(indent)
		sb.WriteString("\"")
		sb.WriteString(l1)
		sb.WriteString("\"\n")
	}

	if len(escaped) > 0 {
		sb.WriteString(indent)
		sb.WriteString("\"")
		sb.WriteString(escaped)
		sb.WriteString("\"\n")
	}

	return sb.String()
}
