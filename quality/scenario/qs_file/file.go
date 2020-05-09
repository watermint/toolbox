package qs_file

import (
	"github.com/watermint/toolbox/essentials/log/esl"
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

func NewScenario(short bool) (sc Scenario, err error) {
	l := esl.Default()
	sc.LocalPath, err = ioutil.TempDir("", "file-upload-scenario")
	if err != nil {
		l.Error("unable to create temp dir", esl.Error(err))
		return
	}

	sc.Files = make(map[string]string)
	sc.Files["zyx/wvu.txt"] = "wvu"
	sc.Files["アイウ/エオ.txt"] = time.Now().String()

	if !short {
		sc.Files["a-b-c/time.txt"] = time.Now().String()
		sc.Files["123.txt"] = "123"
		sc.Files["abc.txt"] = "abc"
		sc.Files["あいう.txt"] = "あいう"
		sc.Files["time.txt"] = time.Now().String()
		sc.Files["987/654.txt"] = "654"
	}

	sc.Ignore = make(map[string]string)
	sc.Ignore[".DS_Store"] = "ignore-dsstore"
	if !short {
		sc.Ignore["987/~$abc"] = "ignore-abc"
		sc.Ignore["d-e-f/.~abc"] = "ignore-dot-tilde"
		sc.Ignore["~123.tmp"] = "ignore-123"
	}

	// Empty folders
	sc.Folders = make(map[string]bool)
	sc.Folders["987"] = true
	sc.Folders["zyx"] = true
	sc.Folders["アイウ"] = true
	if !short {
		sc.Folders["a-b-c"] = true
		sc.Folders["d-e-f"] = true
		sc.Folders["1-2-3"] = true
		sc.Folders["g-h-i/j-k-l"] = true
	}

	// Create test folders
	{
		for f := range sc.Folders {
			if err = os.MkdirAll(filepath.Join(sc.LocalPath, f), 0755); err != nil {
				l.Error("Unable to create folder", esl.Error(err), esl.String("f", f))
				return
			}
		}
	}

	// Create test files
	{
		for f, c := range sc.Files {
			if err = ioutil.WriteFile(filepath.Join(sc.LocalPath, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", esl.Error(err), esl.String("f", f))
				return
			}
		}
		for f, c := range sc.Ignore {
			if err := ioutil.WriteFile(filepath.Join(sc.LocalPath, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", esl.Error(err), esl.String("f", f))
				return
			}
		}
	}

	return
}
