package app

import (
	"github.com/watermint/toolbox/infra/control/app_build"
	"github.com/watermint/toolbox/resources"
	"runtime"
	"strings"
)

var (
	Name           = "watermint toolbox"
	Build          = app_build.Version()
	BuildId        = Build.String()
	Release        = resources.Release()
	Copyright      = "© 2016-2021 Takayuki Okazaki"
	Hash           = ""
	Branch         = ""
	Zap            = ""
	BuilderKey     = ""
	BuildTimestamp = ""
	DefaultWebPort = 7800
)

func UserAgent() string {
	return strings.ReplaceAll(Name, " ", "-") + "/" + BuildId
}

func ReleaseStage() string {
	switch Branch {
	case "current":
		return StageBeta
	case "master", "main":
		return StageRelease
	default:
		return StageDev
	}
}

func IsProduction() bool {
	return Hash != ""
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
