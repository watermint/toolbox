package app_build

import (
	"github.com/watermint/toolbox/essentials/strings/es_version"
	"github.com/watermint/toolbox/resources"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	// Minor version definition

	// BuildMinorOther Build with other than CI or unmanaged CI
	BuildMinorOther = 0

	// BuildMinorGitLab Build with GitLab
	BuildMinorGitLab = 1

	// BuildMinorCircleCiCurrent Build with CircleCI on other than main/master branch
	BuildMinorCircleCiCurrent = 2

	// BuildMinorGitHubActionsCurrent Build with GitHub Actions
	BuildMinorGitHubActionsCurrent = 3

	// BuildMinorCircleCiMaster Build with CircleCI on master branch
	BuildMinorCircleCiMaster = 4

	// BuildMinorCircleCiMain Build with CircleCI on main branch
	BuildMinorCircleCiMain = 4

	// BuildMinorGitHubActionsMain Build with GitHub Actions
	BuildMinorGitHubActionsMain = 8
)

// Release number
func Release() uint64 {
	if r, err := strconv.ParseUint(resources.Release(), 10, 64); err != nil {
		panic(err)
	} else {
		return r
	}
}

// SelectVersion Select or define the version
func SelectVersion(v string) es_version.Version {
	if ver, err := es_version.Parse(v); err != nil {
		return Version()
	} else {
		if ver.Major > 0 {
			return ver
		}
	}
	return Version()
}

func versionCCI() (es_version.Version, bool) {
	var err error
	if branch, ok := os.LookupEnv("CIRCLE_BRANCH"); ok {
		var patchVer uint64

		buildNum, ok := os.LookupEnv("TOOLBOX_PATCH_VERSION")
		if ok {
			patchVer, err = strconv.ParseUint(strings.TrimSpace(buildNum), 10, 64)
			if err != nil {
				panic(err)
			}
		} else {
			buildNum, ok := os.LookupEnv("CIRCLE_BUILD_NUM")
			if ok {
				patchVer, err = strconv.ParseUint(buildNum, 10, 64)
				if err != nil {
					panic(err)
				}
			}
		}

		switch branch {
		case "main":
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorCircleCiMain,
				Patch:      patchVer,
				PreRelease: "",
				Build:      "",
			}, true
		case "master":
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorCircleCiMaster,
				Patch:      patchVer,
				PreRelease: "",
				Build:      "",
			}, true
		default:
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorCircleCiCurrent,
				Patch:      patchVer,
				PreRelease: "",
				Build:      "",
			}, true
		}
	}
	return es_version.Version{}, false
}

func versionGitLab() (es_version.Version, bool) {
	if pipelineId, ok := os.LookupEnv("CI_PIPELINE_IID"); ok {
		id, err := strconv.ParseUint(pipelineId, 10, 64)
		if err != nil {
			panic(err)
		}
		return es_version.Version{
			Major:      Release(),
			Minor:      BuildMinorGitLab,
			Patch:      id,
			PreRelease: "",
			Build:      "",
		}, true
	}
	return es_version.Version{}, false
}

func versionGitHub() (es_version.Version, bool) {
	if runId, ok := os.LookupEnv("GITHUB_RUN_NUMBER"); ok {
		id, err := strconv.ParseUint(runId, 10, 64)
		if err != nil {
			panic(err)
		}

		branch, ok := os.LookupEnv("GITHUB_REF")
		if !ok {
			branch = "unknown"
		}
		switch branch {
		case "refs/heads/main":
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorGitHubActionsMain,
				Patch:      id,
				PreRelease: "",
				Build:      "",
			}, true

		case "refs/heads/current":
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorGitHubActionsCurrent,
				Patch:      id,
				PreRelease: "",
				Build:      "",
			}, true

		default:
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorOther,
				Patch:      id,
				PreRelease: "",
				Build:      "",
			}, true
		}
	}
	return es_version.Version{}, false
}

func Version() es_version.Version {
	if ver, ok := versionGitHub(); ok {
		return ver
	}
	if ver, ok := versionCCI(); ok {
		return ver
	}
	if ver, ok := versionGitLab(); ok {
		return ver
	}
	return es_version.Version{
		Major:      Release(),
		Minor:      BuildMinorOther,
		Patch:      0,
		PreRelease: "dev",
		Build:      time.Now().UTC().Format("20060102T150405Z"),
	}
}
