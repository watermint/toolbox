package download

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		url := "https://postman-echo.com/get?hello=world"
		ws := ctl.Workspace().Test()
		if err := os.MkdirAll(ws, 0755); err != nil {
			t.Error(err)
			return
		}

		p := filepath.Join(ws, "hello.json")
		if err := Download(ctl.Log(), url, p); err != nil {
			t.Error(err)
			return
		}

		r, err := ioutil.ReadFile(p)
		if err != nil {
			t.Error(err)
			return
		}

		type Hello struct {
			Hello string `json:"hello"`
		}
		type Response struct {
			Args *Hello `json:"args"`
		}
		res := &Response{}
		if err = json.Unmarshal(r, res); err != nil {
			t.Error(err)
			return
		}

		if res.Args.Hello != "world" {
			t.Error(res.Args.Hello)
		}
	})
}
