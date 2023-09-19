package build

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_license"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/recipe/dev/ci/artifact"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Package struct {
	rc_recipe.RemarkSecret
	BuildPath  mo_path.ExistingFileSystemPath
	DistPath   mo_path.FileSystemPath
	DeployPath mo_string.OptionalString
	Up         *artifact.Up
}

func (z *Package) Preset() {
}

func (z *Package) platformName() string {
	target, ok := os.LookupEnv(app.EnvNameToolboxBuildTarget)
	if ok {
		switch target {
		case "windows/amd64":
			return "win"
		case "linux/amd64":
			return "linux-intel"
		case "linux/arm64":
			return "linux-arm"
		case "darwin/amd64":
			return "mac-intel"
		case "darwin/arm64":
			return "mac-applesilicon"
		}
	}
	return "unknown"
}

func (z *Package) createPackage(c app_control.Control) (path string, err error) {
	platform := z.platformName()
	name := fmt.Sprintf("tbx-%s-%s.zip", app.Version, platform)
	l := c.Log().With(esl.String("name", name))
	if err := os.MkdirAll(z.DistPath.Path(), 0755); err != nil {
		return "", err
	}

	path = filepath.Join(z.DistPath.Path(), name)
	buildTimestamp, err := time.Parse(time.RFC3339, app.BuildInfo.Timestamp)
	if err != nil {
		return "", err
	}
	docCtl := c.WithLang(es_lang.Default.CodeString())
	docMc := docCtl.Messages()
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	pkg := zip.NewWriter(f)
	if err = pkg.SetComment(fmt.Sprintf("%s %s, %s, %s", app.Name, app.Version, app.Copyright, app.LandingPage)); err != nil {
		return path, err
	}
	pkg.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// LICENSE.txt
	{
		licenseBody := app_ui.MakeMarkdown(docMc, func(ui app_ui.UI) {
			err = dc_license.Generate(docCtl, ui)
		})
		if err != nil {
			return path, err
		}
		w, err := pkg.CreateHeader(&zip.FileHeader{
			Name:     "LICENSE.txt",
			Method:   zip.Deflate,
			Modified: buildTimestamp,
		})
		if err != nil {
			return "", err
		}
		if _, err = w.Write([]byte(licenseBody)); err != nil {
			return "", err
		}
	}

	// README.txt
	{
		doc := dc_readme.New(dc_index.MediaRepository, docMc, false)
		body := dc_section.Generate(dc_index.MediaRepository, dc_section.LayoutPage, docMc, doc)
		w, err := pkg.CreateHeader(&zip.FileHeader{
			Name:     "README.txt",
			Method:   zip.Deflate,
			Modified: buildTimestamp,
		})
		if err != nil {
			return "", err
		}
		if _, err = w.Write([]byte(body)); err != nil {
			return "", err
		}
	}

	// binary
	{
		binaryName := "tbx"
		binarySuffix := ""
		target, ok := os.LookupEnv(app.EnvNameToolboxBuildTarget)
		if ok {
			switch target {
			case "windows/amd64":
				binarySuffix = ".exe"
				binaryName = "tbx-windows-amd64.exe"
			case "linux/amd64":
				binaryName = "tbx-linux-amd64"
			case "linux/arm64":
				binaryName = "tbx-linux-arm64"
			case "darwin/amd64":
				binaryName = "tbx-darwin-amd64"
			case "darwin/arm64":
				binaryName = "tbx-darwin-arm64"
			}
		}
		binaryPath := filepath.Join(z.BuildPath.Path(), binaryName)
		newBinaryPath := filepath.Join(z.BuildPath.Path(), "tbx"+binarySuffix)
		l.Debug("Renaming binary", esl.String("old", binaryName), esl.String("new", newBinaryPath))
		if newBinaryPath == binaryPath {
			l.Debug("Skip renaming")
		} else {
			err := os.Rename(binaryPath, newBinaryPath)
			if err != nil {
				l.Debug("Unable to rename binary", esl.Error(err))
				return "", err
			}
		}

		info, err := os.Lstat(newBinaryPath)
		if err != nil {
			return "", err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return "", err
		}
		header.Method = zip.Deflate
		header.Modified = buildTimestamp

		bin, err := os.Open(newBinaryPath)
		if err != nil {
			return "", err
		}
		defer func() {
			_ = bin.Close()
		}()

		w, err := pkg.CreateHeader(header)
		if _, err = io.Copy(w, bin); err != nil {
			return "", err
		}
	}

	if err = pkg.Flush(); err != nil {
		return "", err
	}
	if err = pkg.Close(); err != nil {
		return "", err
	}

	l.Info("The package created", esl.String("path", path))

	return path, nil
}

func (z *Package) Exec(c app_control.Control) error {
	pkgPath, err := z.createPackage(c)
	if err != nil {
		return err
	}

	if z.DeployPath.IsExists() {
		ea_indicator.SuppressIndicatorForce()

		return rc_exec.Exec(c, z.Up, func(r rc_recipe.Recipe) {
			m := r.(*artifact.Up)
			m.LocalPath = mo_path.NewFileSystemPath(pkgPath)
			m.DropboxPath = mo_path2.NewDropboxPath(filepath.ToSlash(filepath.Join(z.DeployPath.Value(), app.BuildInfo.Branch, "tbx-"+app.BuildId)))
			m.PeerName = "deploy"
		})
	}

	return nil
}

func (z *Package) Test(c app_control.Control) error {
	dest, err := qt_file.MakeTestFolder("pkg", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(dest)
	}()

	binPath := filepath.Join(dest, "tbx")
	err = os.WriteFile(binPath, []byte("This is test content"), 0644)
	if err != nil {
		return err
	}

	return rc_exec.Exec(c, &Package{}, func(r rc_recipe.Recipe) {
		m := r.(*Package)
		m.DistPath = mo_path.NewFileSystemPath(dest)
		m.BuildPath = mo_path.NewExistingFileSystemPath(dest)
	})
}
