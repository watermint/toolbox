package app

import "runtime"

var (
	Name       = "toolbox"
	Version    = "`dev`"
	Hash       = ""
	Zap        = ""
	BuilderKey = ""
	debugMode  = false
)

const (
	Pkg                = "github.com/watermint/toolbox"
	ProjectStatusBadge = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
`
	ProjectLogo = `![watermint toolbox](resources/watermint-toolbox-256x256.png)`
)

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
