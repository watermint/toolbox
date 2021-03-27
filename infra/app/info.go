package app

import (
	"runtime"
	"strings"
)

var (
	Name           = "watermint toolbox"
	Version        = "`dev`"
	Copyright      = "© 2016-2021 Takayuki Okazaki"
	Hash           = ""
	Branch         = ""
	Zap            = ""
	BuilderKey     = ""
	BuildTimestamp = ""
	DefaultWebPort = 7800
)

func UserAgent() string {
	return strings.ReplaceAll(Name, " ", "-") + "/" + Version
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
