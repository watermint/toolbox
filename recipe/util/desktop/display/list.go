package display

import (
	"github.com/kbinani/screenshot"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Display struct {
	Id     int `json:"id"`
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type List struct {
	Displays rp_model.RowReport
}

func (z *List) Preset() {
	z.Displays.SetModel(&Display{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Displays.Open(); err != nil {
		return err
	}
	for i := 0; i < screenshot.NumActiveDisplays(); i++ {
		d := screenshot.GetDisplayBounds(i).Bounds()
		z.Displays.Row(&Display{
			Id:     i,
			X:      d.Min.X,
			Y:      d.Min.Y,
			Width:  d.Size().X,
			Height: d.Size().Y,
		})
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
