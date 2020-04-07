package dbx_response

import (
	"bufio"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_response"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"github.com/watermint/toolbox/infra/network/nw_monitor"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func New(ctx api_context.Context, req *http.Request, res *http.Response) (api_response.Response, error) {
	l := ctx.Log()
	defer nw_monitor.Log(req, res)

	if res == nil {
		l.Debug("Null response")
		return nil, api_response.ErrorNoResponse
	}

	result := ""
	if res.Header != nil {
		result = res.Header.Get(api_response.DropboxApiResHeaderResult)
	}
	if result != "" {
		resFile, err := ioutil.TempFile("", ctx.ClientHash())
		if err != nil {
			l.Debug("Unable to create temp file to store download", zap.Error(err))
			return nil, err
		}
		defer resFile.Close()
		buf := make([]byte, 4096)
		resBody := nw_bandwidth.WrapReader(res.Body)
		resWrite := bufio.NewWriter(resFile)
		var loadedLength int64

		for {
			n, err := resBody.Read(buf)
			if n > 0 {
				l.Debug("writing", zap.Int("writing", n))
				if _, wErr := resWrite.Write(buf[:n]); wErr != nil {
					l.Debug("Error on writing body to the file", zap.Error(wErr))
					return nil, err
				}
				loadedLength += int64(n)
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				l.Debug("Error on reading body", zap.Error(err))
				return nil, err
			}
			if n == 0 {
				// Wait for throttling
				time.Sleep(50 * time.Millisecond)
				continue
			}
		}
		resWrite.Flush()
		res.ContentLength = loadedLength

		return api_response.NewDownload(res, mo_path.NewFileSystemPath(resFile.Name()), result, loadedLength), nil

	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			l.Debug("Unable to read body", zap.Error(err))
			return nil, err
		}
		res.ContentLength = int64(len(body))

		return api_response.New(res, body), nil
	}
}
