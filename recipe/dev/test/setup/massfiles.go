package setup

import (
	"compress/bzip2"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_async"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dfs_copier_batch"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/time/ut_format"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Contributor struct {
	Username string `xml:"username"`
	Id       string `xml:"id"`
}

type Revision struct {
	Id          string      `xml:"id"`
	ParentId    string      `xml:"parentid"`
	Timestamp   string      `xml:"timestamp"`
	Contributor Contributor `xml:"contributor"`
	Comment     string      `xml:"comment"`
	Model       string      `xml:"model"`
	Format      string      `xml:"format"`
	Text        string      `xml:"text"`
	Sha1        string      `xml:"sha1"`
}

type Page struct {
	Title    string      `xml:"title"`
	Ns       string      `xml:"ns"`
	Id       string      `xml:"id"`
	Revision []*Revision `xml:"revision"`
}

type WikimediaLoader struct {
	l    esl.Logger
	skip int
}

func (z WikimediaLoader) LoadBz2(path string, handler func(p Page) error) error {
	f, err := os.Open(path)
	if err != nil {
		z.l.Warn("Can't open the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	bf := bzip2.NewReader(f)

	return z.load(bf, handler)
}

func (z WikimediaLoader) LoadXml(path string, handler func(p Page) error) error {
	f, err := os.Open(path)
	if err != nil {
		z.l.Warn("Can't open the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return z.load(f, handler)
}

func (z WikimediaLoader) load(stream io.Reader, handler func(p Page) error) error {
	d := xml.NewDecoder(stream)
	index := 0
	for {
		t, err := d.Token()
		if err != nil {
			z.l.Warn("cannot parse", esl.Error(err))
			return err
		}
		if t == nil {
			z.l.Debug("Reached to EOL")
			return nil
		}

		switch se := t.(type) {
		case xml.StartElement:
			e := se.Name.Local
			switch e {
			case "page":
				index++
				if index <= z.skip {
					continue
				}

				var page Page
				if err := d.DecodeElement(&page, &se); err != nil {
					z.l.Warn("Can't error", esl.Error(err))
					return err
				}
				if err := handler(page); err != nil {
					z.l.Warn("Can't handle the page", esl.Error(err), esl.Any("page", page))
					return err
				}
			}
		}
	}
}

type Massfiles struct {
	rc_recipe.RemarkSecret
	Peer      dbx_conn.ConnScopedIndividual
	Source    mo_path.ExistingFileSystemPath
	Base      mo_path2.DropboxPath
	Offset    int
	BatchSize mo_int.RangeInt
}

func (z *Massfiles) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
	z.BatchSize.SetRange(0, 1000, 1000)
}

func (z *Massfiles) Exec(c app_control.Control) error {
	l := c.Log()

	sessions := make(map[string]Page)
	offsets := make(map[string]int64)
	sessionMutex := sync.Mutex{}
	batchSize := z.BatchSize.Value()
	ctx := z.Peer.Context()

	pageContent := func(p Page) string {
		switch len(p.Revision) {
		case 0:
			return p.Title
		default:
			return fmt.Sprintf("= %s =\n\n----\n%s", p.Title, p.Revision[0].Text)
		}
	}

	pageTime := func(p Page) string {
		if len(p.Revision) < 1 {
			return dbx_util.ToApiTimeString(time.Now())
		} else {
			pt, valid := ut_format.ParseTimestamp(p.Revision[0].Timestamp)
			if !valid {
				return dbx_util.ToApiTimeString(time.Now())
			} else {
				return dbx_util.ToApiTimeString(pt)
			}
		}
	}

	pageToPath := func(p Page) []string {
		altPageId := "p-" + p.Id + ".txt"
		pageId, err := strconv.ParseInt(p.Id, 10, 32)
		if err != nil {
			l.Debug("Unable to parse pageId", esl.Error(err), esl.String("pageId", p.Id))
			return []string{"unexpected_page_id", altPageId}
		}
		altPageId = fmt.Sprintf("%d/%d/p-%d.txt", pageId/1_000_000, pageId/1000, pageId)
		if len(p.Revision) < 1 {
			return []string{"no_revision", altPageId}
		}
		pt, valid := ut_format.ParseTimestamp(p.Revision[0].Timestamp)
		if valid {
			return []string{
				fmt.Sprintf("%04d", pt.Year()),
				fmt.Sprintf("%04d-%02d", pt.Year(), pt.Month()),
				fmt.Sprintf("%04d-%02d-%02d", pt.Year(), pt.Month(), pt.Day()),
				fmt.Sprintf("%s.txt", p.Id),
			}
		}
		return []string{"invalid_time_format", altPageId}
	}

	commit := func() error {
		l.Info("Commit", esl.Int("size", len(sessions)))
		if len(sessions) < 1 {
			return nil
		}

		commits := make([]dfs_copier_batch.UploadFinish, 0)
		paths := make([]string, 0)
		for sessionId, page := range sessions {
			path := z.Base.ChildPath(pageToPath(page)...).Path()
			offset := offsets[sessionId]
			paths = append(paths, path)
			commits = append(commits, dfs_copier_batch.UploadFinish{
				Cursor: dfs_copier_batch.UploadCursor{
					SessionId: sessionId,
					Offset:    offset,
				},
				Commit: dfs_copier_batch.CommitInfo{
					Path:           path,
					Mode:           "add",
					Autorename:     true,
					ClientModified: pageTime(page),
					Mute:           true,
					StrictConflict: false,
				},
			})
		}

		finish := &dfs_copier_batch.UploadFinishBatch{
			Entries: commits,
		}
		res := ctx.Async("files/upload_session/finish_batch", api_request.Param(finish)).Call(
			dbx_async.Status("files/upload_session/finish_batch/check"),
		)

		if err, f := res.Failure(); f {
			l.Debug("Unable to finish the batch", esl.Error(err))
			return err
		}

		// clean sessions
		sessions = make(map[string]Page)
		offsets = make(map[string]int64)
		l.Info("Commit batch", esl.Strings("paths", paths))
		return nil
	}

	h := func(p Page) error {
		sessionMutex.Lock()
		if batchSize <= len(sessions) {
			if err := commit(); err != nil {
				return err
			}
		}
		sessionMutex.Unlock()

		type StartSessionParam struct {
			Close       bool   `json:"close"`
			SessionType string `json:"session_type,omitempty"`
		}
		type SessionData struct {
			SessionId string `path:"session_id" json:"session_id"`
		}

		content := pageContent(p)

		ssp := &StartSessionParam{
			Close:       true,
			SessionType: "sequential",
		}
		sessionRes := ctx.Upload("files/upload_session/start",
			api_request.Content(es_rewinder.NewReadRewinderOnMemory([]byte(content))),
			api_request.Param(&ssp))
		if err, f := sessionRes.Failure(); f {
			l.Debug("Unable to start the session", esl.Error(err))
			return err
		}
		sessionData := SessionData{}
		if err := sessionRes.Success().Json().Model(&sessionData); err != nil {
			l.Debug("Unable to parse session data", esl.Error(err))
			return err
		}
		if sessionData.SessionId == "" {
			l.Debug("Unable to retrieve session id")
			return errors.New("no session id found")
		}

		sessionMutex.Lock()
		sessions[sessionData.SessionId] = p
		offsets[sessionData.SessionId] = int64(len([]byte(content)))
		sessionMutex.Unlock()

		return nil
	}

	var wl = &WikimediaLoader{
		l:    c.Log(),
		skip: z.Offset,
	}
	sourcePath := z.Source.Path()
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("upload", h)
		q := s.Get("upload")

		switch {
		case strings.HasSuffix(sourcePath, ".xml.bz2"):
			wl.LoadBz2(sourcePath, func(p Page) error {
				q.Enqueue(p)
				return nil
			})
		case strings.HasSuffix(sourcePath, ".xml"):
			wl.LoadXml(sourcePath, func(p Page) error {
				q.Enqueue(p)
				return nil
			})
		default:
			panic("Look like the file is not supported format:" + sourcePath)
		}
	})
	return commit()
}

func (z *Massfiles) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
