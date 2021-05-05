package resources

import (
	"embed"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_resource"
	"strings"
	"time"
)

//go:embed templates/*
var resTemplates embed.FS

//go:embed messages/*
var resMessages embed.FS

//go:embed web/*
var resWeb embed.FS

//go:embed keys/*
var resKeys embed.FS

//go:embed images/*
var resImages embed.FS

//go:embed data/*
var resData embed.FS

//go:embed release/release
var resRelease string

//go:embed release/release_notes
var resReleaseNotes string

//go:embed release/build/*
var resBuildInfo embed.FS

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
		es_resource.NewSecureResource("messages", resMessages),
		es_resource.NewResource("web", resWeb),
		es_resource.NewSecureResource("keys", resKeys),
		es_resource.NewResource("images", resImages),
		es_resource.NewSecureResource("data", resData),
	)
}

// Release release number (major version only)
func Release() string {
	return strings.TrimSpace(resRelease)
}

// ReleaseNotes release notes for the current release.
func ReleaseNotes() string {
	return resReleaseNotes
}

func Build() BuildInfo {
	fallback := func() BuildInfo {
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

	infoJson, err := resBuildInfo.ReadFile("release/build/info.json")
	if err != nil {
		return fallback()
	}

	info := BuildInfo{}
	if err := json.Unmarshal(infoJson, &info); err != nil {
		return fallback()
	}

	return info
}
