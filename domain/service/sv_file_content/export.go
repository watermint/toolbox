package sv_file_content

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"go.uber.org/zap"
	"os"
)

type Export interface {
	Export(path mo_path.DropboxPath) (export *mo_file.Export, localPath mo_path.FileSystemPath, err error)
}

func NewExport(ctx api_context.Context) Export {
	return &exportImpl{ctx: ctx}
}

type exportImpl struct {
	ctx api_context.Context
}

func (z *exportImpl) Export(path mo_path.DropboxPath) (export *mo_file.Export, localPath mo_path.FileSystemPath, err error) {
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
	if !res.IsContentDownloaded() {
		return nil, nil, errors.New("content was not downloaded")
	}
	export = &mo_file.Export{}
	if err := res.Model(export); err != nil {
		// Try remove downloaded file
		if removeErr := os.Remove(res.ContentFilePath().Path()); removeErr != nil {
			l.Debug("Unable to remove exported file", zap.Error(err), zap.String("path", res.ContentFilePath().Path()))
			// fall through
		}

		return nil, nil, err
	}
	return export, res.ContentFilePath(), nil
}
