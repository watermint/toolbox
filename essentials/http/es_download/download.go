package es_download

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_bandwidth"
)

// ErrorNetworkDetail encapsulates the original error for network issues.
type ErrorNetworkDetail struct {
	Err error
}

func (e *ErrorNetworkDetail) Error() string {
	return "network error: unable to connect to the server (possible firewall, proxy, or network issue)"
}

func (e *ErrorNetworkDetail) Unwrap() error {
	return e.Err
}

var (
	ErrorNotFound = errors.New("not found")
)

// Download downloads file from `url` to `path`.
func Download(l esl.Logger, url string, path string) error {
	l.Debug("Try download", esl.String("url", url))
	resp, err := http.Get(url)
	if err != nil {
		l.Debug("Unable to create download request")
		return &ErrorNetworkDetail{Err: err}
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
		l.Debug("Unable to copy from response", esl.Error(err))
		return err
	}
	l.Debug("Finished download", esl.String("url", url))
	return nil
}

func DownloadText(l esl.Logger, url string) (string, error) {
	l.Debug("Try download", esl.String("url", url))
	resp, err := http.Get(url)
	if err != nil {
		l.Debug("Unable to create download request")
		return "", &ErrorNetworkDetail{Err: err}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		// OK, fall through
	case http.StatusNotFound:
		return "", ErrorNotFound
	default:
		l.Debug("Invalid status code", esl.Int("code", resp.StatusCode))
		return "", errors.New("status code " + strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	out := new(bytes.Buffer)

	_, err = io.Copy(out, nw_bandwidth.WrapReader(resp.Body))
	if err != nil {
		l.Debug("Unable to copy from response", esl.Error(err))
		return "", err
	}
	l.Debug("Finished download", esl.String("url", url))
	return out.String(), nil
}
