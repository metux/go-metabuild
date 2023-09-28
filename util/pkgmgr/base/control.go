package base

type PkgControl struct {
	Package      string
	Version      string
	Maintainer   string
	Architecture string
	Section      string
	Description  string
	Depend       []string
	MultiArch    string
	Origin       string
	Bugs         string
	Homepage     string
	Priority     string
}
