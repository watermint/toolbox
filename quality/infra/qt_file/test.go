package qt_file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func MakeDummyFile(name string) (path string, err error) {
	d, err := ioutil.TempFile("", name)
	if err != nil {
		return "", err
	}
	_, err = d.Write([]byte(time.Now().Format(time.RFC3339)))
	if err != nil {
		os.Remove(d.Name())
		return "", err
	}
	d.Close()
	return d.Name(), nil
}

func MakeTestFile(name string, content string) (path string, err error) {
	d, err := ioutil.TempFile("", name)
	if err != nil {
		return "", err
	}
	_, err = d.Write([]byte(content))
	if err != nil {
		os.Remove(d.Name())
		return "", err
	}
	d.Close()
	return d.Name(), nil
}

func MakeTestFolder(name string, withContent bool) (path string, err error) {
	path, err = ioutil.TempDir("", name)
	if err != nil {
		return "", err
	}
	if withContent {
		err := ioutil.WriteFile(filepath.Join(path, "test.dat"), []byte(time.Now().String()), 0644)
		if err != nil {
			os.RemoveAll(path)
			return "", err
		}
	}
	return path, nil
}

func MustMakeTestFolder(ctl app_control.Control, name string, withContent bool) (path string) {
	path, err := MakeTestFolder(name, withContent)
	if err != nil {
		ctl.Log().Error("Unable to create test folder", zap.Error(err))
		ctl.Abort(app_control.Reason(app_control.FailureGeneral))
	}
	return path
}
