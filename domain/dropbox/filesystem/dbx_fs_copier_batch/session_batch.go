package dbx_fs_copier_batch

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"sync"
	"time"
)

type SessionCallback struct {
	CopyPair  es_filesystem.CopyPair
	OnSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry)
	OnFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)
}

type BatchSessions interface {
	// NewSession Start a new session with the commit information.
	// The operation might blocked depends on the backlog size.
	NewSession(commit CommitInfo, callback SessionCallback) (sessionId string, err error)

	// AbortSession Abort uploading session for the session Id.
	// The sessionId will be removed from the batch.
	// The func callbacks OnFailure event.
	AbortSession(sessionId string, err error)

	// FinishBlockUploads Report finish for all block uploads.
	// The func should not block the operation.
	FinishBlockUploads(sessionId string)

	// Shutdown Flush and finish all sessions.
	// The operation might blocked when the backlog exists.
	Shutdown() (err dbx_error.DropboxError)

	// FinishBatchCommit Process upload session commit.
	FinishBatchCommit(batch *FinishBatch)

	// FinishBatchEntry Entry upload session commit batch.
	FinishBatchEntry(count int)
}

func StartBatchSessions(qd eq_queue.Definition, ctx dbx_client.Client, batchSize int) (BatchSessions, BlockSession) {
	batchSessions := &copierBatchSessions{
		ctx:                    ctx,
		batchSize:              batchSize,
		sessionIdToMutex:       sync.Mutex{},
		sessionIdToCommit:      make(map[string]CommitInfo),
		sessionIdToCallback:    make(map[string]SessionCallback),
		finishedSessionIdMutex: sync.Mutex{},
		finishedSessionIds:     make(map[string]bool, 0),
	}

	qd.Define(queueIdBlockBatch, batchSessions.FinishBatchEntry)
	qd.Define(queueIdBlockCommit, batchSessions.FinishBatchCommit)

	batchSessions.queueBatchEntry = qd.Current().MustGet(queueIdBlockBatch)
	batchSessions.queueBatchCommit = qd.Current().MustGet(queueIdBlockCommit)

	blockSession := NewBlockSession(ctx, batchSessions)

	return batchSessions, blockSession
}

type copierBatchSessions struct {
	queueBatchCommit       eq_queue.Queue
	queueBatchEntry        eq_queue.Queue
	ctx                    dbx_client.Client
	batchSize              int
	blockSession           BlockSession
	sessionIdToMutex       sync.Mutex
	sessionIdToCommit      map[string]CommitInfo
	sessionIdToCallback    map[string]SessionCallback
	finishedSessionIdMutex sync.Mutex
	finishedSessionIds     map[string]bool
}

func (z *copierBatchSessions) IsSessionAlive(sessionId string) (found bool) {
	z.sessionIdToMutex.Lock()
	_, found = z.sessionIdToCommit[sessionId]
	z.sessionIdToMutex.Unlock()
	return found
}

