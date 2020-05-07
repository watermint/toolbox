package asset

import (
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/api/gh_context"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MsgUp struct {
	ProgressUpload    app_msg.Message
	ErrorUnableToRead app_msg.Message
}

var (
	MUp = app_msg.Apply(&MsgUp{}).(*MsgUp)
)

type UploadWorker struct {
	file       mo_path.ExistingFileSystemPath
	owner      string
	repository string
	release    *mo_release.Release
	ctx        gh_context.Context
	ctl        app_control.Control
	uploads    rp_model.TransactionReport
	queue      rc_worker.Queue
}

func (z *UploadWorker) Exec() error {
	l := z.ctl.Log().With(
		es_log.String("owner", z.owner),
		es_log.String("repo", z.repository),
		es_log.String("release", z.release.Id),
		es_log.String("path", z.file.Path()))
	ui := z.ctl.UI()

	l.Debug("verify path")
	ls, err := os.Lstat(z.file.Path())
	if err != nil {
		ui.Error(MUp.ErrorUnableToRead.With("Error", err).With("Path", z.file.Path()))
		return err
	}
	if ls.IsDir() {
		l.Debug("read a directory")
		entries, err := ioutil.ReadDir(z.file.Path())
		if err != nil {
			return err
		}
		for _, entry := range entries {
			l.Debug("Enqueue", es_log.String("entry", entry.Name()))
			z.queue.Enqueue(&UploadWorker{
				file:       mo_path.NewExistingFileSystemPath(filepath.Join(z.file.Path(), entry.Name())),
				owner:      z.owner,
				repository: z.repository,
				release:    z.release,
				ctx:        z.ctx,
				ctl:        z.ctl,
				uploads:    z.uploads,
				queue:      z.queue,
			})
		}
		return nil
	}

	ui.Progress(MUp.ProgressUpload.With("File", z.file.Path()))

	uf := &UploadFile{
		File: z.file.Path(),
	}
	asset, err := sv_release_asset.New(z.ctx, z.owner, z.repository, z.release.Id).Upload(z.file)
	if err != nil {
		z.uploads.Failure(err, uf)
		return err
	}
	z.uploads.Success(uf, asset)
	return nil
}

type UploadFile struct {
	File string `json:"file"`
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

func (z *Upload) Exec(c app_control.Control) error {
	if err := z.Uploads.Open(); err != nil {
		return err
	}

	rel, err := sv_release.New(z.Peer.Context(), z.Owner, z.Repository).Get(z.Release)
	if err != nil {
		return err
	}

	q := c.NewQueue()
	q.Enqueue(&UploadWorker{
		file:       z.Asset,
		owner:      z.Owner,
		repository: z.Repository,
		release:    rel,
		ctx:        z.Peer.Context(),
		ctl:        c,
		uploads:    z.Uploads,
		queue:      q,
	})
	return q.Wait()
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
