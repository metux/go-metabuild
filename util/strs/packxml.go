package strs

import (
	"regexp"
	"strings"
)

var (
	reXMLcomments = regexp.MustCompile(`\<\!\-\-.*\-\-\>`)
	reXMLtrim     = regexp.MustCompile(`\>[ \n\r]+\<`)
)

func PackXML(data string) string {
	return strings.Trim(
		reXMLtrim.ReplaceAllString(
			reXMLcomments.ReplaceAllString(string(data), ""), "><"),
		" \n\r")
}
