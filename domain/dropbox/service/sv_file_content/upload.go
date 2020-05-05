package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_request"
	"os"
	"path/filepath"
	"sync"
)

type Upload interface {
	Add(destPath mo_path.DropboxPath, filePath string) (entry mo_file.Entry, err error)
	Overwrite(destPath mo_path.DropboxPath, filePath string) (entry mo_file.Entry, err error)
	Update(destPath mo_path.DropboxPath, filePath string, revision string) (entry mo_file.Entry, err error)
}

type UploadOpt func(o *UploadOpts) *UploadOpts
type UploadOpts struct {
	ChunkSize int64
	Mute      bool
}

const (
	MaxChunkSize     = 150 * 1_048_576 // 150MB
	DefaultChunkSize = MaxChunkSize
)

var (
	warnExceededChunkSize = sync.Once{}
)

func NewUpload(ctx dbx_context.Context, opts ...UploadOpt) Upload {
	uo := &UploadOpts{
		ChunkSize: DefaultChunkSize,
		Mute:      false,
	}
	for _, o := range opts {
		o(uo)
	}
	if uo.ChunkSize < 1 {
		ctx.Log().Debug("Zero or negative chunk size. Fallback to max chunk size", es_log.Int64("givenChunkSize", uo.ChunkSize))
		uo.ChunkSize = MaxChunkSize
	}
	if uo.ChunkSize > MaxChunkSize {
		warnExceededChunkSize.Do(func() {
			ctx.Log().Warn("Chunk size exceed maximum size, chunk size will be adjusted to maximum size", es_log.Int64("givenChunkSize", uo.ChunkSize))
		})
		uo.ChunkSize = MaxChunkSize
	}
	return &uploadImpl{
		ctx: ctx,
		uo:  uo,
	}
}

func ChunkSizeKb(chunkSizeKb int) UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.ChunkSize = int64(chunkSizeKb * 1024)
		return o
	}
}

type uploadImpl struct {
	ctx dbx_context.Context
	uo  *UploadOpts
}

func (z *uploadImpl) Add(destPath mo_path.DropboxPath, filePath string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "add", "")
}

func (z *uploadImpl) Overwrite(destPath mo_path.DropboxPath, filePath string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "overwrite", "")
}

func (z *uploadImpl) Update(destPath mo_path.DropboxPath, filePath string, revision string) (entry mo_file.Entry, err error) {
	return z.upload(destPath, filePath, "update", revision)
}

func (z *uploadImpl) upload(destPath mo_path.DropboxPath, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
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

type UploadParamMode struct {
	Tag    string `json:".tag"`
	Update string `json:"update,omitempty"`
}

type UploadParams struct {
	Path           string           `json:"path"`
	Mode           *UploadParamMode `json:"mode"`
	Mute           bool             `json:"mute"`
	ClientModified string           `json:"client_modified"`
	Autorename     bool             `json:"autorename"`
}

func UploadPath(destPath mo_path.DropboxPath, f os.FileInfo) mo_path.DropboxPath {
	return destPath.ChildPath(filepath.Base(f.Name()))
}

func (z *uploadImpl) makeParams(info os.FileInfo, destPath mo_path.DropboxPath, mode string, revision string) *UploadParams {
	upm := &UploadParamMode{
		Tag:    mode,
		Update: "",
	}
	up := &UploadParams{
		Path:           UploadPath(destPath, info).Path(),
		Mode:           upm,
		Mute:           false,
		ClientModified: dbx_util.RebaseAsString(info.ModTime()),
	}
	switch mode {
	case "update":
		upm.Update = revision
	case "add":
		up.Autorename = true
	}
	return up
}

func (z *uploadImpl) uploadSingle(info os.FileInfo, destPath mo_path.DropboxPath, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	l := z.ctx.Log().With(es_log.String("filePath", filePath), es_log.Int64("size", info.Size()))
	l.Debug("Uploading file")

	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	rr, err := es_rewinder.NewReadRewinder(r, 0)
	if err != nil {
		l.Debug("Unable to create read rewinder", es_log.Error(err))
		return nil, err
	}
	defer r.Close()

	res := z.ctx.Upload("files/upload",
		api_request.Content(rr),
		api_request.Param(z.makeParams(info, destPath, mode, revision)))
	if err, f := res.Failure(); f {
		return nil, err
	}

	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return
}

func (z *uploadImpl) uploadChunked(info os.FileInfo, destPath mo_path.DropboxPath, filePath string, mode string, revision string) (entry mo_file.Entry, err error) {
	l := z.ctx.Log().With(es_log.String("filePath", filePath), es_log.Int64("size", info.Size()))

	total := info.Size()
	var written int64
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	type SessionId struct {
		SessionId string `path:"session_id" json:"session_id"`
	}
	type CursorInfo struct {
		SessionId string `json:"session_id"`
		Offset    int64  `json:"offset"`
	}
	type AppendInfo struct {
		Cursor *CursorInfo `json:"cursor"`
	}
	type CommitInfo struct {
		Cursor *CursorInfo   `json:"cursor"`
		Commit *UploadParams `json:"commit"`
	}

	l.Debug("Upload session start")
	r, err := es_rewinder.NewReadRewinderWithLimit(f, 0, z.uo.ChunkSize)
	if err != nil {
		l.Debug("Unable to create read rewinder", es_log.Error(err))
		return nil, err
	}
	res := z.ctx.Upload("files/upload_session/start",
		api_request.Content(r))
	if err, f := res.Failure(); f {
		return nil, err
	}
	sid := &SessionId{}
	if err := res.Success().Json().Model(sid); err != nil {
		return nil, err
	}
	written += z.uo.ChunkSize
	l = l.With(es_log.String("sessionId", sid.SessionId))

	for (total - written) > z.uo.ChunkSize {
		l.Debug("Append chunk", es_log.Int64("written", written))
		ai := &AppendInfo{
			Cursor: &CursorInfo{
				SessionId: sid.SessionId,
				Offset:    written,
			},
		}
		r, err := es_rewinder.NewReadRewinderWithLimit(f, written, z.uo.ChunkSize)
		if err != nil {
			l.Debug("Unable to create read rewinder", es_log.Error(err))
			return nil, err
		}
		res = z.ctx.Upload("files/upload_session/append_v2",
			api_request.Content(r),
			api_request.Param(ai))
		if err, fail := res.Failure(); fail {
			return nil, err
		}
		written += z.uo.ChunkSize
	}

	l.Debug("Finish")
	ci := &CommitInfo{
		Cursor: &CursorInfo{
			SessionId: sid.SessionId,
			Offset:    written,
		},
		Commit: z.makeParams(info, destPath, mode, revision),
	}
	r, err = es_rewinder.NewReadRewinderWithLimit(f, written, z.uo.ChunkSize)
	if err != nil {
		l.Debug("Unable to create read rewinder", es_log.Error(err))
		return nil, err
	}
	res = z.ctx.Upload("files/upload_session/finish",
		api_request.Content(r),
		api_request.Param(ci))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	entry = &mo_file.Metadata{}
	err = res.Success().Json().Model(entry)
	return entry, nil
}
