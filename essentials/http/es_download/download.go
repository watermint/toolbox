package es_download

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"io"
	"net/http"
	"os"
)

func Download(l es_log.Logger, url string, path string) error {
	l.Debug("Try download", es_log.String("url", url))
	resp, err := http.Get(url)
	if err != nil {
		l.Debug("Unable to create download request")
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		l.Debug("Unable to create download file")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, nw_bandwidth.WrapReader(resp.Body))
	if err != nil {
		l.Debug("Unable to copy from response", es_log.Error(err))
		return err
	}
	l.Debug("Finished download", es_log.String("url", url))
	return nil
}
