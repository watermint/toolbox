package sv_paper

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_paper"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"os"
)

type Legacy interface {
	List(filterBy string, onEntry func(docId string)) error
	ListAccessed(onEntry func(docId string)) error
	ListCreated(onEntry func(docId string)) error
	Metadata(docId string, format string) (export *mo_paper.LegacyPaper, err error)
	Export(docId string, format string) (export *mo_paper.LegacyPaper, exportPath mo_path.FileSystemPath, err error)
}

func NewLegacy(ctx dbx_client.Client) Legacy {
	return &legacyImpl{
		ctx: ctx,
	}
}

type legacyRequest struct {
	FilterBy string `json:"filter_by"`
}

type legacyImpl struct {
	ctx dbx_client.Client
}

func (z legacyImpl) List(filterBy string, onEntry func(docId string)) error {
	l := z.ctx.Log().With(esl.String("filterBy", filterBy))
	q := &legacyRequest{
		FilterBy: filterBy,
	}
	res := z.ctx.List("paper/docs/list", api_request.Param(q)).Call(
		dbx_list.Continue("paper/docs/list/continue"),
		dbx_list.ResultTag("doc_ids"),
		dbx_list.UseHasMore(),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			if id, ok := entry.String(); !ok {
				l.Debug("Unexpected data format", esl.String("entry", entry.RawString()))
				return errors.New("unexpected response data format")
			} else {
				onEntry(id)
				return nil
			}
		}),
	)
	if err, fail := res.Failure(); fail {
		return err
	}
	return nil
}

func (z legacyImpl) ListAccessed(onEntry func(docId string)) error {
	return z.List("docs_accessed", onEntry)
}

func (z legacyImpl) ListCreated(onEntry func(docId string)) error {
	return z.List("docs_created", onEntry)
}

func (z legacyImpl) Metadata(docId string, format string) (export *mo_paper.LegacyPaper, err error) {
	l := z.ctx.Log().With(esl.String("docId", docId), esl.String("format", format))
	p := struct {
		DocId        string `json:"doc_id"`
		ExportFormat string `json:"export_format"`
	}{
		DocId:        docId,
		ExportFormat: format,
	}
	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("Unable to prepare arg", esl.Error(err))
		return nil, err
	}

	res := z.ctx.Download("paper/docs/download", q)
	if err, fail := res.Failure(); fail {
		l.Debug("Got an error", esl.Error(err))
		return nil, err
	}

	resData := dbx_client.ContentResponseData(res)
	export = &mo_paper.LegacyPaper{}
	if err := resData.Model(export); err != nil {
		return nil, err
	}

	return export, nil
}

func (z legacyImpl) Export(docId string, format string) (export *mo_paper.LegacyPaper, exportPath mo_path.FileSystemPath, err error) {
	l := z.ctx.Log().With(esl.String("docId", docId), esl.String("format", format))
	p := struct {
		DocId        string `json:"doc_id"`
		ExportFormat string `json:"export_format"`
	}{
		DocId:        docId,
		ExportFormat: format,
	}
	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("Unable to prepare arg", esl.Error(err))
		return nil, nil, err
	}

	res := z.ctx.DownloadRPC("paper/docs/download", q)
	if err, fail := res.Failure(); fail {
		l.Debug("Got an error", esl.Error(err))
		return nil, nil, err
	}

	contentFilePath, err := res.Success().AsFile()
	if err != nil {
		l.Debug("Unable to retrieve a file", esl.Error(err))
		return nil, nil, err
	}
	resData := dbx_client.ContentResponseData(res)
	export = &mo_paper.LegacyPaper{}
	if err := resData.Model(export); err != nil {
		if removeErr := os.Remove(contentFilePath); removeErr != nil {
			l.Debug("Unable to remove exported file", esl.Error(removeErr), esl.String("path", contentFilePath))
		}
		return nil, nil, err
	}

	return export, mo_path.NewFileSystemPath(contentFilePath), nil
}
