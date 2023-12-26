package release

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/github/api/gh_client"
	"github.com/watermint/toolbox/domain/github/api/gh_client_impl"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_release"
	"github.com/watermint/toolbox/domain/github/model/mo_release_asset"
	"github.com/watermint/toolbox/domain/github/service/sv_reference"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/file/es_filehash"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/ingredient/ig_release/ig_homebrew"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_runtime"
	"github.com/watermint/toolbox/recipe/dev/test"
	"github.com/watermint/toolbox/resources"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrorBuildIsNotReadyForRelease = errors.New("the build does not satisfy release criteria")
	ErrorNoArtifactFound           = errors.New("no artifact found")
)

const (
	homebrewRepoOwner  = "watermint"
	homebrewRepoName   = "homebrew-toolbox"
	homebrewRepoBranch = "master"
)

type Publish struct {
	rc_recipe.RemarkConsole
	rc_recipe.RemarkSecret
	ArtifactPath mo_path2.FileSystemPath
	Branch       string
	ConnGithub   gh_conn.ConnGithubRepo
	SkipTests    bool
	Recipe       *test.Recipe
	Formula      *ig_homebrew.Formula
	Asset        *Asset
	Asseturl     *Asseturl

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
	TagCommitMessage          app_msg.Message
	ReleaseName               app_msg.Message
	ReleaseNameBeta           app_msg.Message
	ReleaseNameDev            app_msg.Message
}

type ArtifactSum struct {
	Filename string
	Size     int64
	MD5      string
	SHA256   string
}

func (z *Publish) Preset() {
	z.Branch = "main"
}

func (z *Publish) artifactAssets(c app_control.Control) (paths []string, sizes map[string]int64, err error) {
	l := c.Log()

	entries, err := os.ReadDir(z.ArtifactPath.Path())
	if err != nil {
		return nil, nil, err
	}
	paths = make([]string, 0)
	sizes = make(map[string]int64)
	for _, e := range entries {
		if !strings.HasPrefix(e.Name(), app_definitions2.ExecutableName+"-"+app_definitions2.BuildId) || !strings.HasSuffix(e.Name(), ".zip") {
			l.Debug("Ignore non artifact file", esl.Any("file", e))
			continue
		}
		path := filepath.Join(z.ArtifactPath.Path(), e.Name())
		paths = append(paths, path)
		info, err := e.Info()
		if err != nil {
			l.Debug("Unable to read file info", esl.Error(err))
			return nil, nil, err
		}
		sizes[path] = info.Size()
	}
	return paths, sizes, nil
}

func (z *Publish) verifyArtifacts(c app_control.Control) (a []*ArtifactSum, err error) {
	l := c.Log()
	a = make([]*ArtifactSum, 0)

	assets, assetSize, err := z.artifactAssets(c)

	h := es_filehash.NewHash(l)
	for _, p := range assets {
		sum := &ArtifactSum{
			Filename: filepath.Base(p),
			Size:     assetSize[p],
		}
		sum.MD5, err = h.MD5(p)
		if err != nil {
			l.Debug("Unable to calc MD5", esl.Error(err))
			return nil, err
		}
		sum.SHA256, err = h.SHA256(p)
		if err != nil {
			l.Debug("Unable to calc SHA256", esl.Error(err))
			return nil, err
		}
		a = append(a, sum)
	}
	if len(a) < 1 {
		return nil, ErrorNoArtifactFound
	}
	return a, nil
}

func (z *Publish) releaseNotes(c app_control.Control, sum []*ArtifactSum) (relNote string, err error) {
	l := c.Log()
	baseUrl := "https://github.com/watermint/toolbox/blob/" + app_definitions2.BuildId
	if app_definitions2.Release == "" {
		l.Error("Release number undefined")
		return "", errors.New("release number undefined")
	}

	md := app_ui.MakeMarkdown(c.WithLang("en").Messages(), func(mui app_ui.UI) {
		mui.Header(z.HeadingReleaseTheme)
		mui.Break()
		mui.Info(app_msg.Raw(resources.ReleaseNotes()))
		mui.Break()

		mui.Header(z.HeadingChanges)

		for _, la := range es_lang.Supported {
			mui.Info(z.ListSpecChange.
				With("Link", baseUrl+fmt.Sprintf("/docs/releases/changes%s.md", app_definitions2.Release)).
				With("Lang", la.Self()),
			)
		}

		mui.Break()
		mui.Header(z.HeadingDocument)

		for _, la := range es_lang.Supported {
			name := "README" + la.Suffix() + ".md"
			mui.Info(z.ListReadme.
				With("Name", name).
				With("Link", baseUrl+"/"+name).
				With("Lang", la.Self()),
			)
		}

		mui.Break()
		mui.Header(z.HeadingBinary)
		mui.WithTable("Binaries", func(mit app_ui.Table) {
			mit.Header(z.BinaryTableHeaderFilename, z.BinaryTableHeaderSize, z.BinaryTableHeaderMD5, z.BinaryTableHeaderSHA256)

			for _, s := range sum {
				mit.RowRaw(s.Filename, fmt.Sprintf("%d", s.Size), s.MD5, s.SHA256)
			}
		})
	})

	relNotesPath := filepath.Join(c.Workspace().Report(), "release_notes.md")
	err = os.WriteFile(relNotesPath, []byte(md), 0644)
	if err != nil {
		l.Debug("Unable to write release notes", esl.Error(err), esl.String("path", relNotesPath))
		return "", err
	}
	l.Info("Release note created", esl.String("path", relNotesPath))

	return md, nil
}

