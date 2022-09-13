package stage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"golang.org/x/exp/rand"
	"time"
)

type UploadAppend struct {
	rc_recipe.RemarkSecret
	Peer dbx_conn.ConnScopedIndividual
	Path mo_path.DropboxPath
}

func (z *UploadAppend) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *UploadAppend) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("path", z.Path.Path()))
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

	var sessionId string
	{
		sessionRes := z.Peer.Client().Upload("files/upload_session/start",
			api_request.Content(es_rewinder.NewReadRewinderOnMemory([]byte{})),
			api_request.Param(&startParam))
		if err, f := sessionRes.Failure(); f {
			return err
		}
		sessionData := SessionData{}
		if err := sessionRes.Success().Json().Model(&sessionData); err != nil {
			return err
		}
		sessionId = sessionData.SessionId
		l.Info("Session started", esl.String("sessionId", sessionId))
	}
	l = l.With(esl.String("sessionId", sessionId))

	var blockSize int64 = 4096 * 1024
	dummyDataSource := rand.New(rand.NewSource(uint64(time.Now().Unix())))
	dummyDataBlock1 := make([]byte, blockSize)
	if _, err := dummyDataSource.Read(dummyDataBlock1); err != nil {
		return err
	}
	dummyDataBlock2 := make([]byte, blockSize/2)
	if _, err := dummyDataSource.Read(dummyDataBlock2); err != nil {
		return err
	}
	//dummyDataTotalSize := len(dummyDataBlock1) + len(dummyDataBlock2)

	type AppendCursor struct {
		SessionId string `json:"session_id"`
		Offset    int64  `json:"offset"`
	}
	type AppendParam struct {
		Cursor AppendCursor `json:"cursor"`
		Close  bool         `json:"close"`
	}

	// Append 2nd block
	{
		appendParam := AppendParam{
			Cursor: AppendCursor{
				SessionId: sessionId,
				Offset:    blockSize,
			},
			Close: true,
		}
		appendRes := z.Peer.Client().Upload("files/upload_session/append_v2",
			api_request.Param(&appendParam),
			api_request.Content(es_rewinder.NewReadRewinderOnMemory(dummyDataBlock2)))
		if err, f := appendRes.Failure(); f {
			return err
		}
	}

	// Append 1st block
	{
		appendParam := AppendParam{
			Cursor: AppendCursor{
				SessionId: sessionId,
				Offset:    0,
			},
			Close: false,
		}
		appendRes := z.Peer.Client().Upload("files/upload_session/append_v2",
			api_request.Param(&appendParam),
			api_request.Content(es_rewinder.NewReadRewinderOnMemory(dummyDataBlock1)))
		if err, f := appendRes.Failure(); f {
			return err
		}
	}

	// Finish
	type CommitInfo struct {
		Path           string `json:"path"`
		Mode           string `json:"mode"`
		Autorename     bool   `json:"autorename"`
		Mute           bool   `json:"mute"`
		StrictConflict bool   `json:"strict_conflict"`
		ClientModified string `json:"client_modified"`
	}

	type FinishParam struct {
		Cursor AppendCursor `json:"cursor"`
		Commit CommitInfo   `json:"commit"`
	}

	// finish
	{
		finishParam := FinishParam{
			Cursor: AppendCursor{
				SessionId: sessionId,
				//Offset:    int64(dummyDataTotalSize),
			},
			Commit: CommitInfo{
				Path:           z.Path.Path(),
				Mode:           "overwrite",
				Autorename:     false,
				Mute:           true,
				StrictConflict: false,
				ClientModified: dbx_util.ToApiTimeString(time.Now().Add(-10 * time.Hour)),
			},
		}
		finishRes := z.Peer.Client().Upload("files/upload_session/finish",
			api_request.Content(es_rewinder.NewReadRewinderOnMemory([]byte{})),
			api_request.Param(&finishParam))

		if err, f := finishRes.Failure(); f {
			return err
		}
		meta := &mo_file.ConcreteEntry{}
		if err := finishRes.Success().Json().Model(meta); err != nil {
			return err
		}
		l.Info("Success", esl.Any("uploaded", meta))
	}
	return nil
}

func (z *UploadAppend) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &UploadAppend{}, func(r rc_recipe.Recipe) {
		m := r.(*UploadAppend)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("upload-append")
	})
}
