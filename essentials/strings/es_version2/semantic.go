package es_version2

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"github.com/watermint/toolbox/essentials/strings/es_regexp2"
	"strconv"
	"strings"
)

// Compare x, y as semantic version number. The result will be 0 if x==y, -1 if x < y, and +1 if x > y.
func Compare(x, y Version) int {
	return x.Compare(y)
}

// Max selects maximum version in versions. Returns 0.0.0 if len(versions) == 0.
func Max(versions ...Version) Version {
	switch len(versions) {
	case 0:
		return Zero()
	case 1:
		return versions[0]
	default:
		max := versions[0]
		for _, v := range versions[1:] {
			if max.Compare(v) < 0 {
				max = v
			}
		}
		return max
	}
}

// Min selects minimum version in versions. Returns 0.0.0 if len(versions) == 0.
func Min(versions ...Version) Version {
	switch len(versions) {
	case 0:
		return Zero()
	case 1:
		return versions[0]
	default:
		min := versions[0]
		for _, v := range versions[1:] {
			if min.Compare(v) > 0 {
				min = v
			}
		}
		return min
	}
}

type Version struct {
	Major      uint64 `json:"major"`
	Minor      uint64 `json:"minor"`
	Patch      uint64 `json:"patch"`
	PreRelease string `json:"pre_release"`
	Build      string `json:"build"`
}

// Zero returns "0.0.0"
func Zero() Version {
	return Version{
		Major:      0,
		Minor:      0,
		Patch:      0,
		PreRelease: "",
		Build:      "",
	}
}

func (z Version) String() string {
	v := make([]byte, 0)
	v = strconv.AppendUint(v, z.Major, 10)
	v = append(v, '.')
	v = strconv.AppendUint(v, z.Minor, 10)
	v = append(v, '.')
	v = strconv.AppendUint(v, z.Patch, 10)
	if z.PreRelease != "" {
		v = append(v, '-')
		v = append(v, z.PreRelease...)
	}
	if z.Build != "" {
		v = append(v, '+')
		v = append(v, z.Build...)
	}
	return string(v)
}

func (z Version) Equals(x Version) bool {
	return z.Compare(x) == 0
}

func (z Version) Compare(x Version) int {
	if z.Major != x.Major {
		if z.Major < x.Major {
			return -1
		}
		return 1
	}
	if z.Minor != x.Minor {
		if z.Minor < x.Minor {
			return -1
		}
		return 1
	}
	if z.Patch != x.Patch {
		if z.Patch < x.Patch {
			return -1
		}
		return 1
	}
	if z.PreRelease != "" && x.PreRelease == "" {
		return -1
	}
	if z.PreRelease == "" && x.PreRelease != "" {
		return 1
	}
	if c := strings.Compare(z.PreRelease, x.PreRelease); c != 0 {
		return c
	}
	return strings.Compare(z.Build, x.Build)
}

var (
	// https://semver.org/
	semanticRegex = es_regexp2.MustNew(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
)

const (
	subexpNameMajor         = "major"
	subexpNameMinor         = "minor"
	subexpNamePatch         = "patch"
	subexpNamePreRelease    = "prerelease"
	subexpNameBuildMetadata = "buildmetadata"
)

func MustParse(v string) Version {
	if x, out := Parse(v); out.IsError() {
		panic(out)
	} else {
		return x
	}
}

// Parse version string as semantic versioning system MAJOR.MINOR.PATCH
// Return 0.0.0 for version and an error outcome if the invalid format.
func Parse(v string) (version Version, outcome eoutcome.ParseOutcome) {
	version = Zero()
	matches, match := semanticRegex.MatchSubExp(v)
	if !match {
		return version, eoutcome.NewParseInvalidFormat("the given string does not match as semantic version format")
	}

	if v, err := strconv.ParseUint(matches[subexpNameMajor], 10, 64); err != nil {
		return version, eoutcome.NewParseInvalidFormat("the major version does not comply version format")
	} else {
		version.Major = v
	}
	if v, err := strconv.ParseUint(matches[subexpNameMinor], 10, 64); err != nil {
		return version, eoutcome.NewParseInvalidFormat("the minor version does not comply version format")
	} else {
		version.Minor = v
	}
	if v, err := strconv.ParseUint(matches[subexpNamePatch], 10, 64); err != nil {
		return version, eoutcome.NewParseInvalidFormat("the patch version does not comply version format")
	} else {
		version.Patch = v
	}
	version.PreRelease = matches[subexpNamePreRelease]
	version.Build = matches[subexpNameBuildMetadata]

	return version, eoutcome.NewParseSuccess()
}
