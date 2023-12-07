package release

import (
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"strings"
)

const (
	AssetPlatformUnknown      = "unknown"
	AssetPlatformMacIntel     = "mac-intel"
	AssetPlatformMacArm       = "mac-arm"
	AssetPlatformLinuxIntel   = "linux-intel"
	AssetPlatformLinuxArm     = "linux-arm"
	AssetPlatformWindowsIntel = "win-intel"
)

func IdentifyPlatform(asset *mo_release_asset.Asset) string {
	name := strings.ToLower(asset.Name)
	if !strings.HasPrefix(name, "tbx") {
		return AssetPlatformUnknown
	}

	switch {
	case strings.HasSuffix(name, "mac-intel.zip"),
		strings.HasSuffix(name, "mac-amd64.zip"),
		strings.HasSuffix(name, "mac-x86_64.zip"):
		return AssetPlatformMacIntel

	case strings.HasSuffix(name, "mac-applesilicon.zip"),
		strings.HasSuffix(name, "mac-arm64.zip"),
		strings.HasSuffix(name, "mac-arm.zip"):
		return AssetPlatformMacArm

	case strings.HasSuffix(name, "linux-intel.zip"),
		strings.HasSuffix(name, "linux-amd64.zip"),
		strings.HasSuffix(name, "linux-x86_64.zip"):
		return AssetPlatformLinuxIntel

	case strings.HasSuffix(name, "linux-arm.zip"),
		strings.HasSuffix(name, "linux-arm64.zip"),
		strings.HasSuffix(name, "linux-arm.zip"):
		return AssetPlatformLinuxArm

	case strings.HasSuffix(name, "win.zip"),
		strings.HasSuffix(name, "win-intel.zip"),
		strings.HasSuffix(name, "win-amd64.zip"),
		strings.HasSuffix(name, "win-x86_64.zip"):
		return AssetPlatformWindowsIntel
	}

	return AssetPlatformUnknown
}
