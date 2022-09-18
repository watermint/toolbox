package pack

import (
	"archive/zip"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"io"
	"os"
	"path/filepath"
	"time"
)

type RemoteZipEntry struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Remote struct {
	Peer             dbx_conn.ConnScopedIndividual
	DropboxPath      mo_path.DropboxPath
	LocalPath        mo_path2.FileSystemPath
	OperationLog     rp_model.TransactionReport
	ProgressDownload app_msg.Message
	ProgressCompress app_msg.Message
}

func (z *Remote) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
	)
	z.OperationLog.SetModel(
		&mo_file.ConcreteEntry{},
		&RemoteZipEntry{},
		rp_model.HiddenColumns(
			"input.id",
			"input.tag",
			"input.path_lower",
			"input.content_hash",
			"input.shared_folder_id",
			"input.parent_shared_folder_id",
			"input.server_modified",
			"input.revision",
		),
	)
}

func (z *Remote) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	zf, err := os.Create(z.LocalPath.Path())
	if err != nil {
		return err
	}
	defer func() {
		_ = zf.Close()
	}()
	zw := zip.NewWriter(zf)
	defer func() {
		_ = zw.Close()
	}()

	sdl := sv_file_content.NewDownload(z.Peer.Client())
	var dlErr, zipErr error
	var dlFile mo_file.Entry
	var dlCopy mo_path2.FileSystemPath

	l := c.Log()
	err = sv_file.NewFiles(z.Peer.Client()).ListEach(z.DropboxPath, func(entry mo_file.Entry) {
		ll := l.With(esl.String("path", entry.PathDisplay()))
		ui := c.UI()

		fEntry, isFile := entry.File()
		if !isFile {
			return
		}

		relPath, rErr := es_filepath.Rel(z.DropboxPath.Path(), entry.PathDisplay())
		if rErr != nil {
			ll.Debug("Unable to calculate rel path", esl.Error(rErr))
			z.OperationLog.Failure(rErr, entry.Concrete())
			return
		}
		ll = ll.With(esl.String("relPath", relPath))

		ui.Progress(z.ProgressDownload.With("Path", entry.PathDisplay()).With("Size", fEntry.Size))
		dlFile, dlCopy, dlErr = sdl.Download(mo_path.NewDropboxPath(entry.PathDisplay()))
		if dlErr != nil {
			ll.Debug("Unable to download", esl.Error(dlErr))
			z.OperationLog.Failure(dlErr, entry.Concrete())
			return
		}
		ll.Debug("Download completed", esl.String("dlCopy", dlCopy.Path()))

		var zipPath string
		if relPath == "." {
			zipPath = entry.Name()
		} else {
			zipPath = filepath.ToSlash(filepath.Join(relPath))
		}

		entryTs, tsErr := dbx_util.Parse(fEntry.ClientModified)
		if tsErr != nil {
			ll.Debug("Unable to parse timestamp", esl.Error(tsErr), esl.String("ts", fEntry.ClientModified))
			entryTs = time.Now()
		}

		ui.Progress(z.ProgressCompress.With("Path", zipPath))
		zef, zipErr := zw.CreateHeader(&zip.FileHeader{
			Name:     zipPath,
			Method:   zip.Deflate,
			Modified: entryTs,
		})
		if zipErr != nil {
			ll.Debug("Unable to create zip", esl.Error(zipErr))
			z.OperationLog.Failure(zipErr, entry.Concrete())
			return
		}

		cpFile, crdErr := os.Open(dlCopy.Path())
		if crdErr != nil {
			ll.Debug("Unable to open downloaded file", esl.Error(crdErr))
			z.OperationLog.Failure(crdErr, entry.Concrete())
			return
		}
		defer func() {
			_ = cpFile.Close()
		}()

		if _, cpErr := io.Copy(zef, cpFile); cpErr != nil {
			ll.Debug("Unable to copy into the zip file", esl.Error(cpErr))
			z.OperationLog.Failure(cpErr, entry.Concrete())
			return
		}

		if flErr := zw.Flush(); flErr != nil {
			ll.Debug("Unable to flush", esl.Error(flErr))
			z.OperationLog.Failure(flErr, entry.Concrete())
			return
		}

		z.OperationLog.Success(dlFile, &RemoteZipEntry{
			Path: zipPath,
			Name: entry.Name(),
		})
	}, sv_file.Recursive(true))

	return lang.NewMultiErrorOrNull(err, dlErr, zipErr)
}

func (z *Remote) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("remote", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	return rc_exec.ExecMock(c, &Remote{}, func(r rc_recipe.Recipe) {
		m := r.(*Remote)
		m.LocalPath = mo_path2.NewFileSystemPath(filepath.Join(p, "test.zip"))
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("remote")
	})
}
