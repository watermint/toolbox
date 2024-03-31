package dc_version

import (
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_lifecycle"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/resources"
	"runtime"
	"strconv"
	"time"
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
		Version:   app_definitions2.Name,
	})
	components = append(components, &VersionInfo{
		Key:       "app.version",
		Component: app_definitions2.Name,
		Version:   app_definitions2.BuildId,
	})
	components = append(components, &VersionInfo{
		Key:       "app.hash",
		Component: ui.Text(MVersionInfo.HeaderAppHash),
		Version:   app_definitions2.BuildInfo.Hash,
	})
	components = append(components, &VersionInfo{
		Key:       "app.branch",
		Component: ui.Text(MVersionInfo.HeaderBranch),
		Version:   app_definitions2.BuildInfo.Branch,
	})
	components = append(components, &VersionInfo{
		Key:       "app.production",
		Component: ui.Text(MVersionInfo.HeaderProduction),
		Version:   strconv.FormatBool(app_definitions2.BuildInfo.Production),
	})
	components = append(components, &VersionInfo{
		Key:       "core.release",
		Component: app_definitions2.Pkg,
		Version:   resources.CoreRelease(),
	})
	components = append(components, &VersionInfo{
		Key:       "build.time",
		Component: ui.Text(MVersionInfo.HeaderBuildTime),
		Version:   app_definitions2.BuildInfo.Timestamp,
	})
	components = append(components, &VersionInfo{
		Key:       "go.version",
		Component: ui.Text(MVersionInfo.HeaderGoVersion),
		Version:   runtime.Version(),
	})
	components = append(components, &VersionInfo{
		Key:       "lifecycle.bestbefore",
		Component: ui.Text(MVersionInfo.HeaderLifecycle),
		Version:   app_lifecycle.LifecycleControl().TimeBestBefore().Format(time.RFC3339),
	})
	components = append(components, &VersionInfo{
		Key:       "lifecycle.expiration",
		Component: ui.Text(MVersionInfo.HeaderLifecycle),
		Version:   app_lifecycle.LifecycleControl().TimeExpiration().Format(time.RFC3339),
	})

	return components
}
