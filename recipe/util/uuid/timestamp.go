package uuid

import (
	"github.com/watermint/toolbox/essentials/strings/es_uuid"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"strings"
	"time"
)

type Timestamp struct {
	Uuid string
}

func (z *Timestamp) Preset() {
}

func (z *Timestamp) Exec(c app_control.Control) error {
	u, oc := es_uuid.Parse(strings.TrimSpace(z.Uuid))
	if oc.IsError() {
		return oc.Cause()
	}
	ts, err := es_uuid.TimestampFromUUIDV7(u)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, ts.Format(time.RFC3339Nano))
	return nil
}

func (z *Timestamp) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Timestamp{}, func(r rc_recipe.Recipe) {
		m := r.(*Timestamp)
		m.Uuid = es_uuid.NewV7().String()
	})
}
