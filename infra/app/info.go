package app

import (
	"runtime"
	"strings"
)

var (
	Name           = "watermint toolbox"
	Version        = "`dev`"
	Copyright      = "© 2016-2020 Takayuki Okazaki"
	Hash           = ""
	Branch         = ""
	Zap            = ""
	BuilderKey     = ""
	BuildTimestamp = ""
	DefaultWebPort = 7800
	debugMode      = false
)

const (
	RepositoryOwner    = "watermint"
	RepositoryName     = "toolbox"
	Pkg                = "github.com/" + RepositoryOwner + "/" + RepositoryName
	ProjectStatusBadge = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
`
	ProjectLogo  = `![watermint toolbox](resources/watermint-toolbox-256x256.png)`
	StageDev     = "dev"
	StageBeta    = "beta"
	StageRelease = "release"
)

func UserAgent() string {
	return strings.ReplaceAll(Name, " ", "-") + "/" + Version
}

func ReleaseStage() string {
	switch Branch {
	case "current":
		return StageBeta
	case "master":
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

func IsDebug() bool {
	return !debugMode
}

func SetDebug(enabled bool) {
	debugMode = enabled
}
