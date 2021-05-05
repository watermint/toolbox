package app

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_build"
	"github.com/watermint/toolbox/resources"
	"runtime"
	"strings"
)

var (
	Name           = "watermint toolbox"
	Version        = app_build.SelectVersion(BuildInfo.Version)
	BuildInfo      = resources.Build()
	BuildId        = Version.String()
	Release        = resources.Release()
	Copyright      = fmt.Sprintf("© 2016-%4d Takayuki Okazaki", BuildInfo.Year)
	LandingPage    = "https://toolbox.watermint.org"
	DefaultWebPort = 7800
)

func UserAgent() string {
	return strings.ReplaceAll(Name, " ", "-") + "/" + BuildId
}

func ReleaseStage() string {
	switch BuildInfo.Branch {
	case "current":
		return StageBeta
	case "master", "main":
		return StageRelease
	default:
		return StageDev
	}
}

func IsProduction() bool {
	return BuildInfo.Production && ReleaseStage() == StageRelease
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
