package engine

import (
	"os"
	"testing"

	magic "github.com/metux/go-magicdict"
)

const (
	StdStage = StagePackage
)

var (
	TestingEngine *Engine
)

func LoadTestingEngine(t *testing.T) {
	// jump.to projects root
	os.Chdir("../")

	if TestingEngine != nil {
		return
	}

	engine := Engine{}
	if err := engine.Load("examples/pkg/zlib.yaml", "examples/settings.yaml"); err != nil {
		t.Fatalf("yaml load failed: %s", err)
	}

	TestingEngine = &engine
}

// extra check whether SpecObj really implements the magic.Entry interface
func test1(t *testing.T, ent magic.Entry) {
	t.Logf("%s", ent.Keys())
}

func TestEngine(t *testing.T) {
	LoadTestingEngine(t)
	if err := TestingEngine.RunStage(StdStage); err != nil {
		t.Fatal("build error:", err)
	}

	test1(t, (*TestingEngine).Global)
	(*TestingEngine).Global.YamlStore("/tmp/metabuild-dump.yaml", 0644)
}
