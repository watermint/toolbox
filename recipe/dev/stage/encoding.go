package stage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

type Encoding struct {
	rc_recipe.RemarkSecret
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Name     string
	Encoding string
}

func (z *Encoding) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *Encoding) selectEncoding(encName string) encoding.Encoding {
	switch strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(encName, "-", ""), "_", "")) {
	case "utf8":
		return unicode.UTF8
	case "utf8bom":
		return unicode.UTF8BOM
	case "utf16":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "utf16bom":
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	case "utf16le":
		return unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	case "utf16be":
		return unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	case "sjis", "shiftjis", "ms932", "cp932":
		return japanese.ShiftJIS
	case "iso2022jp":
		return japanese.ISO2022JP
	case "eucjp":
		return japanese.EUCJP
	case "euckr":
		return korean.EUCKR
	case "gb18030":
		return simplifiedchinese.GB18030
	case "gbk", "cp936":
		return simplifiedchinese.GBK
	case "hzgb2312":
		return simplifiedchinese.HZGB2312
	case "big5":
		return traditionalchinese.Big5
	case "iso88591": // ISO 8859-1
		return charmap.ISO8859_1
	default:
		// fallback to UTF8
		return unicode.UTF8
	}
}

func (z *Encoding) Exec(c app_control.Control) error {
	l := c.Log()
	ec := z.selectEncoding(z.Encoding)
	en, el, err := transform.String(ec.NewEncoder(), z.Name)
	if err != nil {
		return err
	}
	l.Info("Convert", esl.Int("origLength", len(z.Name)), esl.Int("convertedLength", el))
	up := z.Path.ChildPath(en)

	f, err := qt_file.MakeTestFile("encoding", z.Encoding)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	ulf, err := sv_file_content.NewUpload(z.Peer.Context(), sv_file_content.UseCustomFileName(true)).Add(up, f)
	if err != nil {
		return err
	}
	l.Info("Upload completed", esl.Any("meta", ulf))
	return nil
}

func (z *Encoding) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Encoding{}, func(r rc_recipe.Recipe) {
		m := r.(*Encoding)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("encoding")
		m.Encoding = "iso8859-1"
		m.Name = "エンコーディング"
	})
}
