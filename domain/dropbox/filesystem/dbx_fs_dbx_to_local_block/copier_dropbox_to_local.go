package dbx_fs_dbx_to_local_block

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_request"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/io/es_block"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
)

const (
	BlockSize = 4 * 1048576 // 4MiB
)

var (
	ErrorRangeRequestNotSupported = errors.New("range request not supported")
	ErrorInvalidContentLength     = errors.New("invalid content length")
)

type DownloadHeadResponse struct {
	Raw         json.RawMessage
	Name        string `json:"name" path:"name"`
	PathDisplay string `json:"path_display" path:"path_display"`
	Rev         string `json:"rev" path:"rev"`
	Size        int64  `json:"size" path:"size"`
}

func NewDropboxToLocal(ctx dbx_client.Client) es_filesystem.Connector {
	return &copierDropboxToLocal{
		ctx:    ctx,
		target: es_filesystem_local.NewFileSystem(),
	}
}

type copierDropboxToLocal struct {
	ctx       dbx_client.Client
	target    es_filesystem.FileSystem
	bwf       es_block.BlockWriterFactory
	indicator ea_indicator.Indicator
}

func (z *copierDropboxToLocal) Copy(source es_filesystem.Entry, target es_filesystem.Path, onSuccess func(pair es_filesystem.CopyPair, copied es_filesystem.Entry), onFailure func(pair es_filesystem.CopyPair, err es_filesystem.FileSystemError)) {
	l := z.ctx.Log().With(esl.Any("source", source.AsData()), esl.String("target", target.Path()))
	l.Debug("Copy (block download)")
	type P struct {
		Path string `json:"path"`
	}
	pair := es_filesystem.NewCopyPair(source, target)
	p := &P{Path: source.Path().Path()}
	q, err := dbx_request.DropboxApiArg(p)
	if err != nil {
		l.Debug("Unable to create path arg", esl.Error(err))
		onFailure(pair, dbx_fs.NewError(err))
		return
	}

	res := z.ctx.ContentHead("files/download", q)
	if err, fail := res.Failure(); fail {
		l.Debug("Head request failure", esl.Error(err))
		onFailure(pair, dbx_fs.NewError(err))
		return
	}

	if h := res.Header("Accept-Ranges"); h != "bytes" {
		l.Debug("The server does not support range request", esl.String("acceptRanges", h))
		onFailure(pair, dbx_fs.NewError(ErrorRangeRequestNotSupported))
		return
	}

	contentLength, err := strconv.ParseInt(res.Header("Content-Length"), 10, 64)
	if err != nil {
		l.Debug("invalid content length", esl.String("contentLength", res.Header("Content-Length")))
		onFailure(pair, dbx_fs.NewError(ErrorInvalidContentLength))
		return
	}

	// create zero byte file
	if contentLength == 0 {
		l.Debug("Create the zero byte file")
		f, err := os.Create(target.Path())
		if err != nil {
			l.Debug("Unable to create the file", esl.Error(err))
			onFailure(pair, dbx_fs.NewError(err))
		} else {
			_ = f.Close()
			if entry, fsErr := z.target.Info(target); fsErr != nil {
				onFailure(pair, fsErr)
			} else {
				onSuccess(pair, entry)
			}
		}
		return
	}

	resHeader := res.Header(dbx_client.DropboxApiResHeaderResult)
	j, err := es_json.ParseString(resHeader)
	if err != nil {
		l.Debug("Unable to parse response header", esl.Error(err), esl.String("header", resHeader))
		onFailure(pair, dbx_fs.NewError(err))
		return
	}

	apiResult := &DownloadHeadResponse{}
	if err := j.Model(apiResult); err != nil {
		l.Debug("Unable to parse response header", esl.Error(err), esl.String("header", resHeader))
		onFailure(pair, dbx_fs.NewError(err))
		return
	}

	revP := &P{Path: "rev:" + apiResult.Rev}
	revQ, err := dbx_request.DropboxApiArg(revP)
	if err != nil {
		l.Debug("Unable to create path arg", esl.Error(err))
		onFailure(pair, dbx_fs.NewError(err))
		return
	}

	z.indicator.AddTotal(contentLength)

	z.bwf.Open(
		target.Path(),
		contentLength,
		func(w es_block.BlockWriter, offset, blockSize int64) {
			requestRange := fmt.Sprintf("bytes=%d-%d", offset, min(offset+blockSize-1, contentLength))
			res = z.ctx.Download("files/download", revQ, api_request.Header("Range", requestRange))
			if err, fail := res.Failure(); fail {
				l.Debug("Error on download", esl.Error(err))
				w.Abort(offset, err)
				return
			}
			w.WriteBlock(res.Success().Body(), offset)
			l.Debug("A part downloaded", esl.String("Range", requestRange))
			z.indicator.AddProgress(blockSize)

		}, func(w es_block.BlockWriter, size int64) {
			if entry, fsErr := z.target.Info(target); fsErr != nil {
				onFailure(pair, fsErr)
			} else {
				onSuccess(pair, entry)
			}

		}, func(w es_block.BlockWriter, offset int64, err error) {
			onFailure(pair, dbx_fs.NewError(err))
		},
	)
}

func (z *copierDropboxToLocal) Startup(qd eq_queue.Definition) (err es_filesystem.FileSystemError) {
	z.bwf = es_block.NewWriterFactory(z.ctx.Log(), z.ctx.Feature().Concurrency(), BlockSize)
	z.indicator = ea_indicator.Global().NewIndicator(0,
		mpb.PrependDecorators(
			decor.Name("Download ", decor.WC{W: 20}),
			decor.AverageSpeed(decor.UnitKiB, "% 1.f"),
		),
		mpb.AppendDecorators(
			decor.CountersKibiByte(" % .1f / % .1f"),
			decor.OnComplete(
				decor.Spinner(mpb.DefaultSpinnerStyle, decor.WC{W: 5}), "DONE",
			),
		),
	)
	return nil
}

func (z *copierDropboxToLocal) Shutdown() (err es_filesystem.FileSystemError) {
	z.bwf.Wait()
	z.indicator.Done()
	return nil
}
