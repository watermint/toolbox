package uuid

import (
	"errors"
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
	switch u.Version() {
	case es_uuid.Version1, es_uuid.Version2, es_uuid.Version6:
		return errors.New("the command currently support only UUID v7")
	case es_uuid.Version3, es_uuid.Version4, es_uuid.Version5:
		return errors.New("given UUID is not a time-based UUID")
	case es_uuid.Version7:
		break
	case es_uuid.Version8:
		return errors.New("given UUID is vendor specific UUID, not supported")
	default:
		return errors.New("unsupported UUID version")
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
