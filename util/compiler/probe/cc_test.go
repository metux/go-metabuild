package probe

import (
	"testing"
)

func TestCompilerID(t *testing.T) {
	triplet, cross := CheckCross()
	t.Logf("cross: %s %t", triplet, cross)

	target, host, err := DetectCC()
	t.Logf("target: %s\n", target)
	t.Logf("host:   %s\n", host)
	if err != nil {
		t.Errorf("error: %s\n", err)
	}
}
