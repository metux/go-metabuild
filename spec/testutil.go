package spec

// this file is just for tests in other packages

import (
	"os"
	"testing"

	"github.com/metux/go-metabuild/spec/global"
)

var (
	GlobalTestSpec *global.Global
)

func LoadTestSpec(t *testing.T) {
	os.Chdir("../examples")

	if GlobalTestSpec != nil {
		return
	}
	if gs, err := global.LoadGlobal("pkg/zlib/metabuild.yaml", "conf/settings.xml"); err == nil {
		GlobalTestSpec = &gs
	} else {
		t.Fatalf("yaml load failed: %s", err)
	}
}
