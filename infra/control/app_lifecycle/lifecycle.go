package app_lifecycle

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"time"
)

type MsgAppLifecycle struct {
	LifecycleWarnBestBefore     app_msg.Message
	LifecycleWarnExpired        app_msg.Message
	LifecycleErrorExpired       app_msg.Message
	LifecycleUpgradeInstruction app_msg.Message
}

var (
	MAppLifecycle = app_msg.Apply(&MsgAppLifecycle{}).(*MsgAppLifecycle)
)

type Lifecycle interface {
	TimeBestBefore() (bestBefore time.Time)
	TimeExpiration() (expired time.Time)
	IsExpired() (expired bool)
	IsBeyondBestBefore() (bestBefore bool)
	Verify(ui app_ui.UI)
}

func LifecycleControl() Lifecycle {
	return &lifecycleImpl{}
}

type lifecycleImpl struct {
}

func (z lifecycleImpl) binaryBuildTime() (buildTime time.Time) {
	t, err := time.Parse(time.RFC3339, app_definitions.BuildInfo.Timestamp)
	if err != nil {
		return time.Now()
	}
	return t
}

func (z lifecycleImpl) TimeBestBefore() (bestBefore time.Time) {
	return z.binaryBuildTime().Add(app_definitions.LifecycleExpirationWarning)
}

func (z lifecycleImpl) TimeExpiration() (expired time.Time) {
	return z.binaryBuildTime().Add(app_definitions.LifecycleExpirationCritical)
}

func (z lifecycleImpl) IsExpired() (expired bool) {
	return time.Now().After(z.TimeExpiration())
}

func (z lifecycleImpl) IsBeyondBestBefore() (bestBefore bool) {
	return time.Now().After(z.TimeBestBefore())
}

func (z lifecycleImpl) Verify(ui app_ui.UI) {
	if z.IsExpired() {
		switch app_definitions.LifecycleExpirationMode {
		case app_definitions.LifecycleExpirationShutdown:
			ui.Error(MAppLifecycle.LifecycleErrorExpired)
			ui.Info(MAppLifecycle.LifecycleUpgradeInstruction.With("UpgradeUrl", app_definitions.LifecycleUpgradeUrl))
			app_exit.Abort(app_exit.FailureBinaryExpired)
		case app_definitions.LifecycleExpirationWarningOnly:
			ui.Error(MAppLifecycle.LifecycleWarnExpired)
			ui.Info(MAppLifecycle.LifecycleUpgradeInstruction.With("UpgradeUrl", app_definitions.LifecycleUpgradeUrl))
		}
	} else if z.IsBeyondBestBefore() {
		ui.Error(MAppLifecycle.LifecycleWarnBestBefore.With("ExpirationDate", z.TimeExpiration().Format(time.RFC3339)))
		ui.Info(MAppLifecycle.LifecycleUpgradeInstruction.With("UpgradeUrl", app_definitions.LifecycleUpgradeUrl))
	}
}
