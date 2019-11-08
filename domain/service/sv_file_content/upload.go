package sv_file_content

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"os"
	"path/filepath"
)

type Upload interface {
	Add(destPath mo_path.Path, filePath string) (entry mo_file.Entry, err error)
	Overwrite(destPath mo_path.Path, filePath string) (entry mo_file.Entry, err error)
	Update(destPath mo_path.Path, filePath string, revision string) (entry mo_file.Entry, err error)
}

type UploadOpt func(o *UploadOpts) *UploadOpts
type UploadOpts struct {
	ChunkSize int64
	Mute      bool
}

const (
	DefaultChunkSize = 150 * 1048576
)

func NewUpload(ctx api_context.Context, opts ...UploadOpt) Upload {
	uo := &UploadOpts{
		ChunkSize: DefaultChunkSize,
		Mute:      false,
	}
	for _, o := range opts {
		o(uo)
	}
	return &uploadImpl{
		ctx: ctx,
		uo:  uo,
	}
}

type uploadImpl struct {
	ctx api_context.Context
	uo  *UploadOpts
}

func (z *uploadImpl) Add(destPath mo_path.Path, filePath string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "add", "")
}

func (z *uploadImpl) Overwrite(destPath mo_path.Path, filePath string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "overwrite", "")
}

func (z *uploadImpl) Update(destPath mo_path.Path, filePath string, revision string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "update", revision)
}

func (z *uploadImpl) upload(destPath mo_path.Path, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	info, err := os.Lstat(filePath)
	if err != nil {
		return nil, err
	}
	if info.Size() < z.uo.ChunkSize {
		return z.uploadSingle(info, destPath, filePath, mode, revision)
	} else {
		return z.uploadChunked(info, destPath, filePath, mode, revision)
	}
}

type uploadParamMode struct {
	Tag    string `json:".tag"`
	Update string `json:"update,omitempty"`
}

type uploadParams struct {
	Path           string           `json:"path"`
	Mode           *uploadParamMode `json:"mode"`
	Mute           bool             `json:"mute"`
	ClientModified string           `json:"client_modified"`
}

func (z *uploadImpl) makeParams(info os.FileInfo, destPath mo_path.Path, mode string, revision string) *uploadParams {
	upm := &uploadParamMode{
		Tag:    mode,
		Update: "",
	}
	if mode == "update" && revision != "" {
		upm.Update = revision
	}

	return &uploadParams{
		Path:           destPath.ChildPath(filepath.Base(info.Name())).Path(),
		Mode:           upm,
		Mute:           false,
		ClientModified: api_util.RebaseAsString(info.ModTime()),
	}
}

func (z *uploadImpl) uploadSingle(info os.FileInfo, destPath mo_path.Path, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	res, err := z.ctx.Upload("files/upload").
		Param(z.makeParams(info, destPath, mode, revision)).Content(r).Call()
	if err != nil {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (z *uploadImpl) uploadChunked(info os.FileInfo, destPath mo_path.Path, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	panic("not yet")
}