func (z *Publish) endToEndTest(c app_control.Control) error {
	l := c.Log()
	if c.Feature().IsTest() {
		l.Info("Skip tests")
		return nil
	}

	l.Info("Testing all end to end test")
	err := rc_exec.Exec(c, z.Recipe, func(r rc_recipe.Recipe) {
		m := r.(*test.Recipe)
		m.All = true
	})
	return err
}

func (z *Publish) ghCtx(c app_control.Control) gh_client.Client {
	if c.Feature().IsTest() {
		return gh_client_impl.NewMock("mock", c)
	} else {
		return z.ConnGithub.Client()
	}
}

func (z *Publish) createTag(c app_control.Control) error {
	l := c.Log().With(
		esl.String("owner", app_definitions2.RepositoryOwner),
		esl.String("repository", app_definitions2.RepositoryName),
		esl.String("version", app_definitions2.BuildId),
		esl.String("hash", app_definitions2.BuildInfo.Hash))
	svt := sv_reference.New(z.ghCtx(c), app_definitions2.RepositoryOwner, app_definitions2.RepositoryName)
	l.Debug("Create tag")
	tag, err := svt.Create(
		"refs/tags/"+app_definitions2.BuildId,
		app_definitions2.BuildInfo.Hash,
	)
	if err != nil && err != qt_errors.ErrorMock {
		l.Debug("Unable to create tag", esl.Error(err))
		return err
	}
	if err == qt_errors.ErrorMock {
		return nil
	}
	l.Info("The tag created", esl.String("tag", tag.Ref))
	return nil
}

func (z *Publish) createReleaseDraft(c app_control.Control, relNote string) (rel *mo_release.Release, err error) {
	l := c.Log().With(
		esl.String("owner", app_definitions2.RepositoryOwner),
		esl.String("repository", app_definitions2.RepositoryName),
		esl.String("version", app_definitions2.BuildId),
		esl.String("hash", app_definitions2.BuildInfo.Hash))
	ui := c.UI()

	relName := ""
	switch app_definitions2.ReleaseStage() {
	case app_definitions2.StageDev:
		relName = ui.Text(z.ReleaseNameDev.With("Version", app_definitions2.BuildId))
	case app_definitions2.StageBeta:
		relName = ui.Text(z.ReleaseNameBeta.With("Version", app_definitions2.BuildId))
	case app_definitions2.StageRelease:
		relName = ui.Text(z.ReleaseName.With("Version", app_definitions2.BuildId))
	}

	svr := sv_release.New(z.ghCtx(c), app_definitions2.RepositoryOwner, app_definitions2.RepositoryName)
	rel, err = svr.CreateDraft(
		app_definitions2.BuildId,
		relName,
		relNote,
		z.Branch,
	)
	if err != nil && err != qt_errors.ErrorMock {
		l.Debug("Unable to create release draft", esl.Error(err))
		return nil, err
	}
	if err == qt_errors.ErrorMock {
		return &mo_release.Release{}, nil
	}
	l.Info("Release created", esl.String("release", rel.Url))
	return rel, nil
}

func (z *Publish) uploadAssets(c app_control.Control, rel *mo_release.Release) (uploaded map[string]*mo_release_asset.Asset, err error) {
	l := c.Log()
	assets, _, err := z.artifactAssets(c)
	if err != nil {
		return nil, err
	}

	sva := sv_release_asset.New(z.ghCtx(c), app_definitions2.RepositoryOwner, app_definitions2.RepositoryName, rel.Id)

	// filename -> asset info
	uploaded = make(map[string]*mo_release_asset.Asset)
	for _, p := range assets {
		l.Info("Uploading asset", esl.String("path", p))
		a, err := sva.Upload(mo_path2.NewExistingFileSystemPath(p))
		if err != nil && err != qt_errors.ErrorMock {
			return nil, err
		}
		if err == qt_errors.ErrorMock {
			continue
		}
		l.Info("Uploaded", esl.Any("asset", a.Name))
		uploaded[p] = a
	}
	return uploaded, nil
}

func (z *Publish) publishRelease(c app_control.Control, release *mo_release.Release) error {
	l := c.Log()
	svr := sv_release.New(z.ghCtx(c), app_definitions2.RepositoryOwner, app_definitions2.RepositoryName)
	published, err := svr.Publish(release.Id)
	if err != nil {
		l.Warn("Unable to publish the release", esl.Error(err))
		return err
	}

	l.Info("Release published", esl.Any("release", published))
	return nil
}

