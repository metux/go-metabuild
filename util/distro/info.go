package distro

type DistroInfo struct {
	Id              string
	IdLike          string
	Name            string
	VersionId       string
	Version         string
	VersionCodename string
	PrettyName      string
	HomeUrl         string
	SupportUrl      string
	BugReportUrl    string
	PkgFormat       string
}
