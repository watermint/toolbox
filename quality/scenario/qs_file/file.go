package qs_file

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Scenario struct {
	LocalPath string
	// path -> content
	Files map[string]string
	// path -> content
	Ignore map[string]string
	// folders
	Folders map[string]bool
}

func (z *Scenario) Create(short bool) (err error) {
	l := es_log.Default()
	z.LocalPath, err = ioutil.TempDir("", "file-upload-scenario")
	if err != nil {
		l.Error("unable to create temp dir", es_log.Error(err))
		return err
	}

	z.Files = make(map[string]string)
	z.Files["zyx/wvu.txt"] = "wvu"
	z.Files["アイウ/エオ.txt"] = time.Now().String()

	if !short {
		z.Files["a-b-c/time.txt"] = time.Now().String()
		z.Files["123.txt"] = "123"
		z.Files["abc.txt"] = "abc"
		z.Files["あいう.txt"] = "あいう"
		z.Files["time.txt"] = time.Now().String()
		z.Files["987/654.txt"] = "654"
	}

	z.Ignore = make(map[string]string)
	z.Ignore[".DS_Store"] = "ignore-dsstore"
	if !short {
		z.Ignore["987/~$abc"] = "ignore-abc"
		z.Ignore["d-e-f/.~abc"] = "ignore-dot-tilde"
		z.Ignore["~123.tmp"] = "ignore-123"
	}

	// Empty folders
	z.Folders = make(map[string]bool)
	z.Folders["987"] = true
	z.Folders["zyx"] = true
	z.Folders["アイウ"] = true
	if !short {
		z.Folders["a-b-c"] = true
		z.Folders["d-e-f"] = true
		z.Folders["1-2-3"] = true
		z.Folders["g-h-i/j-k-l"] = true
	}

	// Create test folders
	{
		for f := range z.Folders {
			if err := os.MkdirAll(filepath.Join(z.LocalPath, f), 0755); err != nil {
				l.Error("Unable to create folder", es_log.Error(err), es_log.String("f", f))
				return err
			}
		}
	}

	// Create test files
	{
		for f, c := range z.Files {
			if err := ioutil.WriteFile(filepath.Join(z.LocalPath, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", es_log.Error(err), es_log.String("f", f))
				return err
			}
		}
		for f, c := range z.Ignore {
			if err := ioutil.WriteFile(filepath.Join(z.LocalPath, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", es_log.Error(err), es_log.String("f", f))
				return err
			}
		}
	}

	return nil
}
