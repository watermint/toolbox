package resources

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"strings"
	"time"
)

type BuildInfo struct {
	Version    string `json:"version"`
	Hash       string `json:"hash"`
	Branch     string `json:"branch"`
	Timestamp  string `json:"timestamp"`
	Year       int    `json:"year"`
	Zap        string `json:"zap"`
	Xap        string `json:"xap"`
	Production bool   `json:"production"`
}

var (
	CurrentBundle = NewBundle()
)

func NewBundle() es_resource.Bundle {
	return es_resource.New(
		es_resource.NewResource("templates", resTemplates),
		es_resource.NewNonTraversableResource("messages", resMessages),
		es_resource.NewResource("web", resWeb),
		es_resource.NewNonTraversableResource("keys", resKeys),
		es_resource.NewResource("images", resImages),
		es_resource.NewNonTraversableResource("data", resData),
		es_resource.NewNonTraversableResource("build", resBuildInfo),
		es_resource.NewNonTraversableResource("release", resRelease),
	)
}

func CoreRelease() string {
	rel, err := NewBundle().Release().Bytes("release")
	if err != nil {
		panic("`release` resource not found: " + err.Error())
	}
	return strings.TrimSpace(string(rel))
}

// Release release number (major version only) from the resource
func Release() string {
	rel, err := CurrentBundle.Release().Bytes("release")
	if err != nil {
		panic("`release` resource not found: " + err.Error())
	}
	return strings.TrimSpace(string(rel))
}

// ReleaseNotes release notes for the current release.
func ReleaseNotes() string {
	rel, err := CurrentBundle.Release().Bytes("release_notes")
	if err != nil {
		panic("`release` resource not found: " + err.Error())
	}
	return string(rel)
}

func Build() BuildInfo {
	return BuildFromResource(CurrentBundle.Build())
}

func buildInfoFallback() BuildInfo {
	return BuildInfo{
		Version:    "",
		Hash:       "",
		Branch:     "",
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Year:       time.Now().UTC().Year(),
		Zap:        "",
		Xap:        "",
		Production: false,
	}
}

func BuildFromResource(res es_resource.Resource) BuildInfo {
	infoJson, err := res.Bytes("info.json")
	if err != nil {
		return buildInfoFallback()
	}

	info := BuildInfo{}
	if err := json.Unmarshal(infoJson, &info); err != nil {
		return buildInfoFallback()
	}

	return info
}
