package app_definitions

import "github.com/watermint/toolbox/resources"

// Project information
const (
	// CoreRepositoryOwner is the owner of the core repository.
	CoreRepositoryOwner = "watermint"

	// CoreRepositoryName is the name of the core repository.
	CoreRepositoryName = "toolbox"

	// CorePkg is the package name of the core repository.
	CorePkg = "github.com/" + CoreRepositoryOwner + "/" + CoreRepositoryName

	// SupplementRepositoryOwner is the owner of the supplement repository.
	SupplementRepositoryOwner = "watermint"

	// SupplementRepositoryName is the name of the supplement repository.
	SupplementRepositoryName = "toolbox-supplement"

	// SupplementRepositoryLicenseUrl is the URL of the license files
	SupplementRepositoryLicenseUrl = "https://raw.githubusercontent.com/watermint/toolbox-supplement/main/licenses/"
)

var (
	ApplicationRepositoryOwner = CoreRepositoryOwner
	ApplicationRepositoryName  = CoreRepositoryName
)

// Project structure definitions
var (
	RecipePackageNames = []string{
		"ingredient",
		"recipe",
		"citron",
	}

	RecipeFlavors = []string{
		RecipeFlavorCitron,
		RecipeFlavorLime,
	}
)

const (
	RecipeFlavorCitron = "citron"
	RecipeFlavorLime   = "lime"
)

// Peer names
const (
	// Peer name for deployment
	PeerDeploy = "deploy"
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

	// EnvNameToolboxLicenseSalt Env variable name for license salt
	EnvNameToolboxLicenseSalt = resources.EnvNameToolboxLicenseSalt

	// EnvNameToolboxLicenseKey Env variable name for license key
	EnvNameToolboxLicenseKey = "TOOLBOX_LICENSE_KEY"
)

// Document
var (
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
	LogStats   = "stats"

	// Log names
	LogNameStart  = "recipe.log"
	LogNameFinish = "result.log"
)
