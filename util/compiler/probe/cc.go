package probe

import (
	"log"
	"os"
	"strings"

	cmdrun "github.com/metux/go-metabuild/util/cmd"
	"github.com/metux/go-metabuild/util/compiler/base"
)

func probeCC(ci *base.CompilerInfo, envvar string, prefix string) bool {
	ci.Language = base.LangC

	if fromenv := cmdrun.EnvCmdline(envvar); len(fromenv) != 0 {
		return probeCCcmd(ci, fromenv)
	}

	compilers := []string{"gcc", "clang", "cc"}
	for _, c := range compilers {
		if probeCCcmd(ci, cmdrun.StrCmdline(prefix+c)) {
			return true
		}
	}
	return false
}

func probeCCcmd(ci *base.CompilerInfo, cmd []string) bool {
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

func DetectCC() (base.CompilerInfo, base.CompilerInfo, error) {

	crossPrefix, crossCompile := CheckCross()

	infoTarget := base.CompilerInfo{Language: base.LangC, CrossForHost: crossCompile, CrossPrefix: crossPrefix}
	infoHost := base.CompilerInfo{Language: base.LangC, CrossForBuild: crossCompile}

	if !probeCC(&infoTarget, "CC", crossPrefix) {
		return infoTarget, infoHost, base.ErrNoUsableCompiler
	}

	if crossCompile {
		if os.Getenv("HOST_CC") != "" {
			if !probeCC(&infoHost, "HOST_CC", "") {
				return base.CompilerInfo{}, base.CompilerInfo{}, base.ErrNoUsableCompiler
			}
		} else if crossPrefix != "" {
			if !probeCC(&infoHost, "", "") {
				return base.CompilerInfo{}, base.CompilerInfo{}, base.ErrNoUsableCompiler
			}
		} else {
			return infoTarget, infoHost, base.ErrCrossMissingCC
		}
	}

	// no cross, so return the target compiler twice
	return infoTarget, infoTarget, nil
}
