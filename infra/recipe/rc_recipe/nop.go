package rc_recipe

import "github.com/watermint/toolbox/infra/control/app_control"

type Nop struct {
}

func (z Nop) Preset() {
}

func (z Nop) Exec(c app_control.Control) error {
	return nil
}

func (z Nop) Test(c app_control.Control) error {
	return nil
}
