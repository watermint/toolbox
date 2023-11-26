package eregexp

import "testing"

func TestReImpl_MatchSubExp(t *testing.T) {
	// https://semver.org/
	semanticRegex := `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

	re := MustNew(semanticRegex)
	{
		v102, match := re.MatchSubExp("1.0.2")
		if !match {
			t.Error(match)
		}
		if v, ok := v102["major"]; !ok || v != "1" {
			t.Error(v, ok)
		}
		if v, ok := v102["minor"]; !ok || v != "0" {
			t.Error(v, ok)
		}
		if v, ok := v102["patch"]; !ok || v != "2" {
			t.Error(v, ok)
		}
		if v, ok := v102["prerelease"]; !ok || v != "" {
			t.Error(v, ok)
		}
	}

	{
		v123, match := re.MatchSubExp("1.2.3-beta+exp.sha.5114f85")
		if !match {
			t.Error(match)
		}
		if v, ok := v123["major"]; !ok || v != "1" {
			t.Error(v, ok)
		}
		if v, ok := v123["minor"]; !ok || v != "2" {
			t.Error(v, ok)
		}
		if v, ok := v123["patch"]; !ok || v != "3" {
			t.Error(v, ok)
		}
		if v, ok := v123["prerelease"]; !ok || v != "beta" {
			t.Error(v, ok)
		}
		if v, ok := v123["buildmetadata"]; !ok || v != "exp.sha.5114f85" {
			t.Error(v, ok)
		}
	}

	{
		_, match := re.MatchSubExp("1-0-2")
		if match {
			t.Error(match)
		}
	}
}
