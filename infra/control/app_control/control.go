package app_control

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Control interface {
	// UI
	UI() app_ui.UI

	// Logger
	Log() es_log.Logger

	// HTTP Capture logger
	Capture() es_log.Logger

	// Workspace
	Workspace() app_workspace.Workspace

	// Message container
	Messages() app_msg_container.Container

	// Feature
	Feature() app_feature.Feature

	// Create new worker queue
	NewQueue() rc_worker.Queue

	// Workspace bundle
	WorkBundle() app_workspace.Bundle

	// Fork control instance with feature
	WithFeature(feature app_feature.Feature) Control

	// Fork control instance with UI
	WithUI(ui app_ui.UI) Control

	// Fork control with lang
	WithLang(targetLang string) Control

	// Fork control with bundle
	WithBundle(wb app_workspace.Bundle) Control
}
