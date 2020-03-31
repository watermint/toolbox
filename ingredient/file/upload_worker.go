package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_filecompare"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
)

type UploadWorker struct {
	dropboxBasePath string
	localBasePath   string
	localFilePath   string
	dbxEntry        mo_file.Entry
	ctx             api_context.Context
	ctl             app_control.Control
	up              sv_file_content.Upload
	estimateOnly    bool
	status          *UploadStatus
	upload          *Upload
}

func (z *UploadWorker) Exec() (err error) {
	ui := z.ctl.UI()
	upRow := &UploadRow{File: z.localFilePath}
	l := z.ctl.Log().With(
		zap.String("dropboxBasePath", z.dropboxBasePath),
		zap.String("localBasePath", z.localBasePath),
		zap.String("localFilePath", z.localFilePath),
	)
	l.Debug("Prepare upload")

	rel, err := ut_filepath.Rel(z.localBasePath, filepath.Dir(z.localFilePath))
	if err != nil {
		l.Debug("unable to calculate rel path", zap.Error(err))
		z.upload.Uploaded.Failure(err, upRow)
		z.status.error()
		return err
	}
	dp := mo_path.NewDropboxPath(z.dropboxBasePath)
	switch {
	case rel == ".":
		l.Debug("upload to base path")
	case strings.HasPrefix(rel, ".."):
		l.Debug("invalid rel path", zap.String("rel", rel))
		z.upload.Uploaded.Failure(errors.New("invalid path"), &UploadRow{File: z.localFilePath})
		z.status.error()
		return errors.New("invalid rel path")
	default:
		dp = dp.ChildPath(rel)
	}

	info, err := os.Lstat(z.localFilePath)
	if err != nil {
		z.upload.Uploaded.Failure(err, upRow)
		z.status.error()
		return err
	}
	upRow.Size = info.Size()

	// Verify proceed
	if z.dbxEntry != nil {
		same, err := ut_filecompare.Compare(l, z.localFilePath, info, z.dbxEntry)
		if err != nil {
			z.upload.Uploaded.Failure(err, upRow)
			z.status.error()
			return err
		}
		if same {
			//ui.Info("usecase.uc_file_upload.progress.skip", app_msg.P{
			//	"File": z.localFilePath,
			//})

			z.upload.Skipped.Skip(app_msg.M("usecase.uc_file_upload.skip.file_exists"), upRow)
			z.status.skip()
			return nil
		}
	}

	if z.estimateOnly {
		z.status.upload(info.Size(), z.upload.ChunkSizeKb)
		l.Debug("Skip upload (estimate only)")
		return nil
	}

	ui.Info(z.upload.ProgressUpload.With("File", z.localFilePath))

	var entry mo_file.Entry
	if z.upload.Overwrite {
		entry, err = z.up.Overwrite(dp, z.localFilePath)
		if err != nil {
			z.upload.Uploaded.Failure(err, upRow)
			z.status.error()
			return err
		}
	} else {
		entry, err = z.up.Add(dp, z.localFilePath)
		if err != nil {
			z.upload.Uploaded.Failure(err, upRow)
			z.status.error()
			return err
		}
	}
	z.upload.Uploaded.Success(upRow, entry.Concrete())
	z.status.upload(info.Size(), z.upload.ChunkSizeKb)
	return nil
}
