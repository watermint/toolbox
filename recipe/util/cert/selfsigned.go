package cert

import (
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/security/es_cert"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Selfsigned struct {
	Days mo_int.RangeInt
	Out  mo_path.FileSystemPath
}

func (z *Selfsigned) Preset() {
	z.Days.SetRange(1, 3650, 365)
}

func (z *Selfsigned) Exec(c app_control.Control) error {
	crt, key, err := es_cert.CreateSelfSigned(z.Days.Value())
	if err != nil {
		return err
	}
	keyFile, err := os.Create(filepath.Join(z.Out.Path(), "key.pem"))
	if err != nil {
		return err
	}
	defer func() {
		_ = keyFile.Close()
	}()
	if _, err = keyFile.Write(key); err != nil {
		return err
	}

	certFile, err := os.Create(filepath.Join(z.Out.Path(), "cert.pem"))
	if err != nil {
		return err
	}
	defer func() {
		_ = certFile.Close()
	}()
	if _, err = certFile.Write(crt); err != nil {
		return err
	}
	return nil
}

func (z *Selfsigned) Test(c app_control.Control) error {
	d, err := qt_file.MakeTestFolder("selfsigned", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()
	return rc_exec.ExecMock(c, &Selfsigned{}, func(r rc_recipe.Recipe) {
		m := r.(*Selfsigned)
		m.Out = mo_path.NewFileSystemPath(d)
	})
}
