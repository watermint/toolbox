package app_control_impl

import (
	"database/sql"
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/database/orm"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_error"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_queue"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"gorm.io/gorm"
	"path/filepath"
)

func New(wb app_workspace.Bundle, ui app_ui.UI, feature app_feature.Feature, seq eq_sequence.Sequence, authRepo api_auth.Repository, er app_error.ErrorReport) app_control.Control {
	return &ctlImpl{
		seq:         seq,
		wb:          wb,
		ui:          ui,
		feature:     feature,
		errorReport: er,
		authRepo:    authRepo,
	}
}

func ForkQuiet(ctl app_control.Control, name string) (app_control.Control, error) {
	wb, err := app_workspace.ForkBundleWithLevel(ctl.WorkBundle(), name, esl.LevelQuiet)
	if err != nil {
		return nil, err
	}
	lg := wb.Logger().Logger()
	qui := app_ui.NewDiscard(ctl.Messages(), lg)
	qfe := ctl.Feature().AsQuiet()
	cc, ok := ctl.(ForkableControl)
	if !ok {
		return nil, errors.New("this control is not cloneable")
	}
	seq, er := app_queue.NewSequence(lg, qfe, qui, wb)
	fc := cc.WithSequence(seq).WithErrorReport(er).WithBundle(wb).WithFeature(qfe).WithUI(qui)
	if err := er.Up(fc); err != nil {
		return nil, err
	}
	return fc, nil
}

func WithForkedQuiet(ctl app_control.Control, name string, f func(c app_control.Control) error) error {
	cf, err := ForkQuiet(ctl, name)
	if err != nil {
		return err
	}
	defer func() {
		_ = cf.WorkBundle().Close()
	}()
	return f(cf)
}

type ForkableControl interface {
	app_control.Control

	// Fork control with error report
	WithErrorReport(er app_error.ErrorReport) ForkableControl

	// Fork control with sequence
	WithSequence(seq eq_sequence.Sequence) ForkableControl
}

type ctlImpl struct {
	feature     app_feature.Feature
	ui          app_ui.UI
	wb          app_workspace.Bundle
	seq         eq_sequence.Sequence
	errorReport app_error.ErrorReport
	authRepo    api_auth.Repository
}

func (z ctlImpl) NewOrm(path string) (db *gorm.DB, err error) {
	return orm.NewOrm(z.Log(), path)
}

func (z ctlImpl) NewOrmOnMemory() (db *gorm.DB, err error) {
	return orm.NewOrmOnMemory(z.Log())
}

func (z ctlImpl) NewDatabase(name string) (db *sql.DB, path string, err error) {
	l := z.Log().With(esl.String("name", name))
	dbPath := filepath.Join(z.Workspace().Database(), es_filepath.Escape(name))
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		l.Debug("Unable to open the database", esl.Error(err), esl.String("path", dbPath))
		return nil, "", err
	}
	return db, dbPath, nil
}

func (z ctlImpl) AuthRepository() api_auth.Repository {
	return z.authRepo
}

func (z ctlImpl) NewQueue() eq_queue.Definition {
	return app_queue.NewQueue(z.Log(), z.Feature(), z.WorkBundle())
}

func (z ctlImpl) WithSequence(seq eq_sequence.Sequence) ForkableControl {
	z.seq = seq
	return z
}

func (z ctlImpl) WithErrorReport(er app_error.ErrorReport) ForkableControl {
	z.errorReport = er
	return z
}

func (z ctlImpl) NewKvsFactory() (factory kv_storage.Factory) {
	factory = kv_storage_impl.NewFactory(z.wb.Workspace().KVS(), z.wb.Logger().Logger())
	return
}

func (z ctlImpl) NewKvs(name string) (kvs kv_storage.Storage, err error) {
	kvs0 := kv_storage_impl.NewBitCask(name, z.wb.Logger().Logger())
	kvs = kvs0
	err = kvs0.Open(z.wb.Workspace().KVS())
	return
}

func (z ctlImpl) Close() {
	z.errorReport.Down()
}

func (z ctlImpl) Sequence() eq_sequence.Sequence {
	return z.seq
}

func (z ctlImpl) WithLang(targetLang string) app_control.Control {
	l := z.Log().With(esl.String("targetLang", targetLang))
	usrLang := es_lang.Select(targetLang, es_lang.Supported)
	priority := es_lang.Priority(usrLang)
	containers := make(map[es_lang.Iso639One]app_msg_container.Container)

	for _, la := range priority {
		mc, err := app_msg_container_impl.NewSingle(la)
		if err != nil {
			l.Debug("Unable to load resource for language",
				esl.String("la", la.String()),
				esl.Error(err))
			return z
		}
		containers[la.Code()] = mc
	}

	mc := app_msg_container_impl.NewMultilingual(priority, containers)
	z.ui = z.ui.WithContainer(mc)
	return z
}

func (z ctlImpl) WithFeature(feature app_feature.Feature) app_control.Control {
	z.feature = feature
	return z
}

func (z ctlImpl) WithUI(ui app_ui.UI) app_control.Control {
	z.ui = ui
	return z
}

func (z ctlImpl) WithBundle(wb app_workspace.Bundle) app_control.Control {
	z.wb = wb
	return z
}

func (z ctlImpl) Feature() app_feature.Feature {
	return z.feature
}

func (z ctlImpl) Messages() app_msg_container.Container {
	return z.ui.Messages()
}

func (z ctlImpl) Workspace() app_workspace.Workspace {
	return z.wb.Workspace()
}

func (z ctlImpl) WorkBundle() app_workspace.Bundle {
	return z.wb
}

func (z ctlImpl) UI() app_ui.UI {
	return z.ui
}

func (z ctlImpl) Log() esl.Logger {
	return z.wb.Logger().Logger()
}

func (z ctlImpl) Capture() esl.Logger {
	return z.wb.Capture().Logger()
}
