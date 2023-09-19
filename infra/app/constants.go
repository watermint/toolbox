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

// Auth database
const (
	AuthDatabaseDefaultName = "secrets.db"
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

	// Env variable name of quiet mode on test.
	EnvNameTestQuiet = "TOOLBOX_TEST_QUIET"

	// Env variable name of deploy token. Expected format is JSON encoded tokens.
	EnvNameDeployToken = "TOOLBOX_DEPLOY_TOKEN"

	// Env variable name for skip end to end tests. Expected format is bool.
	EnvNameEndToEndSkipTest = "TOOLBOX_SKIPENDTOENDTEST"

	// Env variable name for end to end token. Expected format is JSON encoded tokens
	EnvNameEndToEndToken = "TOOLBOX_ENDTOEND_TOKEN"

	// Env variable name for replay store path.
	EnvNameReplayPath = "TOOLBOX_REPLAY_PATH"

	// Env variable name for replay store shared link.
	EnvNameReplayUrl = "TOOLBOX_REPLAY_URL"

	// Env variable name for toolbox home
	EnvNameToolboxHome = "TOOLBOX_HOME"

	// Env variable name for build signature
	EnvNameToolboxBuilderKey = "TOOLBOX_BUILDERKEY"

	// Env variable name for build app keys data
	EnvNameToolboxAppKeys = "TOOLBOX_APPKEYS"

	// EnvNameToolboxBuildTarget Env variable name for build target
	EnvNameToolboxBuildTarget = "TOOLBOX_BUILD_TARGET"
)

// Document
const (
	// Project status badge
	ProjectStatusBadge = `
[![Build](https://github.com/watermint/toolbox/actions/workflows/build.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/build.yml)
[![Test](https://github.com/watermint/toolbox/actions/workflows/test.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/test.yml)
[![CodeQL](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/watermint/toolbox/actions/workflows/codeql-analysis.yml)
[![Codecov](https://codecov.io/gh/watermint/toolbox/branch/main/graph/badge.svg?token=CrE8reSVvE)](https://codecov.io/gh/watermint/toolbox)
`

	// Project logo
	ProjectLogo = `![watermint toolbox](resources/images/watermint-toolbox-256x256.png)`
)

const (
	// Log prefix
	LogToolbox = "toolbox"
	LogCapture = "capture"
	LogSummary = "summary"

	// Log names
	LogNameStart  = "recipe.log"
	LogNameFinish = "result.log"
)
