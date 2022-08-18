package asset

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type UploadFile struct {
	File string `json:"file"`
}

type UploadTarget struct {
	Path    string
	Release *mo_release.Release
}

type Upload struct {
	rc_recipe.RemarkExperimental
	rc_recipe.RemarkIrreversible
	Asset      mo_path.ExistingFileSystemPath
	Owner      string
	Repository string
	Release    string
	Peer       gh_conn.ConnGithubRepo
	Uploads    rp_model.TransactionReport
}

func (z *Upload) Preset() {
	z.Uploads.SetModel(
		&UploadFile{},
		&mo_release_asset.Asset{},
		rp_model.HiddenColumns("result.id"),
	)
}

func (z *Upload) upload(target *UploadTarget, c app_control.Control, s eq_sequence.Stage) error {
	q := s.Get("upload")
	l := c.Log().With(
		esl.String("owner", z.Owner),
		esl.String("repo", z.Repository),
		esl.String("release", target.Release.Id),
		esl.String("path", target.Path))

	l.Debug("verify path")
	ls, err := os.Lstat(target.Path)
	if err != nil {
		return err
	}
	if ls.IsDir() {
		l.Debug("read a directory")
		entries, err := os.ReadDir(target.Path)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			l.Debug("Enqueue", esl.String("entry", entry.Name()))
			q.Enqueue(&UploadTarget{
				Path:    filepath.Join(target.Path, entry.Name()),
				Release: target.Release,
			})
		}
		return nil
	}

	uf := &UploadFile{
		File: target.Path,
	}
	asset, err := sv_release_asset.New(z.Peer.Context(), z.Owner, z.Repository, target.Release.Id).Upload(mo_path.NewExistingFileSystemPath(target.Path))
	if err != nil {
		z.Uploads.Failure(err, uf)
		return err
	}
	z.Uploads.Success(uf, asset)
	return nil
}

func (z *Upload) Exec(c app_control.Control) error {
	if err := z.Uploads.Open(); err != nil {
		return err
	}

	rel, err := sv_release.New(z.Peer.Context(), z.Owner, z.Repository).Get(z.Release)
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("upload", z.upload, c, s)
		q := s.Get("upload")
		q.Enqueue(&UploadTarget{
			Path:    z.Asset.Path(),
			Release: rel,
		})
	})
	return nil
}

func (z *Upload) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("github-release-asset", true)
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Upload{}, func(r rc_recipe.Recipe) {
		m := r.(*Upload)
		m.Owner = "watermint"
		m.Repository = "toolbox_sandbox"
		m.Release = "0.0.2"
		m.Asset = mo_path.NewExistingFileSystemPath(f)
	})
}