// FinishBatchCommit Transaction finish batch: Exec from queue:
func (z *copierBatchSessions) FinishBatchCommit(batch *FinishBatch) {
	l := z.ctx.Log()
	l.Debug("Start batch", esl.Strings("batchSessionIds", batch.Batch))

	p := UploadFinishBatch{
		Entries: make([]UploadFinish, 0),
	}
	deadSessionId := make([]string, 0)
	z.sessionIdToMutex.Lock()
	for _, sessionId := range batch.Batch {
		if ci, ok := z.sessionIdToCommit[sessionId]; ok {
			p.Entries = append(p.Entries, UploadFinish{
				Cursor: UploadCursor{
					SessionId: sessionId,
					Offset:    0,
				},
				Commit: ci,
			})
		} else {
			l.Debug("Unable to find the sessionId", esl.String("sessionId", sessionId))
			deadSessionId = append(deadSessionId, sessionId)
		}
	}
	z.sessionIdToMutex.Unlock()

	endSessionNoLock := func(sessionId string) {
		l.Debug("End session", esl.String("sessionId", sessionId))
		delete(z.sessionIdToCommit, sessionId)
		delete(z.sessionIdToCallback, sessionId)
	}

	defer func() {
		if len(deadSessionId) > 0 {
			// try find callbacks, panic if not found
			z.sessionIdToMutex.Lock()
			for _, sid := range deadSessionId {
				if cb, ok := z.sessionIdToCallback[sid]; ok {
					cb.OnFailure(cb.CopyPair, dbx_fs.NewError(errors.New("unable to find upload session id")))
				} else {
					l.Error("Unable to find upload session id", esl.String("sessionId", sid))
				}
				endSessionNoLock(sid)
			}
			z.sessionIdToMutex.Unlock()
		}
	}()

	res := z.ctx.Async("files/upload_session/finish_batch", api_request.Param(&p)).Call(
		dbx_async.Status("files/upload_session/finish_batch/check"),
	)

	broadcastError := func(err error) {
		z.sessionIdToMutex.Lock()
		for _, e := range p.Entries {
			if cb, ok := z.sessionIdToCallback[e.Cursor.SessionId]; ok {
				cb.OnFailure(cb.CopyPair, dbx_fs.NewError(err))
			} else {
				l.Debug("Callback not found", esl.String("sessionId", e.Cursor.SessionId))
			}
			endSessionNoLock(e.Cursor.SessionId)
		}
		z.sessionIdToMutex.Unlock()
	}

	if err, f := res.Failure(); f {
		l.Debug("Unable to finish the batch", esl.Error(err))
		broadcastError(err)
		return
	}

	resEntries, found := res.Success().Json().FindArray("entries")
	if !found {
		l.Debug("`entries` not found in the response")
		broadcastError(errors.New("no entries data found"))
		return
	}

	if len(resEntries) != len(p.Entries) {
		l.Debug("Inconsistent number of entries between the param and the response",
			esl.Int("request", len(p.Entries)),
			esl.Int("response", len(resEntries)),
		)
		broadcastError(errors.New("inconsistent entries count"))
		return
	}

	z.sessionIdToMutex.Lock()
	for i, reqEntry := range p.Entries {
		sid := reqEntry.Cursor.SessionId
		ll := l.With(esl.String("sessionId", sid))
		resEntryJson := resEntries[i]
		resEntryTag, found := resEntryJson.FindString("\\.tag")
		if !found || resEntryTag != "success" {
			reason := resEntryJson.RawString() // TODO: change it more better way to extract message
			if cb, ok := z.sessionIdToCallback[sid]; ok {
				ll.Debug("Error on finish batch")
				cb.OnFailure(cb.CopyPair, dbx_fs.NewError(errors.New(reason)))
			} else {
				ll.Debug("Callback not found")
			}
		} else {
			entryMeta := &mo_file.Metadata{}
			if err := resEntryJson.Model(entryMeta); err != nil {
				ll.Debug("Unable to unmarshal metadata", esl.Error(err))
				if cb, ok := z.sessionIdToCallback[sid]; ok {
					ll.Debug("Error on finish batch")
					cb.OnFailure(cb.CopyPair, dbx_fs.NewError(err))
				} else {
					ll.Debug("Callback not found")
				}
			} else {
				if cb, ok := z.sessionIdToCallback[sid]; ok {
					ll.Debug("Success entry", esl.Any("entry", entryMeta))
					cb.OnSuccess(cb.CopyPair, dbx_fs.NewEntry(entryMeta))
				} else {
					ll.Debug("Callback not found")
				}
			}
		}

		// end session
		endSessionNoLock(sid)
	}
	z.sessionIdToMutex.Unlock()
}

func (z *copierBatchSessions) AbortSession(sessionId string, err error) {
	l := z.ctx.Log().With(esl.String("sessionId", sessionId))
	l.Debug("Abort session", esl.Error(err))

	z.sessionIdToMutex.Lock()
	delete(z.sessionIdToCallback, sessionId)
	delete(z.sessionIdToCommit, sessionId)
	z.sessionIdToMutex.Unlock()

	z.finishedSessionIdMutex.Lock()
	delete(z.finishedSessionIds, sessionId)
	z.finishedSessionIdMutex.Unlock()
}

