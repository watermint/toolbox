package release

import (
	"fmt"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/strings/es_version"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/recipe/dev/spec"
	"github.com/watermint/toolbox/resources"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

const (
	MinimumSpecDocVersion = 59
)

type Doc struct {
	rc_recipe.RemarkSecret
	Peer gh_conn.ConnGithubPublic
}

type Metadata struct {
	// Release version
	Version *es_version.Version `json:"release_major"`

	// Tag name like "90.8.991"
	TagName string `json:"version_full" path:"tag_name"`

	// Release page url like "https://github.com/watermint/toolbox/releases/tag/90.8.991"
	ReleasePage string `json:"release_page" path:"html_url"`

	// Release name like "Release 90.8.991"
	ReleaseName string `json:"release_name" path:"name"`

	// Publish date like "2021-04-30T07:24:33Z"
	PublishedAt string `json:"published_at" path:"published_at"`

	// Draft flag
	Draft bool `json:"draft" path:"draft"`

	// Pre release flag
	PreRelease bool `json:"pre_release" path:"prerelease"`

	// Release notes
	ReleaseNotes string `json:"release_notes" path:"body"`
}

type Notes struct {
	Version *es_version.Version `json:"version"`
	Date    time.Time           `json:"date"`
	Notes   string              `json:"notes"`
	Url     string              `json:"url"`
}

func (z *Doc) Preset() {
}

func (z *Doc) updateRelease(c app_control.Control, release *Notes, prjRoot string) error {
	l := c.Log().With(esl.Uint64("Release", release.Version.Major))
	releasePostPath := filepath.Join(prjRoot, "docs", "_posts")
	tmplBytes, err := app_resource.Bundle().Templates().Bytes("release_post.md.tmpl")
	if err != nil {
		l.Debug("Unable to find the release detail", esl.Error(err))
		return err
	}
	tmpl, err := template.New("release_notes").Parse(string(tmplBytes))
	if err != nil {
		l.Debug("Unable to compile the template", esl.Error(err))
		return err
	}
	name := fmt.Sprintf("%s-release-%d.md", release.Date.UTC().Format("2006-01-02"), release.Version.Major)
	f, err := os.Create(filepath.Join(releasePostPath, name))
	if err != nil {
		l.Debug("Unable to create the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	return tmpl.Execute(f, map[string]interface{}{
		"Title":        fmt.Sprintf("Release %d", release.Version.Major),
		"Lang":         "en",
		"ReleasePage":  release.Url,
		"Release":      strconv.FormatUint(release.Version.Major, 10),
		"ReleaseNotes": strings.ReplaceAll(release.Notes, "{{.", "{% raw %}{{.{% endraw %}"),
	})
}

func (z *Doc) updateSpec(c app_control.Control, prjRoot string) error {
	l := c.Log()
	for _, lg := range lang.Supported {
		l.Info("Generating Spec document")
		err := rc_exec.Exec(c, &spec.Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*spec.Doc)
			rr.Lang = mo_string.NewOptional(lg.CodeString())
			rr.FilePath = mo_string.NewOptional(filepath.Join(prjRoot, lg.CodeString(), fmt.Sprintf("spec_%s.json.gz", app.Release)))
		})
		if err != nil {
			l.Error("Failed to generate documents", esl.Error(err))
			return err
		}
	}
	return nil
}

func (z *Doc) updateChanges(c app_control.Control, release *es_version.Version, prjRoot string) error {
	l := c.Log().With(esl.Uint64("Release", release.Major))

	for _, lg := range lang.Supported {
		ll := l.With(esl.String("lang", lg.CodeString()))

		webLangPath := lg.CodeString() + "/"
		if lg.IsDefault() {
			webLangPath = ""
		}

		ll.Info("Generating release changes")
		err := rc_exec.Exec(c, &spec.Diff{}, func(r rc_recipe.Recipe) {
			rr := r.(*spec.Diff)
			rr.DocLang = mo_string.NewOptional(lg.CodeString())
			rr.Release1 = mo_string.NewOptional(fmt.Sprintf("%d", release.Major-1))
			rr.Release2 = mo_string.NewOptional(fmt.Sprintf("%d", release.Major))
			rr.FilePath = mo_string.NewOptional(filepath.Join(prjRoot, "docs", webLangPath, "releases", fmt.Sprintf("changes%d.md", release.Major)))
		})
		if err != nil {
			ll.Error("Failed to generate documents", esl.Error(err))
			return err
		}
	}

	return nil
}

