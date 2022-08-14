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

	// BuildMinorDocker Build with Docker on local
	BuildMinorDocker = 0

	// BuildMinorGitLab Build with GitLab
	BuildMinorGitLab = 1

	// BuildMinorCircleCiCurrent Build with CircleCI on other than main/master branch
	BuildMinorCircleCiCurrent = 2

	// BuildMinorGitHubActions Build with GitHub Actions
	BuildMinorGitHubActions = 3

	// BuildMinorCircleCiMaster Build with CircleCI on master branch
	BuildMinorCircleCiMaster = 4

	// BuildMinorCircleCiMain Build with CircleCI on main branch
	BuildMinorCircleCiMain = 8
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

func Version() es_version.Version {
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
			}
		case "master":
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorCircleCiMaster,
				Patch:      patchVer,
				PreRelease: "",
				Build:      "",
			}
		default:
			return es_version.Version{
				Major:      Release(),
				Minor:      BuildMinorCircleCiCurrent,
				Patch:      patchVer,
				PreRelease: "",
				Build:      "",
			}
		}
		panic("branch not found")
	}

	if runId, ok := os.LookupEnv("GITHUB_RUN_ID"); ok {
		id, err := strconv.ParseUint(runId, 10, 64)
		if err != nil {
			panic(err)
		}
		return es_version.Version{
			Major:      Release(),
			Minor:      BuildMinorGitHubActions,
			Patch:      id,
			PreRelease: "",
			Build:      "",
		}
	}

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
		}
	}

	return es_version.Version{
		Major:      Release(),
		Minor:      BuildMinorDocker,
		Patch:      0,
		PreRelease: "dev",
		Build:      time.Now().UTC().Format("20060102T150405Z"),
	}
}