func (z *copierBatchSessions) FinishBatchEntry(count int) {
	l := z.ctx.Log()

	var finishBatch []string = make([]string, 0)

	z.sessionIdToMutex.Lock()
	z.finishedSessionIdMutex.Lock()
	if len(z.finishedSessionIds) >= z.batchSize {
		l.Debug("finished sessions exceeds batch size")
		for sid := range z.finishedSessionIds {
			finishBatch = append(finishBatch, sid)
			delete(z.finishedSessionIds, sid)
			if len(finishBatch) >= z.batchSize {
				break
			}
		}
	} else if len(z.finishedSessionIds) == len(z.sessionIdToCommit) {
		l.Debug("Look like all sessions completed")
		for sid := range z.finishedSessionIds {
			finishBatch = append(finishBatch, sid)
			delete(z.finishedSessionIds, sid)
			if len(finishBatch) >= z.batchSize {
				break
			}
		}
	}
	l.Debug("Finished sizes",
		esl.Int("finished", len(z.finishedSessionIds)),
		esl.Int("sessions", len(z.sessionIdToCommit)),
		esl.Int("commits", len(finishBatch)))

	z.sessionIdToMutex.Unlock()
	z.finishedSessionIdMutex.Unlock()

	if x := len(finishBatch); x > 0 {
		l.Debug("Enqueue finish batch", esl.Int("batchSize", x))
		//z.queueBatchCommit.Enqueue(FinishBatch{
		//	Batch: finishBatch,
		//})
		z.FinishBatchCommit(&FinishBatch{
			Batch: finishBatch,
		})
	} else {
		l.Debug("No session found for the batch")
	}
}

func (z *copierBatchSessions) FinishBlockUploads(sessionId string) {
	l := z.ctx.Log().With(esl.String("sessionId", sessionId))
	l.Debug("finish session")

	z.finishedSessionIdMutex.Lock()
	z.finishedSessionIds[sessionId] = true
	z.finishedSessionIdMutex.Unlock()

	//z.queueBatchEntry.Enqueue(len(z.finishedSessionIds))
	z.FinishBatchEntry(len(z.finishedSessionIds))
}

func (z *copierBatchSessions) Shutdown() (err dbx_error.DropboxError) {
	for {
		if len(z.sessionIdToCommit) > 0 {
			time.Sleep(1 * time.Second)
		}
		return nil
	}
}

func (z *copierBatchSessions) NewSession(commit CommitInfo, callback SessionCallback) (sessionId string, err error) {
	l := z.ctx.Log().With(esl.String("path", commit.Path))
	l.Debug("New session", esl.Any("commitInfo", commit))
	type StartSessionParam struct {
		Close       bool   `json:"close"`
		SessionType string `json:"session_type,omitempty"`
	}
	type SessionData struct {
		SessionId string `path:"session_id" json:"session_id"`
	}
	startParam := StartSessionParam{
		Close:       false,
		SessionType: "concurrent",
	}
	l.Debug("Start session")
	sessionRes := z.ctx.Upload("files/upload_session/start",
		api_request.Content(es_rewinder.NewReadRewinderOnMemory([]byte{})),
		api_request.Param(&startParam))
	if err, f := sessionRes.Failure(); f {
		l.Debug("Unable to start the session", esl.Error(err))
		return "", dbx_error.NewErrors(err)
	}
	sessionData := SessionData{}
	if err := sessionRes.Success().Json().Model(&sessionData); err != nil {
		l.Debug("Unable to parse session data", esl.Error(err))
		return "", err
	}
	if sessionData.SessionId == "" {
		l.Debug("Unable to retrieve session id")
		return "", errors.New("no session id found")
	}

	// update session id to commit info
	z.sessionIdToMutex.Lock()
	z.sessionIdToCommit[sessionData.SessionId] = commit
	z.sessionIdToCallback[sessionData.SessionId] = callback
	z.sessionIdToMutex.Unlock()

	return sessionData.SessionId, nil
}
