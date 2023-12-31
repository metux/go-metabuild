package check

import (
	"github.com/metux/go-magicdict/api"
)

type Key = api.Key

const (
	// types
	KeyCHeader     = Key("c/header")
	KeyCFunction   = Key("c/function")
	KeyCType       = Key("c/type")
	KeyCCompiler   = Key("c/compiler")
	KeyCXXCompiler = Key("c++/compiler")
	KeyPkgConfig   = Key("pkgconf")
	KeyPkgConfAdd  = Key("pkgconf/add")

	KeyCDefines = Key("c/defines")
	KeyCLDFlags = Key("c/ldflags")
	KeyCCFlags  = Key("c/cflags")

	KeyInstallDirs  = Key("install-dirs")
	KeyTargetDistro = Key("target-distro")

	KeyMandatory = Key("mandatory")
	KeyBuild     = Key("build")
	KeyDefines   = Key("defines")

	KeyHashAttrs = Key("@hash")
	KeyDone      = Key("@done")
	KeyPresent   = Key("@present")

	KeyId   = Key("id")
	KeyType = Key("type")

	KeyGitDescribe = Key("git/describe")

	KeyI18nLinguas = Key("i18n/linguas")
	KeyI18nPoDir   = Key("i18n/po/dir")
)
