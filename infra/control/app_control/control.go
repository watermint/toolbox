package app_control

import (
	"database/sql"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"gorm.io/gorm"
)

type Control interface {
	// UI returns UI instance
	UI() app_ui.UI

	// Log returns logger instance
	Log() esl.Logger

	// Capture HTTP Capture logger
	Capture() esl.Logger

	// Workspace returns workspace instance
	Workspace() app_workspace.Workspace

	// Messages Message container
	Messages() app_msg_container.Container

	// Feature returns feature flags
	Feature() app_feature.Feature

	// NewQueue Create new queue definition
	NewQueue() eq_queue.Definition

	// Sequence Async queue sequence
	Sequence() eq_sequence.Sequence

	// AuthRepository returns auth repository
	AuthRepository() api_auth.Repository

	// NewKvs Create new KVS. The caller must close the storage before exit.
	NewKvs(name string) (kvs kv_storage.Storage, err error)

	// NewKvsFactory Create new KVS factory. The caller must close the factory before exit.
	NewKvsFactory() (factory kv_storage.Factory)

	// NewDatabase Create new database. The caller must close the database before exit.
	NewDatabase(name string) (db *sql.DB, path string, err error)

	// NewOrm Create new ORM instance. The caller must close the ORM before exit.
	NewOrm(path string) (db *gorm.DB, err error)

	// NewOrmOnMemory Create new ORM instance on memory. The caller must close the ORM before exit.
	NewOrmOnMemory() (db *gorm.DB, err error)

	// WorkBundle Workspace bundle
	WorkBundle() app_workspace.Bundle

	// WithFeature Fork control instance with feature
	WithFeature(feature app_feature.Feature) Control

	// WithUI Fork control instance with UI
	WithUI(ui app_ui.UI) Control

	// WithLang Fork control with lang
	WithLang(targetLang string) Control

	// WithBundle Fork control with bundle
	WithBundle(wb app_workspace.Bundle) Control
}

type ControlCloser interface {
	Control

	Close()
}
