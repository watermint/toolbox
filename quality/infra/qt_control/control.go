package qt_control

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_feature_impl"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"io/ioutil"
)

func WithControl(f func(c app_control.Control) error) error {
	l := es_log.Default()
	home, err := ioutil.TempDir("", "control")
	if err != nil {
		l.Debug("unable to create home", es_log.Error(err))
		return err
	}
	wb, err := app_workspace.NewBundle(home, app_budget.BudgetUnlimited, es_log.LevelQuiet)
	if err != nil {
		l.Debug("unable to create bundle", es_log.Error(err))
		return err
	}
	defer wb.Close()

	com := app_opt.Default()
	fe := app_feature_impl.NewFeature(com, wb.Workspace())
	mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{})
	ui := app_ui.NewDiscard(mc, wb.Logger().Logger())
	ctl := app_control_impl.New(wb, ui, fe)

	return f(ctl)
}
