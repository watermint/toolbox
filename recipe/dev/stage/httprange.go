package stage

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"io"
	"os"
	"path/filepath"
)

type HttpRange struct {
	rc_recipe.RemarkSecret
	Peer        dbx_conn.ConnScopedIndividual
	DropboxPath mo_path.DropboxPath
	LocalPath   mo_path2.FileSystemPath
}

func (z *HttpRange) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
}

func (z *HttpRange) Exec(c app_control.Control) error {
	l := c.Log()

	type P struct {
		Path string `json:"path"`
	}
	p := &P{Path: z.DropboxPath.Path()}
	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("Unable to create an arg", esl.Error(err))
		return err
	}
	res := z.Peer.Client().ContentHead("files/download", q)
	if err, fail := res.Failure(); fail {
		return err
	}
	if h := res.Header("Accept-Ranges"); h != "bytes" {
		l.Debug("The server does not support range request", esl.String("acceptRanges", h))
		return errors.New("the server does not support range request")
	}
	contentLength := es_number.New(res.Header("Content-Length"))
	if !contentLength.IsInt() {
		l.Debug("invalid content length", esl.String("contentLength", res.Header("Content-Length")))
		return errors.New("invalid content length")
	}

	contentResponse := dbx_client.ContentResponseData(res)
	contentFile := &mo_file.File{}
	if err := contentResponse.Model(contentFile); err != nil {
		l.Debug("Unable to parse the result", esl.Error(err))
		return errors.New("unable to parse the result")
	}

	localPath := filepath.Join(z.LocalPath.Path(), contentFile.EntryName)
	l = l.With(esl.String("localPath", localPath))
	l.Debug("Create the file")

	localFile, err := os.Create(localPath)
	if err != nil {
		l.Debug("Unable to create the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = localFile.Close()
	}()

	var chunkSize int64
	chunkSize = 4 * 1048576

	type DownloadChunk struct {
		Offset    int64
		ChunkSize int64
	}

	downloader := func(chunk *DownloadChunk) error {
		requestRange := fmt.Sprintf("bytes=%d-%d", chunk.Offset, es_number.Min(chunk.Offset+chunk.ChunkSize-1, contentLength).Int64())
		res = z.Peer.Client().Download("files/download", q, api_request.Header("Range", requestRange))
		if err, fail := res.Failure(); fail {
			l.Debug("Error on download", esl.Error(err))
			return err
		}
		written, err := io.Copy(localFile, bytes.NewReader(res.Success().Body()))
		if err != nil {
			l.Debug("Could not write to local", esl.Error(err))
			return err
		}
		l.Info("A part downloaded", esl.String("Range", requestRange), esl.Int64("written", written))
		return nil
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("download", downloader)
		var loaded int64
		for loaded < contentLength.Int64() {
			s.Get("download").Enqueue(&DownloadChunk{
				Offset:    loaded,
				ChunkSize: chunkSize,
			})
			loaded += chunkSize
		}
	})

	return nil
}

func (z *HttpRange) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("range", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &HttpRange{}, func(r rc_recipe.Recipe) {
		m := r.(*HttpRange)
		m.LocalPath = mo_path2.NewFileSystemPath(f)
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("http_range.bin")
	})
}
