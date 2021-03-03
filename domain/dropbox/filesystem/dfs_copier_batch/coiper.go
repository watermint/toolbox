package dfs_copier_batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"sync"
	"time"
)

func NewLocalToDropboxBatch(ctl app_control.Control, ctx dbx_context.Context, batchSize int) es_filesystem.Connector {
	warehouse := NewBlockWarehouse(ctl.Log(), 4096*1024, ctl.Feature().Concurrency()*2)

	return &copierLocalToDropboxBatch{
		batchSize:    batchSize,
		queue:        nil,
		sessions:     nil,
		block:        nil,
		warehouse:    warehouse,
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
	warehouse    BlockWarehouse
	ctl          app_control.Control
	ctx          dbx_context.Context
	backlogCount sync.WaitGroup
}

type CopyBatchUploadBlock struct {
	SessionId   string `json:"session_id"`
	Path        string `json:"path"`
	Receipt     int64  `json:"receipt"`
	Offset      int64  `json:"offset"`
	IsLastBlock bool   `json:"is_last_block"`
}

func (z *copierLocalToDropboxBatch) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	qd.Define(queueIdBlockUpload, z.uploadBlock)
	qd.Define(queueIdBlockCheck, z.checkSession)

	z.sessions, z.block = StartBatchSessions(qd, z.ctx, z.batchSize)
	z.queue = qd.Current()
	z.warehouse.Startup()

	return nil
}

func (z *copierLocalToDropboxBatch) Shutdown() (err es_filesystem.FileSystemError) {
	l := z.ctl.Log()

	l.Debug("Waiting for shutdown")
	z.backlogCount.Wait()

	l.Debug("Shutdown warehouse:")
	z.warehouse.Shutdown()

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

	data, err := z.warehouse.Receive(BlockWarehouseReceipt(upload.Receipt))
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
		Close: upload.IsLastBlock,
	}
	l.Debug("Upload block")
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

func (z *copierLocalToDropboxBatch) checkSession(check *SessionCheck) error {
	l := z.ctx.Log().With(esl.Any("session", check))
	l.Debug("Check session")

	if z.warehouse.Status(check.Path) {
		l.Debug("The session is still in progress")
	} else {
		l.Debug("Check re-entry")
		time.Sleep(500 * time.Millisecond)
		z.queue.MustGet(queueIdBlockCheck).Enqueue(check)
	}
	return nil
}

func (z *copierLocalToDropboxBatch) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (upload)")
	cp := es_filesystem.NewCopyPair(source, target)

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

	// NOTE: This will not block:
	z.warehouse.Load(source.Path().Path(),
		func(bw BlockWarehouse, path string, offset int64, isLastBlock bool, receipt BlockWarehouseReceipt) {
			l.Debug("The block loaded", esl.String("path", path), esl.Int64("offset", offset),
				esl.Bool("isLastBlock", isLastBlock), esl.Int64("receipt", int64(receipt)))

			z.block.AddBlock(sessionId, offset, isLastBlock)

			q := z.queue.MustGet(queueIdBlockUpload)
			q.Enqueue(&CopyBatchUploadBlock{
				SessionId:   sessionId,
				Path:        path,
				Receipt:     int64(receipt),
				Offset:      offset,
				IsLastBlock: isLastBlock,
			})
		}, func(bw BlockWarehouse, path string, err error) {
			l.Debug("Error during load blocks", esl.Error(err))
			z.block.FinishFailure(sessionId, err)
		},
	)

	z.queue.MustGet(queueIdBlockCheck).Enqueue(&SessionCheck{
		SessionId: sessionId,
		Path:      source.Path().Path(),
	})
}