// Returns versionMajor -> release metadata
func (z *Doc) listReleases(c app_control.Control) (releases map[uint64]*Metadata, err error) {
	l := c.Log()
	releasesData, err := sv_release.New(z.Peer.Client(), app.RepositoryOwner, app.RepositoryName).List()
	if err != nil {
		l.Debug("Unable to retrieve release information", esl.Error(err))
		return nil, err
	}

	releases = make(map[uint64]*Metadata)
	for _, rd := range releasesData {
		meta := &Metadata{}
		if err := api_parser.ParseModelRaw(meta, rd.Raw); err != nil {
			l.Debug("Unable to parse", esl.Error(err))
			return nil, err
		}

		ver, err := es_version.Parse(meta.TagName)
		if err != nil {
			l.Debug("Unable to parse the version", esl.Error(err))
			return nil, err
		}

		meta.Version = &ver

		l.Debug("Release", esl.Any("release", meta))

		if meta.Draft || meta.PreRelease {
			l.Debug("Skip pre-release or draft", esl.Any("release", meta))
			continue
		}

		if r, ok := releases[ver.Major]; !ok {
			releases[ver.Major] = meta
		} else if ver.Compare(*r.Version) > 0 {
			releases[ver.Major] = meta
		}
	}
	return releases, nil
}
func (z *Doc) removeOldReleaseNotes(c app_control.Control, base string) error {
	l := c.Log()
	path := filepath.Join(base, "docs/_posts")
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		l.Debug("Unable to read the directory", esl.Error(err))
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		if strings.HasSuffix(entry.Name(), ".md") {
			l.Info("Removing old release note", esl.String("file", entryPath))
			err = os.Remove(entryPath)
			if err != nil {
				l.Warn("Unable to remove the file", esl.Error(err))
				return err
			}
		}
	}

	return nil
}

func (z *Doc) Exec(c app_control.Control) error {
	l := c.Log()
	releases, err := z.listReleases(c)
	if err != nil {
		return err
	}

	prjBase, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}

	l.Info("Update spec", esl.Uint64("Release", app.Version.Major))
	if err := z.updateSpec(c, filepath.Join(prjBase, "resources/release")); err != nil {
		l.Debug("Unable to update spec", esl.Error(err))
		return err
	}
	l.Info("Update changes", esl.Uint64("Release", app.Version.Major))
	if err := z.updateChanges(c, &app.Version, prjBase); err != nil {
		l.Debug("Unable to update changes", esl.Error(err))
		return err
	}

	l.Info("Remove old release notes")
	if err = z.removeOldReleaseNotes(c, prjBase); err != nil {
		l.Debug("Unable to remove old release notes")
		return err
	}

	l.Info("Update release notes", esl.Uint64("Release", app.Version.Major))
	notes := &Notes{
		Version: &app.Version,
		Date:    time.Now(),
		Notes:   resources.ReleaseNotes(),
		Url:     "https://github.com/watermint/toolbox/releases/latest",
	}
	if err := z.updateRelease(c, notes, prjBase); err != nil {
		l.Debug("Unable to update release notes", esl.Error(err))
		return err
	}

	for _, r := range releases {
		l.Info("Release", esl.Any("release", r.Version))
		relDate, err := time.Parse(time.RFC3339, r.PublishedAt)
		if err != nil {
			l.Warn("Unable to parse the release date, skip this release", esl.Error(err))
			continue
		}
		notes := &Notes{
			Version: r.Version,
			Date:    relDate,
			Notes:   r.ReleaseNotes,
			Url:     r.ReleasePage,
		}
		if err := z.updateRelease(c, notes, prjBase); err != nil {
			return err
		}
		if err := z.updateChanges(c, r.Version, prjBase); err != nil {
			return err
		}
	}

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Doc{}, rc_recipe.NoCustomValues)
}
