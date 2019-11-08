package sv_file_content

import (
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"go.uber.org/zap"
	"io"
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
	Autorename     bool             `json:"autorename"`
}

func (z *uploadImpl) makeParams(info os.FileInfo, destPath mo_path.Path, mode string, revision string) *uploadParams {
	upm := &uploadParamMode{
		Tag:    mode,
		Update: "",
	}
	up := &uploadParams{
		Path:           destPath.ChildPath(filepath.Base(info.Name())).Path(),
		Mode:           upm,
		Mute:           false,
		ClientModified: api_util.RebaseAsString(info.ModTime()),
	}
	switch mode {
	case "update":
		upm.Update = revision
	case "add":
		up.Autorename = true
	}
	return up
}

func (z *uploadImpl) uploadSingle(info os.FileInfo, destPath mo_path.Path, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	l := z.ctx.Log().With(zap.String("filePath", filePath), zap.Int64("size", info.Size()))
	l.Debug("Uploading file")

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
	l := z.ctx.Log().With(zap.String("filePath", filePath), zap.Int64("size", info.Size()))

	total := info.Size()
	var written int64
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r := io.LimitReader(f, z.uo.ChunkSize)

	type sessionId struct {
		SessionId string `json:"session_id"`
	}
	type cursorInfo struct {
		SessionId string `json:"session_id"`
		Offset    int64  `json:"offset"`
	}
	type appendInfo struct {
		Cursor *cursorInfo `json:"cursor"`
	}
	type commitInfo struct {
		Cursor *cursorInfo   `json:"cursor"`
		Commit *uploadParams `json:"commit"`
	}

	l.Debug("Upload session start")
	res, err := z.ctx.Upload("files/files/upload_session/start").Content(r).Call()
	if err != nil {
		return nil, err
	}
	sid := &sessionId{}
	if err := res.Model(sid); err != nil {
		return nil, err
	}
	written += z.uo.ChunkSize
	l = l.With(zap.String("sessionId", sid.SessionId))

	for (total - written) > z.uo.ChunkSize {
		l.Debug("Append chunk", zap.Int64("written", written))
		ai := &appendInfo{
			Cursor: &cursorInfo{
				SessionId: sid.SessionId,
				Offset:    written,
			},
		}
		r = io.LimitReader(f, z.uo.ChunkSize)
		_, err := z.ctx.Upload("files/upload_session/append_v2").Param(ai).Content(r).Call()
		if err != nil {
			return nil, err
		}
		written += z.uo.ChunkSize
	}

	l.Debug("Finish")
	ci := &commitInfo{
		Cursor: &cursorInfo{
			SessionId: sid.SessionId,
			Offset:    written,
		},
		Commit: z.makeParams(info, destPath, mode, revision),
	}
	res, err = z.ctx.Upload("files/upload_session/finish").Param(ci).Content(f).Call()
	if err != nil {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	if err := res.Model(entry); err != nil {
		return nil, err
	}
	return entry, nil
}
