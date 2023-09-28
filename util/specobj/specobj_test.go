package specobj

import (
	"testing"

	magic "github.com/metux/go-magicdict"
)

// test whether MagicDict supports Entry interface
func t1(t *testing.T, ent magic.Entry) {
	t.Logf("%s", ent.Keys())
}

func TestSpecObj(t *testing.T) {
	so, err := magic.YamlLoad("../../examples/pkg/zlib.yaml", "../../examples/settings.yaml")
	if err != nil {
		t.Fatalf("failed loading yaml: %s", err)
	}
	t1(t, so)
}
