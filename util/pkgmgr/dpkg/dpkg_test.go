package dpkg

import (
	"testing"
)

func TestDpkgSearchFile(t *testing.T) {
	dpkg := NewDpkg([]string{}, []string{}, "")

	m := dpkg.SearchInstalledFile("libc.so.6")

	if len(m) == 0 {
		t.Fatalf("cant find libc in dpkg database ... something seems really wrong")
	}

	for _, v := range m {
		t.Logf("pkg::: %+v", v)
	}
}
