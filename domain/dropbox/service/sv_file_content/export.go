package sv_file_content

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"go.uber.org/zap"
	"os"
)

type Export interface {
	Export(path mo_path.DropboxPath) (export *mo_file.Export, localPath mo_path2.FileSystemPath, err error)
}

func NewExport(ctx dbx_context.Context) Export {
	return &exportImpl{ctx: ctx}
}

type exportImpl struct {
	ctx dbx_context.Context
}

func (z *exportImpl) Export(path mo_path.DropboxPath) (export *mo_file.Export, localPath mo_path2.FileSystemPath, err error) {
	l := z.ctx.Log()
	p := struct {
		Path string `json:"path"`
	}{
		Path: path.Path(),
	}

	res, err := z.ctx.Download("files/export").Param(p).Call()
	if err != nil {
		return nil, nil, err
	}
	contentFilePath, err := res.Body().AsFile()
	if err != nil {
		return nil, nil, err
	}
	resData := dbx_context.ContentResponseData(res)
	export = &mo_file.Export{}
	if _, err := resData.Model(export); err != nil {
		// Try remove downloaded file
		if removeErr := os.Remove(contentFilePath); removeErr != nil {
			l.Debug("Unable to remove exported file",
				zap.Error(err),
				zap.String("path", contentFilePath))
			// fall through
		}

		return nil, nil, err
	}
	return export, mo_path2.NewFileSystemPath(contentFilePath), nil
}
