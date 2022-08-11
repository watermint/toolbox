package encoding

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/text/es_encoding"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
)

type From struct {
	Encoding    string
	In          mo_path.ExistingFileSystemPath
	Out         mo_path.FileSystemPath
	ErrNotFound app_msg.Message
}

func (z *From) Preset() {
}

func (z *From) Exec(c app_control.Control) error {
	enc := es_encoding.SelectEncoding(z.Encoding)
	if enc == nil {
		c.UI().Error(z.ErrNotFound.With("Encoding", z.Encoding))
		return errors.New("encoding not found")
	}
	in, err := os.Open(z.In.Path())
	if err != nil {
		return err
	}
	defer func() {
		_ = in.Close()
	}()
	out, err := os.Create(z.Out.Path())
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	tr := transform.NewReader(in, enc.NewDecoder())
	_, err = io.Copy(out, tr)
	return err
}

func (z *From) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("encoding", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()

	origContent, err := os.ReadFile("from_test_data.txt")
	if err != nil {
		return err
	}

	scenario := map[string]string{
		"sjis":   "from_test_data_sjis.txt",
		"euc-jp": "from_test_data_eucjp.txt",
	}

	for enc, dataFile := range scenario {
		of := filepath.Join(p, enc)
		err = rc_exec.Exec(c, &From{}, func(r rc_recipe.Recipe) {
			m := r.(*From)
			m.In = mo_path.NewExistingFileSystemPath(dataFile)
			m.Out = mo_path.NewFileSystemPath(of)
			m.Encoding = enc
		})
		if err != nil {
			return err
		}

		encContent, err := os.ReadFile(of)
		if err != nil {
			return err
		}
		if bc := bytes.Compare(origContent, encContent); bc != 0 {
			c.Log().Warn("Content mismatch", esl.Int("compareResult", bc), esl.String("encoding", enc))
			return errors.New("content mismatch")
		}
	}
	return nil
}
