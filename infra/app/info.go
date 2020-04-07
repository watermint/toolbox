package app

import "runtime"

var (
	Name           = "watermint toolbox"
	Version        = "`dev`"
	Copyright      = "© 2016-2020 Takayuki Okazaki"
	Hash           = ""
	Zap            = ""
	BuilderKey     = ""
	BuildTimestamp = ""
	DefaultWebPort = 7800
	debugMode      = false
)

const (
	Pkg                = "github.com/watermint/toolbox"
	ProjectStatusBadge = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
`
	ProjectLogo = `![watermint toolbox](resources/watermint-toolbox-256x256.png)`
)

func UserAgent() string {
	return Name + "/" + Version
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
