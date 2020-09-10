package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
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

	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("unable to marshal parameter", esl.Error(err))
		return nil, nil, err
	}

	res := z.ctx.Download("files/export", q)
	if err, f := res.Failure(); f {
		return nil, nil, err
	}
	contentFilePath, err := res.Success().AsFile()
	if err != nil {
		return nil, nil, err
	}
	resData := dbx_context.ContentResponseData(res)
	export = &mo_file.Export{}
	if err := resData.Model(export); err != nil {
		// Try remove downloaded file
		if removeErr := os.Remove(contentFilePath); removeErr != nil {
			l.Debug("Unable to remove exported file",
				esl.Error(err),
				esl.String("path", contentFilePath))
			// fall through
		}

		return nil, nil, err
	}
	return export, mo_path2.NewFileSystemPath(contentFilePath), nil
}
