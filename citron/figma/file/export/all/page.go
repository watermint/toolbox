package all

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/domain/figma/model/mo_project"
	"github.com/watermint/toolbox/domain/figma/service/sv_file"
	"github.com/watermint/toolbox/domain/figma/service/sv_project"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Project struct {
	Project            *mo_project.Project
	HasProjectConflict bool
}

type ProjectFile struct {
	Project         *Project
	File            *mo_file.File
	HasFileConflict bool
}

type ProjectFilePage struct {
	ProjectFile *ProjectFile
	Node        mo_file.NodeWithPath
}

func (z ProjectFilePage) Folder() string {
	var prj string
	if z.ProjectFile.Project.HasProjectConflict {
		prj = z.ProjectFile.Project.Project.Name + " (" + z.ProjectFile.Project.Project.Id + ")"
	} else {
		prj = z.ProjectFile.Project.Project.Name
	}

	return prj
}

// Path generates relative path of file/page from export root.
func (z ProjectFilePage) Path(ext string) string {
	var file string
	if z.ProjectFile.HasFileConflict {
		file = z.ProjectFile.File.Name + " (" + z.ProjectFile.File.Key + ")"
	} else {
		file = z.ProjectFile.File.Name
	}

	return filepath.Join(z.Folder(), file+"-"+z.Node.Path(" ", "-")+"."+ext)
}

type Page struct {
	Peer   fg_conn.ConnFigmaApi
	TeamId string
	Scale  mo_int.RangeInt
	Format mo_string.SelectString
	Path   mo_path.ExistingFileSystemPath
}

func (z *Page) Preset() {
	z.Scale.SetRange(sv_file.ImageScaleMin, sv_file.ImageScaleMax, sv_file.ImageScaleDefault)
	z.Format.SetOptions("pdf", sv_file.ImageFormats...)
}

func (z *Page) scanProject(teamId string, s eq_sequence.Stage) error {
	q := s.Get("scan_files")
	projects, err := sv_project.New(z.Peer.Client()).List(teamId)
	if err != nil {
		return err
	}
	for i, p := range projects {
		hasConflict := false
		for j, r := range projects {
			if i != j && p.Name == r.Name {
				hasConflict = true
			}
		}
		q.Enqueue(&Project{
			Project:            p,
			HasProjectConflict: hasConflict,
		})
	}
	return nil
}

func (z *Page) scanFiles(project *Project, s eq_sequence.Stage) error {
	q := s.Get("scan_pages")
	files, err := sv_project.New(z.Peer.Client()).Files(project.Project.Id)
	if err != nil {
		return err
	}
	for i, file := range files {
		hasConflict := false
		for j, r := range files {
			if i != j && file.Name == r.Name {
				hasConflict = true
			}
		}
		q.Enqueue(&ProjectFile{
			Project:         project,
			File:            file,
			HasFileConflict: hasConflict,
		})
	}
	return nil
}

func (z *Page) scanPages(file *ProjectFile, s eq_sequence.Stage, c app_control.Control) error {
	q := s.Get("download_page")
	doc, err := sv_file.New(z.Peer.Client()).Info(file.File.Key)
	if err != nil {
		return err
	}
	docLastModified, err := dbx_util.Parse(doc.LastModified)
	if err != nil {
		return err
	}

	pages := doc.NodesWithPathByType("CANVAS")
	for _, page := range pages {
		pfp := &ProjectFilePage{
			ProjectFile: file,
			Node:        page,
		}
		if len(page.Node.Children) < 1 {
			c.Log().Debug("No children found. Skip this page", esl.Any("node", page.Node))
			return nil
		}

		pageFolder := filepath.Join(z.Path.Path(), pfp.Folder())
		_, err := os.Lstat(pageFolder)
		if os.IsNotExist(err) {
			err = os.MkdirAll(pageFolder, 0755)
			if err != nil {
				return err
			}
		}

		pagePath := filepath.Join(z.Path.Path(), pfp.Path(z.Format.Value()))
		stat, err := os.Stat(pagePath)
		switch err {
		case nil:
			if stat.ModTime().Before(docLastModified) {
				q.Enqueue(pfp)
			} else {
				c.Log().Debug("Download skip",
					esl.String("path", pagePath),
					esl.Time("existingFileTime", stat.ModTime()),
					esl.Time("documentLastModified", docLastModified),
					esl.Any("page", pfp))
			}
		default:
			q.Enqueue(pfp)
		}
	}
	return nil
}

func (z *Page) downloadPage(page *ProjectFilePage, c app_control.Control) error {
	svf := sv_file.New(z.Peer.Client())
	urls, err := svf.Image(page.ProjectFile.File.Key, page.Node.Node.Id, z.Scale.Value(), z.Format.Value())
	if err != nil {
		return err
	}
	url, ok := urls[page.Node.Node.Id]
	if !ok || url == "" {
		return errors.New("no image was generated on the Figma side")
	}

	path := filepath.Join(z.Path.Path(), page.Path(z.Format.Value()))

	err = es_download.Download(c.Log(), url, path)
	return err
}

func (z *Page) Exec(c app_control.Control) error {
	if r, m := sv_project.VerifyTeamId(z.TeamId); r != sv_project.VerifyTeamIdLooksOkay {
		c.UI().Error(m)
		return errors.New(c.UI().Text(m))
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("download_page", z.downloadPage, c)
		s.Define("scan_pages", z.scanPages, s, c)
		s.Define("scan_files", z.scanFiles, s)
		s.Define("scan_projects", z.scanProject, s)

		scanProject := s.Get("scan_projects")
		scanProject.Enqueue(z.TeamId)

	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}))

	return lastErr
}

func (z *Page) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("page", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	return rc_exec.ExecMock(c, &Page{}, func(r rc_recipe.Recipe) {
		m := r.(*Page)
		m.TeamId = "1234"
		m.Path = mo_path.NewExistingFileSystemPath(p)
	})
}
