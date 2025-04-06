package release

import (
	"os"

	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Draft struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkIrreversible
	Peer       gh_conn.ConnGithubRepo
	Owner      string
	Repository string
	Tag        string
	Name       string
	Branch     string
	BodyFile   mo_path2.FileSystemPath
	Release    rp_model.RowReport
}

func (z *Draft) Preset() {
	z.Release.SetModel(&mo_release.Release{})
}

func (z *Draft) Exec(c app_control.Control) error {
	body, err := os.ReadFile(z.BodyFile.Path())
	if err != nil {
		return err
	}
	if err := z.Release.Open(); err != nil {
		return err
	}
	svr := sv_release.New(z.Peer.Client(), z.Owner, z.Repository)
	rel, err := svr.CreateDraft(z.Tag, z.Name, string(body), z.Branch)
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
		m.Branch = "main"
		m.BodyFile = mo_path2.NewFileSystemPath(f)
	})
}
