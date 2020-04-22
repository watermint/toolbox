package download

import (
	"github.com/watermint/toolbox/infra/network/nw_bandwidth"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

func Download(l *zap.Logger, url string, path string) error {
	l.Debug("Try download", zap.String("url", url))
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
		l.Debug("Unable to copy from response", zap.Error(err))
		return err
	}
	l.Debug("Finished download", zap.String("url", url))
	return nil
}
