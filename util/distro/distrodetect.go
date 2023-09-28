package distro

import (
	"github.com/go-ini/ini"
)

// FIXME: find something smaller than go-ini
func DistroDetect(sysroot string) (DistroInfo, error) {
	// try os-release
	cfg, err := ini.Load(sysroot + "/etc/os-release")
	if err != nil {
		return DistroInfo{}, err
	}

	sect := cfg.Section("")
	inf := DistroInfo{
		Id:              sect.Key("ID").String(),
		IdLike:          sect.Key("ID_LIKE").String(),
		Name:            sect.Key("NAME").String(),
		Version:         sect.Key("VERSION").String(),
		VersionId:       sect.Key("VERSION_ID").String(),
		VersionCodename: sect.Key("VERSION_CODENAME").String(),
		PrettyName:      sect.Key("PRETTY_NAME").String(),
		HomeUrl:         sect.Key("HOME_URL").String(),
		SupportUrl:      sect.Key("SUPPORT_URL").String(),
		BugReportUrl:    sect.Key("BUG_REPORT_URL").String(),
	}

	if inf.IdLike == "" {
		inf.IdLike = inf.Id
	}

	return inf, nil
}
