package apply

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_tag"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_template"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Remote struct {
	Peer                 dbx_conn.ConnScopedIndividual
	Template             mo_path.ExistingFileSystemPath
	Path                 mo_path2.DropboxPath
	ProgressAddTags      app_msg.Message
	ProgressPutFile      app_msg.Message
	ProgressCreateFolder app_msg.Message
}

func (z *Remote) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeFilesMetadataWrite,
	)
}

func (z *Remote) addTags(path es_filesystem.Path, tags []string) error {
	svt := sv_file_tag.New(z.Peer.Client())
	for _, tag := range tags {
		if err := svt.Add(mo_path2.NewDropboxPath(path.Path()), tag); err != nil {
			return err
		}
	}
	return nil
}

func (z *Remote) putFile(path es_filesystem.Path, f io.ReadSeeker) error {
	l := esl.Default().With(esl.String("path", path.Path()))
	rr, err := es_rewinder.NewReadRewinder(f, 0)
	if err != nil {
		l.Debug("Unable to create read rewinder", esl.Error(err))
		return err
	}
	svu := sv_file_content.NewUploadStream(z.Peer.Client(), false, true)
	entry, err := svu.Upload(mo_path2.NewDropboxPath(path.Path()), rr, time.Now())
	if err != nil {
		l.Debug("Unable to upload", esl.Error(err))
		return err
	}
	l.Debug("Uploaded", esl.Any("entry", entry.Concrete()))
	return nil
}

func (z *Remote) Exec(c app_control.Control) error {
	tmpl, err := os.ReadFile(z.Template.Path())
	if err != nil {
		return err
	}
	tmplData := es_template.Root{}
	if err := json.Unmarshal(tmpl, &tmplData); err != nil {
		return err
	}

	fs := dbx_fs.NewFileSystem(z.Peer.Client())
	ap := es_template.NewApply(fs,
		es_template.ApplyOpts{
			HandlerTagAdd: func(path es_filesystem.Path, tags []string) error {
				c.UI().Progress(z.ProgressAddTags.With("Path", path.Path()).With("Tags", strings.Join(tags, ", ")))
				return z.addTags(path, tags)
			},
			HandlerPutFile: func(path es_filesystem.Path, f io.ReadSeeker) error {
				c.UI().Progress(z.ProgressPutFile.With("Path", path.Path()))
				return z.putFile(path, f)
			},
			OnCreateFolder: func(path es_filesystem.Path) {
				c.UI().Progress(z.ProgressCreateFolder.With("Path", path.Path()))
			},
		},
	)
	return ap.Apply(dbx_fs.NewPath("", z.Path), tmplData)
}

func (z *Remote) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("tmpl", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	tmplPath := filepath.Join(f, "template.json")
	err = os.WriteFile(tmplPath, []byte(`{}`), 0644)
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Remote{}, func(r rc_recipe.Recipe) {
		m := r.(*Remote)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("apply")
		m.Template = mo_path.NewExistingFileSystemPath(tmplPath)
	})
}
