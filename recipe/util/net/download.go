package net

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Download struct {
	Url mo_url.Url
	Out mo_path.FileSystemPath
}

func (z *Download) Preset() {
}

func (z *Download) Exec(c app_control.Control) error {
	return es_download.Download(c.Log(), z.Url.Value(), z.Out.Path())
}

func (z *Download) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("download", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	dlPath := filepath.Join(p, "hello.json")
	err = rc_exec.Exec(c, &Download{}, func(r rc_recipe.Recipe) {
		m := r.(*Download)
		m.Out = mo_path.NewFileSystemPath(dlPath)
		m.Url, _ = mo_url.NewUrl("https://postman-echo.com/get?hello=world")
	})
	if err != nil {
		return err
	}

	content, err := os.ReadFile(dlPath)
	if err != nil {
		return err
	}
	r := gjson.GetBytes(content, "args.hello")
	if r.String() != "world" {
		return errors.New("invalid download")
	}
	return nil
}
