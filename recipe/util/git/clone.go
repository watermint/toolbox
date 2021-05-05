package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Clone struct {
	Url        string
	LocalPath  mo_path.FileSystemPath
	RemoteName string
	Reference  mo_string.OptionalString
}

func (z *Clone) Preset() {
	z.RemoteName = "origin"
}

func (z *Clone) Exec(c app_control.Control) error {
	l := c.Log()
	repo, err := git.PlainClone(z.LocalPath.Path(), false, &git.CloneOptions{
		URL:           z.Url,
		Auth:          nil,
		RemoteName:    z.RemoteName,
		Progress:      es_stdout.NewDefaultOut(c.Feature()),
		ReferenceName: plumbing.ReferenceName(z.Reference.Value()),
	})
	if err != nil {
		l.Debug("Unable to clone the repository", esl.Error(err))
		return err
	}
	head, err := repo.Head()
	if err != nil {
		l.Debug("Unable to retrieve head", esl.Error(err))
		return err
	}
	l.Info("Head", esl.Any("hash", head.Hash()), esl.Any("target", head.Target()), esl.Any("type", head.Type()))
	return nil
}

func (z *Clone) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("git", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.Exec(c, &Clone{}, func(r rc_recipe.Recipe) {
		m := r.(*Clone)
		m.Url = "https://github.com/watermint/toolbox_sandbox.git"
		m.LocalPath = mo_path.NewFileSystemPath(f)
	})
}
