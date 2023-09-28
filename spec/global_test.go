package spec

import (
	"testing"
)

func TestLoad(t *testing.T) {
	LoadTestSpec(t)

	t.Log("-- Checks")
	for _, y := range GlobalTestSpec.GetChecks() {
		y.Init()
		t.Logf("check %s\n", y)
	}

	t.Log("")

	t.Log("-- Objects")
	for x, y := range GlobalTestSpec.GetTargetObjects() {
		t.Logf("object %s: %s -- %s\n", x, y, y.Type())
	}

	t.Log("")
	t.Logf("done")
}