func (z *Publish) updateHomebrewFormula(c app_control.Control, macIntel, macArm, linuxIntel, linuxArm *mo_release_asset.Asset) error {
	baseUrl := "https://github.com/watermint/toolbox/releases/download/" + app_definitions2.BuildId + "/"
	return rc_exec.Exec(c, z.Formula, func(r rc_recipe.Recipe) {
		m := r.(*ig_homebrew.Formula)
		m.Owner = homebrewRepoOwner
		m.Repository = homebrewRepoName
		m.Branch = homebrewRepoBranch
		m.Message = "Release " + app_definitions2.BuildId
		m.FormulaName = "toolbox.rb"

		m.AssetPathMacIntel = mo_path2.NewExistingFileSystemPath(macIntel.Name)
		m.DownloadUrlMacIntel = baseUrl + macIntel.Name
		m.AssetPathMacArm = mo_path2.NewExistingFileSystemPath(macArm.Name)
		m.DownloadUrlMacArm = baseUrl + macArm.Name
		m.AssetPathLinuxIntel = mo_path2.NewExistingFileSystemPath(linuxIntel.Name)
		m.DownloadUrlLinuxIntel = baseUrl + linuxIntel.Name
		m.AssetPathLinuxArm = mo_path2.NewExistingFileSystemPath(linuxArm.Name)
		m.DownloadUrlLinuxArm = baseUrl + linuxArm.Name
	})
}

func (z *Publish) Exec(c app_control.Control) error {
	l := c.Log()
	ready := true

	if app_definitions2.IsProduction() {
		l.Info("Verify embedded resources")
		qt_runtime.Suite(c)
	} else {
		l.Info("Run as dev mode")
		ready = false
	}

	if z.SkipTests {
		ready = false
	} else {
		err := z.endToEndTest(c)
		if err != nil {
			return err
		}
	}

	sum, err := z.verifyArtifacts(c)
	if err != nil {
		return err
	}

	relNote, err := z.releaseNotes(c, sum)
	if err != nil {
		return err
	}

	if err := z.createTag(c); err != nil {
		return nil
	}

	rel, err := z.createReleaseDraft(c, relNote)
	if err != nil {
		return err
	}

	// assets: filename -> asset data
	assets, err := z.uploadAssets(c, rel)
	if err != nil {
		return err
	}

	if !ready {
		l.Warn("The build does not satisfy release criteria")
		return ErrorBuildIsNotReadyForRelease
	}

	if err := z.publishRelease(c, rel); err != nil {
		return err
	}

	var assetLinuxArm, assetLinuxIntel, assetMacArm, assetMacIntel *mo_release_asset.Asset
	for _, a := range assets {
		switch IdentifyPlatform(a) {
		case AssetPlatformMacIntel:
			assetMacIntel = a
		case AssetPlatformMacArm:
			assetMacArm = a
		case AssetPlatformLinuxIntel:
			assetLinuxIntel = a
		case AssetPlatformLinuxArm:
			assetLinuxArm = a
		}
	}

	l.Info("updating Homebrew formula",
		esl.Any("macIntel", assetMacIntel),
		esl.Any("macArm", assetMacArm),
		esl.Any("linuxIntel", assetLinuxIntel),
		esl.Any("linuxArm", assetLinuxArm),
	)

	if err := z.updateHomebrewFormula(c, assetMacIntel, assetMacArm, assetLinuxIntel, assetLinuxArm); err != nil {
		return err
	}

	err = rc_exec.Exec(c, z.Asseturl, func(r rc_recipe.Recipe) {
		m := r.(*Asseturl)
		m.SourceOwner = app_definitions2.RepositoryOwner
		m.SourceRepo = app_definitions2.RepositoryName
		m.TargetOwner = homebrewRepoOwner
		m.TargetRepo = homebrewRepoName
		m.TargetBranch = homebrewRepoBranch
	})
	if err != nil {
		l.Error("Unable to update asset URL", esl.Error(err))
		return err
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
		app_definitions2.BuildId = "dev-test"
		err = os.WriteFile(filepath.Join(d, app_definitions2.ExecutableName+"-"+app_definitions2.BuildId+"-"+platform+".zip"), []byte("Test artifact"), 0644)
		if err != nil {
			c.Log().Warn("Unable to create test artifact", esl.Error(err))
			return err
		}
	}
	defer os.RemoveAll(d)

	err = rc_exec.ExecMock(c, &Publish{}, func(r rc_recipe.Recipe) {
		m := r.(*Publish)
		m.ArtifactPath = mo_path2.NewFileSystemPath(d)
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); !errors.Is(err, ErrorBuildIsNotReadyForRelease) && err != nil {
		return err
	}
	return nil
}
