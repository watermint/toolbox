package dc_version

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/resources"
	"runtime"
	"strconv"
)

type MsgVersionInfo struct {
	HeaderAppHash    app_msg.Message
	HeaderBuildTime  app_msg.Message
	HeaderBranch     app_msg.Message
	HeaderProduction app_msg.Message
	HeaderGoVersion  app_msg.Message
	HeaderLifecycle  app_msg.Message
}

var (
	MVersionInfo = app_msg.Apply(&MsgVersionInfo{}).(*MsgVersionInfo)
)

type VersionInfo struct {
	Key       string `json:"key"`
	Component string `json:"component"`
	Version   string `json:"version"`
}

func VersionComponents(ui app_ui.UI) []*VersionInfo {
	components := make([]*VersionInfo, 0)
	components = append(components, &VersionInfo{
		Key:       "app.name",
		Component: "app.name",
		Version:   app_definitions.Name,
	})
	components = append(components, &VersionInfo{
		Key:       "app.version",
		Component: app_definitions.Name,
		Version:   app_definitions.BuildId,
	})
	components = append(components, &VersionInfo{
		Key:       "app.hash",
		Component: ui.Text(MVersionInfo.HeaderAppHash),
		Version:   app_definitions.BuildInfo.Hash,
	})
	components = append(components, &VersionInfo{
		Key:       "app.branch",
		Component: ui.Text(MVersionInfo.HeaderBranch),
		Version:   app_definitions.BuildInfo.Branch,
	})
	components = append(components, &VersionInfo{
		Key:       "app.production",
		Component: ui.Text(MVersionInfo.HeaderProduction),
		Version:   strconv.FormatBool(app_definitions.BuildInfo.Production),
	})
	components = append(components, &VersionInfo{
		Key:       "core.release",
		Component: app_definitions.CorePkg,
		Version:   resources.CoreRelease(),
	})
	components = append(components, &VersionInfo{
		Key:       "build.time",
		Component: ui.Text(MVersionInfo.HeaderBuildTime),
		Version:   app_definitions.BuildInfo.Timestamp,
	})
	components = append(components, &VersionInfo{
		Key:       "go.version",
		Component: ui.Text(MVersionInfo.HeaderGoVersion),
		Version:   runtime.Version(),
	})

	return components
}
