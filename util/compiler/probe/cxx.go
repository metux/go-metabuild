package probe

import (
	"log"
	"os"
	"strings"

	cmdrun "github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler/base"
)

func probeCXX(ci *base.CompilerInfo, envvar string, prefix string) bool {
	ci.Language = base.LangCxx

	if fromenv := cmdrun.EnvCmdline(envvar); len(fromenv) != 0 {
		return probeCXXcmd(ci, fromenv)
	}

	// FIXME: clang detection yet broken
	compilers := []string{"g++", "clang++", "c++"}
	for _, c := range compilers {
		if probeCXXcmd(ci, cmdrun.StrCmdline(prefix+c)) {
			return true
		}
	}
	return false
}

func probeCXXcmd(ci *base.CompilerInfo, cmd []string) bool {
	if len(cmd) == 0 || cmd[0] == "" {
		return false
	}
	ci.Command = cmd

	// try to guess from help output

	out, err := cmdrun.RunOutOne(append(ci.Command, "--help"), false)
	if err != nil {
		log.Println("probeCompilerId: cant get help page")
		return false
	}
	if strings.Contains(out, "clang LLVM compiler") {
		return probeClang(ci)
	}
	if probeGCC(ci) {
		return true
	}

	// cant detect it
	log.Println("neither gcc nor clang")
	return false
}

func DetectCXX() (base.CompilerInfo, base.CompilerInfo, error) {

	crossPrefix, crossCompile := CheckCross()

	infoTarget := base.CompilerInfo{Language: base.LangCxx, CrossForHost: crossCompile, CrossPrefix: crossPrefix}
	infoHost := base.CompilerInfo{Language: base.LangCxx, CrossForBuild: crossCompile}

	if !probeCXX(&infoTarget, "CXX", crossPrefix) {
		return infoTarget, infoHost, base.ErrNoUsableCompiler
	}

	if crossCompile {
		if os.Getenv("HOST_CXX") != "" {
			if !probeCXX(&infoHost, "HOST_CXX", "") {
				return base.CompilerInfo{}, base.CompilerInfo{}, base.ErrNoUsableCompiler
			}
		} else if crossPrefix != "" {
			if !probeCXX(&infoHost, "", "") {
				return base.CompilerInfo{}, base.CompilerInfo{}, base.ErrNoUsableCompiler
			}
		} else {
			return infoTarget, infoHost, base.ErrCrossMissingCXX
		}
	}

	// no cross, so return the target compiler twice
	return infoTarget, infoTarget, nil
}
