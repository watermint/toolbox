package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type UploadVO struct {
	Peer        app_conn.ConnUserFile
	LocalPath   string
	DropboxPath string
	Overwrite   bool
	ChunkSize   int
}

const (
	reportUpload = "uploaded"
)

type UploadWorker struct {
	dropboxBasePath string
	localBasePath   string
	localFilePath   string
	ctx             api_context.Context
	ctl             app_control.Control
	up              sv_file_content.Upload
	rep             rp_model.Report
	overwrite       bool
}

type UploadRow struct {
	File string `json:"file"`
}

func (z *UploadWorker) Exec() (err error) {
	ui := z.ctl.UI()
	ui.Info("recipe.file.upload.progress", app_msg.P{
		"File": z.localFilePath,
	})
	upRow := &UploadRow{File: z.localFilePath}
	l := z.ctl.Log().With(
		zap.String("dropboxBasePath", z.dropboxBasePath),
		zap.String("localBasePath", z.localBasePath),
		zap.String("localFilePath", z.localFilePath),
	)

	rel, err := filepath.Rel(z.localBasePath, filepath.Dir(z.localFilePath))
	if err != nil {
		l.Debug("unable to calculate rel path", zap.Error(err))
		z.rep.Failure(err, upRow)
		return err
	}
	dp := mo_path.NewPath(z.dropboxBasePath)
	switch {
	case rel == ".":
		l.Debug("upload to base path")
	case strings.HasPrefix(rel, ".."):
		l.Debug("invalid rel path", zap.String("rel", rel))
		z.rep.Failure(errors.New("invalid path"), &UploadRow{File: z.localFilePath})
		return errors.New("invalid rel path")
	default:
		dp = dp.ChildPath(filepath.ToSlash(rel))
	}
	info, err := os.Lstat(z.localFilePath)
	if err != nil {
		z.rep.Failure(err, upRow)
		return err
	}

	meta, err := sv_file.NewFiles(z.ctx).Resolve(sv_file_content.UploadPath(dp, info))
	if err != nil {
		l.Debug("Unable to resolve path", zap.Error(err))
	} else {
		f := meta.Concrete()
		if f.Size == info.Size() {
			hash, err := api_util.ContentHash(z.localFilePath)
			if err != nil {
				z.rep.Failure(err, upRow)
				return err
			}
			if f.ContentHash == hash {
				z.rep.Skip(app_msg.M("recipe.file.upload.skip.file_exists"), upRow)
				return nil
			}
		}
	}

	var entry mo_file.Entry
	if z.overwrite {
		entry, err = z.up.Overwrite(dp, z.localFilePath)
		if err != nil {
			z.rep.Failure(err, upRow)
			return err
		}
	} else {
		entry, err = z.up.Add(dp, z.localFilePath)
		if err != nil {
			z.rep.Failure(err, upRow)
			return err
		}
	}
	z.rep.Success(upRow, entry.Concrete())
	return nil
}

type Upload struct {
}

func (z *Upload) Console() {
}

func (z *Upload) Requirement() app_vo.ValueObject {
	return &UploadVO{
		ChunkSize: 150 * 1024,
	}
}

func (z *Upload) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UploadVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportUpload)
	if err != nil {
		return err
	}
	defer rep.Close()

	up := sv_file_content.NewUpload(ctx, sv_file_content.ChunkSize(int64(vo.ChunkSize*1024)))
	q := k.NewQueue()

	info, err := os.Lstat(vo.LocalPath)
	if err != nil {
		return err
	}

	var scanFolder func(path string) error
	scanFolder = func(path string) error {
		entries, err := ioutil.ReadDir(path)
		if err != nil {
			return err
		}
		var lastErr error
		for _, e := range entries {
			p := filepath.Join(path, e.Name())
			if api_util.IsFileNameIgnored(p) {
				continue
			}
			if e.IsDir() {
				lastErr = scanFolder(e.Name())
			} else {
				q.Enqueue(&UploadWorker{
					dropboxBasePath: vo.DropboxPath,
					localBasePath:   vo.LocalPath,
					localFilePath:   p,
					ctl:             k.Control(),
					ctx:             ctx,
					up:              up,
					rep:             rep,
					overwrite:       vo.Overwrite,
				})
			}
		}
		return lastErr
	}

	var lastErr error
	if info.IsDir() {
		lastErr = scanFolder(vo.LocalPath)
	} else {
		q.Enqueue(&UploadWorker{
			dropboxBasePath: vo.DropboxPath,
			localBasePath:   vo.LocalPath,
			localFilePath:   vo.LocalPath,
			ctl:             k.Control(),
			up:              up,
			rep:             rep,
			overwrite:       vo.Overwrite,
		})
	}

	q.Wait()

	return lastErr
}

func (z *Upload) Test(c app_control.Control) error {
	l := c.Log()
	fileCandidates := []string{"README.md", "upload.go", "upload_test.go"}
	file := ""
	for _, f := range fileCandidates {
		if _, err := os.Lstat(f); err == nil {
			file = f
			break
		}
	}
	if file == "" {
		l.Warn("No file to upload")
		return qt_test.NotEnoughResource()
	}

	{
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + app_test.TestTeamFolderName,
			Overwrite:   true,
		}
		if !app_test.ApplyTestPeers(c, vo) {
			return qt_test.NotEnoughResource()
		}
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}

	// Chunked
	{
		vo := &UploadVO{
			LocalPath:   file,
			DropboxPath: "/" + app_test.TestTeamFolderName,
			Overwrite:   true,
			ChunkSize:   1,
		}
		if !app_test.ApplyTestPeers(c, vo) {
			return qt_test.NotEnoughResource()
		}
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err != nil {
			return err
		}
	}
	return nil
}

func (z *Upload) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportUpload,
			rp_model.TransactionHeader(&UploadRow{}, &mo_file.ConcreteEntry{}),
			rp_model.HiddenColumns(
				"result.id",
				"result.tag",
			)),
	}
}
