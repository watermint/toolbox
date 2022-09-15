package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"os"
)

type Export interface {
	Export(path mo_path.DropboxPath, opts ...ExportOpt) (export *mo_file.Export, localPath mo_path2.FileSystemPath, err error)
}

type ExportOpts struct {
	Path         string `json:"path"`
	ExportFormat string `json:"export_format,omitempty"`
}

func (z ExportOpts) Apply(opts []ExportOpt) ExportOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type ExportOpt func(o ExportOpts) ExportOpts

func ExportFormat(format string) ExportOpt {
	return func(o ExportOpts) ExportOpts {
		o.ExportFormat = format
		return o
	}
}

func NewExport(ctx dbx_client.Client) Export {
	return &exportImpl{ctx: ctx}
}

type exportImpl struct {
	ctx dbx_client.Client
}

func (z *exportImpl) Export(path mo_path.DropboxPath, opts ...ExportOpt) (export *mo_file.Export, localPath mo_path2.FileSystemPath, err error) {
	l := z.ctx.Log()
	p := ExportOpts{
		Path: path.Path(),
	}.Apply(opts)

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
	resData := dbx_client.ContentResponseData(res)
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
