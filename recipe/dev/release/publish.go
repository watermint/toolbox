package release

import (
	"bytes"
	"errors"
	"fmt"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_filehash"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_runtime"
	"github.com/watermint/toolbox/recipe/dev/test"
	"go.uber.org/zap"
	"golang.org/x/text/language/display"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrorBuildIsNotReadyForRelease = errors.New("the build does not satisfy release criteria")
	ErrorNoArtifactFound           = errors.New("no artifact found")
)

type Publish struct {
	TestResource              string
	ArtifactPath              mo_path2.FileSystemPath
	HeadingReleaseTheme       app_msg.Message
	HeadingChanges            app_msg.Message
	ListSpecChange            app_msg.Message
	HeadingDocument           app_msg.Message
	ListReadme                app_msg.Message
	HeadingBinary             app_msg.Message
	BinaryTableHeaderFilename app_msg.Message
	BinaryTableHeaderSize     app_msg.Message
	BinaryTableHeaderMD5      app_msg.Message
	BinaryTableHeaderSHA256   app_msg.Message
}

type ArtifactSum struct {
	Filename string
	Size     int64
	MD5      string
	SHA256   string
}

func (z *Publish) Preset() {
	z.TestResource = defaultTestResource
}

func (z *Publish) verifyArtifacts(c app_control.Control) (a []*ArtifactSum, err error) {
	l := c.Log()
	a = make([]*ArtifactSum, 0)

	entries, err := ioutil.ReadDir(z.ArtifactPath.Path())
	if err != nil {
		return nil, err
	}

	h := ut_filehash.NewHash(l)
	for _, e := range entries {
		if !strings.HasPrefix(e.Name(), "tbx-"+app.Version) || !strings.HasSuffix(e.Name(), ".zip") {
			l.Debug("Ignore non artifact file", zap.Any("file", e))
			continue
		}
		p := filepath.Join(z.ArtifactPath.Path(), e.Name())
		sum := &ArtifactSum{
			Filename: e.Name(),
			Size:     e.Size(),
		}
		sum.MD5, err = h.MD5(p)
		if err != nil {
			l.Debug("Unable to calc MD5", zap.Error(err))
			return nil, err
		}
		sum.SHA256, err = h.SHA256(p)
		if err != nil {
			l.Debug("Unable to calc SHA256", zap.Error(err))
			return nil, err
		}
		a = append(a, sum)
	}
	if len(a) < 1 {
		return nil, ErrorNoArtifactFound
	}
	return a, nil
}

func (z *Publish) releaseNotes(c app_control.Control, sum []*ArtifactSum) error {
	l := c.Log()
	baseUrl := "https://github.com/watermint/toolbox/blob/" + app.Version

	var buf bytes.Buffer
	mui := app_ui.NewMarkdown(c.Messages(), &buf, true)
	mui.Header(z.HeadingReleaseTheme)
	mui.Break()

	mui.Header(z.HeadingChanges)

	for _, lang := range app_lang.SupportedLanguages {
		mui.Info(z.ListSpecChange.
			With("Link", baseUrl+"/doc/generated"+app_lang.PathSuffix(lang)+"/changes.md").
			With("Lang", display.Self.Name(lang)),
		)
	}

	mui.Break()
	mui.Header(z.HeadingDocument)

	for _, lang := range app_lang.SupportedLanguages {
		name := "README" + app_lang.PathSuffix(lang) + ".md"
		mui.Info(z.ListReadme.
			With("Name", name).
			With("Link", baseUrl+"/"+name).
			With("Lang", display.Self.Name(lang)),
		)
	}

	mui.Break()
	mui.Header(z.HeadingBinary)
	mit := mui.InfoTable("Binaries")
	mit.Header(z.BinaryTableHeaderFilename, z.BinaryTableHeaderSize, z.BinaryTableHeaderMD5, z.BinaryTableHeaderSHA256)

	for _, s := range sum {
		mit.RowRaw(s.Filename, fmt.Sprintf("%d", s.Size), s.MD5, s.SHA256)
	}
	mit.Flush()

	relNotesPath := filepath.Join(c.Workspace().Report(), "release_notes.md")
	err := ioutil.WriteFile(relNotesPath, buf.Bytes(), 0644)
	if err != nil {
		l.Debug("Unable to write release notes", zap.Error(err), zap.String("path", relNotesPath))
		return err
	}
	fmt.Println(buf.String())
	l.Info("Release note created", zap.String("path", relNotesPath))

	return nil
}

func (z *Publish) endToEndTest(c app_control.Control) error {
	l := c.Log()
	if c.Feature().IsTest() {
		l.Info("Skip tests")
		return nil
	}

	if c.Feature().IsProduction() {
		l.Info("Prepare resources")
		if !api_auth_impl.IsCacheAvailable(c, qt_endtoend.EndToEndPeer, []string{
			api_auth.DropboxTokenFull,
			api_auth.DropboxTokenBusinessAudit,
			api_auth.DropboxTokenBusinessManagement,
			api_auth.DropboxTokenBusinessFile,
			api_auth.DropboxTokenBusinessInfo,
		}) {
			return qt_errors.ErrorNotEnoughResource
		}
	}

	l.Info("Ensure end to end resource availability")
	if !dbx_conn_impl.IsEndToEndTokenAllAvailable(c) {
		l.Error("At least one of end to end resource is not available.")
		return errors.New("end to end resource is not available")
	}

	l.Info("Testing all end to end test")
	err := rc_exec.Exec(c, &test.Recipe{}, func(r rc_recipe.Recipe) {
		m := r.(*test.Recipe)
		m.All = true
		_, err := os.Lstat(z.TestResource)
		if err == nil {
			m.Resource = mo_string.NewOptional(z.TestResource)
		} else {
			l.Warn("Unable to read test resource", zap.String("path", z.TestResource), zap.Error(err))
		}
	})
	return err
}

func (z *Publish) Exec(c app_control.Control) error {
	l := c.Log()
	ready := true

	if app.IsProduction() {
		l.Info("Verify embedded resources")
		qt_runtime.Suite(c)
	} else {
		l.Info("Run as dev mode")
		ready = false
	}

	err := z.endToEndTest(c)
	if err != nil {
		return err
	}

	sum, err := z.verifyArtifacts(c)
	if err != nil {
		return err
	}

	err = z.releaseNotes(c, sum)
	if err != nil {
		return err
	}

	if !ready {
		l.Warn("The build does not satisfy release criteria")
		return ErrorBuildIsNotReadyForRelease
	}
	l.Info("The build is ready to publish")
	return nil
}

func (z *Publish) Test(c app_control.Control) error {
	d, err := qt_file.MakeTestFolder("release-publish", false)
	if err != nil {
		return err
	}

	platforms := []string{"linux", "mac", "win"}
	for _, platform := range platforms {
		app.Version = "dev-test"
		err = ioutil.WriteFile(filepath.Join(d, "tbx-"+app.Version+"-"+platform+".zip"), []byte("Test artifact"), 0644)
		if err != nil {
			c.Log().Warn("Unable to create test artifact", zap.Error(err))
			return err
		}
	}
	defer os.RemoveAll(d)

	err = rc_exec.Exec(c, &Publish{}, func(r rc_recipe.Recipe) {
		m := r.(*Publish)
		m.ArtifactPath = mo_path2.NewFileSystemPath(d)
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != ErrorBuildIsNotReadyForRelease && err != nil {
		return err
	}
	return nil
}
