package dfs_copier_batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/io/es_block"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"sync"
)

func NewLocalToDropboxBatch(ctl app_control.Control, ctx dbx_context.Context, batchSize int) es_filesystem.Connector {
	l := ctl.Log()
	if batchSize < 1 {
		l.Debug("Batch size less than one, fallback to one")
		batchSize = 1
	} else if 1000 < batchSize {
		l.Debug("Batch size greater than 1,000. fallback to 1,000")
		batchSize = 1000
	}

	return &copierLocalToDropboxBatch{
		batchSize:    batchSize,
		queue:        nil,
		sessions:     nil,
		block:        nil,
		fs:           es_block.NewPlainFileSystem(ctl.Log(), 4096*1024),
		ctl:          ctl,
		ctx:          ctx,
		backlogCount: sync.WaitGroup{},
	}
}

type copierLocalToDropboxBatch struct {
	batchSize    int
	queue        eq_queue.Container
	sessions     BatchSessions
	block        BlockSession
	fs           es_block.BlockFileSystem
	ctl          app_control.Control
	ctx          dbx_context.Context
	backlogCount sync.WaitGroup
}

type CopyBatchUploadBlock struct {
	SessionId string `json:"session_id"`
	Path      string `json:"path"`
	Offset    int64  `json:"offset"`
}

func (z *copierLocalToDropboxBatch) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	qd.Define(queueIdBlockUpload, z.uploadBlock)

	z.sessions, z.block = StartBatchSessions(qd, z.ctx, z.batchSize)
	z.queue = qd.Current()

	return nil
}

func (z *copierLocalToDropboxBatch) Shutdown() (err es_filesystem.FileSystemError) {
	l := z.ctl.Log()

	l.Debug("Waiting for shutdown")
	z.backlogCount.Wait()

	l.Debug("Shutdown sessions")
	if sErr := z.sessions.Shutdown(); sErr != nil {
		l.Debug("There was an error during shutdown", esl.Error(sErr))
		return filesystem.NewError(sErr)
	}
	return nil
}

func (z *copierLocalToDropboxBatch) uploadBlock(upload *CopyBatchUploadBlock) error {
	l := z.ctl.Log().With(esl.String("sessionId", upload.SessionId),
		esl.String("path", upload.Path), esl.Int64("offset", upload.Offset))

	data, isLastBlock, err := z.fs.ReadBlock(upload.Path, upload.Offset)
	if err != nil {
		l.Debug("Unable to receive the block", esl.Error(err))
		z.block.FinishFailure(upload.SessionId, err)
		return err
	}

	p := UploadAppend{
		Cursor: UploadCursor{
			SessionId: upload.SessionId,
			Offset:    upload.Offset,
		},
		Close: isLastBlock,
	}
	l.Debug("Upload block", esl.Bool("isLastBlock", isLastBlock))
	res := z.ctx.Upload("files/upload_session/append_v2",
		api_request.Param(&p),
		api_request.Content(es_rewinder.NewReadRewinderOnMemory(data)),
	)
	if err, f := res.Failure(); f {
		l.Debug("Error response from append", esl.Error(err))
		z.block.FinishFailure(upload.SessionId, err)
		return err
	}

	l.Debug("Block upload success")
	z.block.FinishSuccess(upload.SessionId, upload.Offset)
	return nil
}

func (z *copierLocalToDropboxBatch) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")
	cp := es_filesystem.NewCopyPair(source, target)

	localPath := source.Path().Path()
	targetDbxPath, dbxErr := filesystem.ToDropboxPath(target)
	if dbxErr != nil {
		l.Debug("unable to convert to Dropbox path", esl.Error(dbxErr))
		onFailure(cp, dbxErr)
		return
	}

	z.backlogCount.Add(1)

	sessionId, sesErr := z.sessions.NewSession(CommitInfo{
		Path:           targetDbxPath.Path(),
		Mode:           "overwrite",
		Autorename:     false,
		ClientModified: dbx_util.ToApiTimeString(source.ModTime()),
		Mute:           true,
		StrictConflict: false,
	}, SessionCallback{
		CopyPair: cp,
		OnSuccess: func(pair es_filesystem.CopyPair, copied es_filesystem.Entry) {
			onSuccess(pair, copied)
			z.backlogCount.Done()
		},
		OnFailure: func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError) {
			onFailure(pair, err)
			z.backlogCount.Done()
		},
	})

	l = l.With(esl.String("sessionId", sessionId))
	if sesErr != nil {
		l.Debug("Unable to start an upload session", esl.Error(sesErr))
		onFailure(cp, filesystem.NewError(sesErr))
		z.backlogCount.Done()
		return
	}

	l.Debug("Estimate blocks")
	offsets, fsErr := z.fs.FileBlocks(localPath)
	if fsErr != nil {
		l.Debug("Unable to retrieve file info", esl.Error(fsErr))
		onFailure(cp, filesystem.NewError(fsErr))
		z.backlogCount.Done()
		return
	}
	l.Debug("Estimated blocks", esl.Int("numBlocks", len(offsets)))

	for _, offset := range offsets {
		l.Debug("The block loaded", esl.Int64("offset", offset))

		z.block.AddBlock(sessionId, offset)
		q := z.queue.MustGet(queueIdBlockUpload)
		q.Enqueue(&CopyBatchUploadBlock{
			SessionId: sessionId,
			Path:      localPath,
			Offset:    offset,
		})
	}
}
