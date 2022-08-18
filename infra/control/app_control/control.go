package app_control

import (
	"github.com/watermint/toolbox/essentials/cache"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Control interface {
	// UI
	UI() app_ui.UI

	// Logger
	Log() esl.Logger

	// HTTP Capture logger
	Capture() esl.Logger

	// Workspace
	Workspace() app_workspace.Workspace

	// Message container
	Messages() app_msg_container.Container

	// Feature
	Feature() app_feature.Feature

	// Create new queue definition
	NewQueue() eq_queue.Definition

	// Async queue sequence
	Sequence() eq_sequence.Sequence

	// Get or create new cache
	NewCache(namespace, name string) cache.Cache

	// Create new KVS. The caller must close the storage before exit.
	NewKvs(name string) (kvs kv_storage.Storage, err error)

	// Create new KVS factory. The caller must close the factory before exit.
	NewKvsFactory() (factory kv_storage.Factory)

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

type ControlCloser interface {
	Control

	Close()
}
