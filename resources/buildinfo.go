package resources

import (
	"embed"
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

func NewBundle() es_resource.Bundle {
	return es_resource.New(
		es_resource.NewResource("templates", resTemplates),
		es_resource.NewNonTraversableResource("messages", resMessages),
		es_resource.NewResource("web", resWeb),
		es_resource.NewNonTraversableResource("keys", resKeys),
		es_resource.NewResource("images", resImages),
		es_resource.NewNonTraversableResource("data", resData),
		es_resource.NewNonTraversableResource("build", resBuildInfo),
	)
}

// Release release number (major version only) from the resource
func Release() string {
	return strings.TrimSpace(resRelease)
}

// ReleaseNotes release notes for the current release.
func ReleaseNotes() string {
	return resReleaseNotes
}

func Build() BuildInfo {
	return BuildFromResource(resBuildInfo)
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

func BuildFromResource(res embed.FS) BuildInfo {
	infoJson, err := res.ReadFile("build/info.json")
	if err != nil {
		return buildInfoFallback()
	}

	info := BuildInfo{}
	if err := json.Unmarshal(infoJson, &info); err != nil {
		return buildInfoFallback()
	}

	return info
}
