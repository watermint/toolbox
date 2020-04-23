package app

// Project information
const (
	// project owner on GitHub
	RepositoryOwner = "watermint"

	// Project name on GitHub
	RepositoryName = "toolbox"

	// Package name
	Pkg = "github.com/" + RepositoryOwner + "/" + RepositoryName
)

// Peer names
const (
	// Peer name for deployment
	PeerDeploy = "deploy"

	// Peer name for end to end test.
	PeerEndToEndTest = "end_to_end_test"
)

// Stages of deployment
const (
	StageDev     = "dev"
	StageBeta    = "beta"
	StageRelease = "release"
)

// Environment variable names
const (
	// Env variable name of verbose debug mode. Expected format is bool.
	EnvNameDebugVerbose = "TOOLBOX_DEBUG_VERBOSE"

	// Env variable name of deploy token. Expected format is JSON encoded tokens.
	EnvNameDeployToken = "TOOLBOX_DEPLOY_TOKEN"

	// Env variable name for skip end to end tests. Expected format is bool.
	EnvNameEndToEndSkipTest = "TOOLBOX_SKIPENDTOENDTEST"

	// Env variable name for end to end token. Expected format is JSON encoded tokens
	EnvNameEndToEndToken = "TOOLBOX_ENDTOEND_TOKEN"

	// Env variable name for test resource file. Expected format is file path.
	EnvNameTestResource = "TOOLBOX_TEST_RESOURCE"
)

// Test resources
const (
	// Keys for ControlTestExtension keys
	CtlTestExtUseMock = "use_mock"
)

// Document
const (
	// Project status badge
	ProjectStatusBadge = `
[![CircleCI](https://circleci.com/gh/watermint/toolbox.svg?style=shield)](https://circleci.com/gh/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg)](https://coveralls.io/github/watermint/toolbox)
`

	// Project logo
	ProjectLogo = `![watermint toolbox](resources/watermint-toolbox-256x256.png)`
)
