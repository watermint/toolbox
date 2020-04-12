package release

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
)

type Draft struct {
	Peer       gh_conn.ConnGithubRepo
	Owner      string
	Repository string
	Tag        string
	Name       string
	BodyFile   mo_path2.FileSystemPath
	Release    rp_model.RowReport
}

func (z *Draft) Preset() {
	z.Release.SetModel(&mo_release.Release{})
}

func (z *Draft) Exec(c app_control.Control) error {
	body, err := ioutil.ReadFile(z.BodyFile.Path())
	if err != nil {
		return err
	}
	if err := z.Release.Open(); err != nil {
		return err
	}
	svr := sv_release.New(z.Peer.Context(), z.Owner, z.Repository)
	rel, err := svr.CreateDraft(z.Tag, z.Name, string(body))
	if err != nil {
		return err
	}
	z.Release.Row(rel)
	return nil
}

func (z *Draft) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("github-release.txt", "This is test release")
	if err != nil {
		return err
	}
	defer func() {
		os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Draft{}, func(r rc_recipe.Recipe) {
		m := r.(*Draft)
		m.Tag = "0.0.2"
		m.Name = "0.0.2"
		m.Owner = "watermint"
		m.Repository = "toolbox_sandbox"
		m.BodyFile = mo_path2.NewFileSystemPath(f)
	})
}
