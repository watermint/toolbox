package file

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filehash"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Hash struct {
	Algorithm mo_string.SelectString
	File      mo_path.ExistingFileSystemPath
}

func (z *Hash) Preset() {
	z.Algorithm.SetOptions(
		"sha1",
		"md5",
		"sha1",
		"sha256",
	)
}

func (z *Hash) printDigest(c app_control.Control, digest string, err error) error {
	if err != nil {
		return err
	}
	ui_out.TextOut(c, digest)
	return nil
}

func (z *Hash) Exec(c app_control.Control) error {
	h := es_filehash.NewHash(c.Log())
	switch z.Algorithm.Value() {
	case "md5":
		d, err := h.MD5(z.File.Path())
		return z.printDigest(c, d, err)
	case "sha1":
		d, err := h.SHA1(z.File.Path())
		return z.printDigest(c, d, err)
	case "sha256":
		d, err := h.SHA256(z.File.Path())
		return z.printDigest(c, d, err)
	}
	return errors.New("unsupported algorithm")
}

func (z *Hash) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("hash", "hello")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.Exec(c, &Hash{}, func(r rc_recipe.Recipe) {
		m := r.(*Hash)
		m.File = mo_path.NewExistingFileSystemPath(f)
	})
}
