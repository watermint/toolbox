package filesystem

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_model"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/em_file"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func NewModelToDropbox(modelRoot em_file.Folder, ctx dbx_context.Context, opts ...sv_file_content.UploadOpt) es_filesystem.Connector {
	return &copierModelToDropbox{
		ctx:        ctx,
		uploadOpts: opts,
		modelRoot:  modelRoot,
	}
}

type copierModelToDropbox struct {
	ctx        dbx_context.Context
	uploadOpts []sv_file_content.UploadOpt
	modelRoot  em_file.Folder
}

func (z copierModelToDropbox) Copy(source es_filesystem.Entry, target es_filesystem.Path) (copied es_filesystem.Entry, err es_filesystem.FileSystemError) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")

	sourceNode := em_file.ResolvePath(z.modelRoot, source.Path().Path())
	if sourceNode == nil {
		l.Debug("Unable to find the source node")
		return nil, es_filesystem_model.NewError(errors.New("source node not found"), es_filesystem_model.ErrorTypePathNotFound)
	}

	if sourceNode.Type() != em_file.FileNode {
		l.Debug("Node is not a file")
		return nil, es_filesystem_model.NewError(errors.New("source node is not a file"), es_filesystem_model.ErrorTypeOther)
	}

	targetDbxPath, err := ToDropboxPath(target)
	if err != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(err))
		return nil, err
	}
	content := sourceNode.(em_file.File).Content()

	tmpDir, ioErr := ioutil.TempDir("", "model_to_dropbox")
	if ioErr != nil {
		l.Debug("unable to create temp file", esl.Error(ioErr))
		return nil, NewError(ioErr)
	}

	tmpFilePath := filepath.Join(tmpDir, sourceNode.Name())
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	if errIO := ioutil.WriteFile(tmpFilePath, content, 0644); errIO != nil {
		l.Debug("Unable to write to the file", esl.Error(errIO))
		return nil, es_filesystem_local.NewError(errIO)
	}

	if errIO := os.Chtimes(tmpFilePath, time.Now(), source.ModTime()); errIO != nil {
		l.Debug("Unable to modify time", esl.Error(err))
	}

	svc := sv_file_content.NewUpload(z.ctx, z.uploadOpts...)
	dbxEntry, dbxErr := svc.Overwrite(targetDbxPath, tmpFilePath)
	if dbxErr != nil {
		l.Debug("Unable to upload file", esl.Error(dbxErr))
		return nil, NewError(dbxErr)
	}

	l.Debug("successfully uploaded", esl.Any("entry", dbxEntry.Concrete()))
	return NewEntry(dbxEntry), nil
}
