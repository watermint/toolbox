package feed

import (
	"encoding/json"
	"github.com/mmcdole/gofeed"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Json struct {
	Url     mo_url.Url
	Compact bool
}

func (z *Json) Preset() {
}

func (z *Json) Exec(c app_control.Control) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(z.Url.Value())
	if err != nil {
		return err
	}
	if z.Compact {
		d, err := json.Marshal(feed)
		if err != nil {
			return err
		}
		ui_out.TextOut(c, string(d))
	} else {
		d, err := json.MarshalIndent(feed, "", "  ")
		if err != nil {
			return err
		}
		ui_out.TextOut(c, string(d))
	}
	return nil
}

func (z *Json) Test(c app_control.Control) error {
	feedUrl, err := mo_url.NewUrl("https://blog.golang.org/feed.atom")
	if err != nil {
		return err
	}
	return rc_exec.Exec(c, &Json{}, func(r rc_recipe.Recipe) {
		m := r.(*Json)
		m.Url = feedUrl
	})
}
