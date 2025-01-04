package legacypaper

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_paper"
	"github.com/watermint/toolbox/essentials/file/es_filemove"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type PaperExport struct {
	MemberEmail   string `json:"member_email"`
	PaperDocId    string `json:"paper_doc_id"`
	PaperOwner    string `json:"paper_owner"`
	PaperTitle    string `json:"paper_title"`
	PaperRevision int64  `json:"paper_revision"`
	ExportPath    string `json:"export_path"`
}

type Export struct {
	Path     mo_path.FileSystemPath
	FilterBy mo_string.SelectString
	Format   mo_string.SelectString
	Paper    rp_model.RowReport
	Peer     dbx_conn.ConnScopedTeam
}

func (z *Export) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.Paper.SetModel(&PaperExport{})
	z.FilterBy.SetOptions("docs_created", "docs_created", "docs_accessed")
	z.Format.SetOptions("html", "html", "markdown")
}

func (z *Export) exportMemberDoc(md *MemberDoc, c app_control.Control) error {
	l := c.Log().With(esl.String("memberEmail", md.MemberEmail), esl.String("docId", md.PaperDocId))
	mc := z.Peer.Client().AsMemberId(md.MemberId).WithPath(dbx_client.Namespace(md.MemberRootNamespaceId))
	meta, path, err := sv_paper.NewLegacy(mc).Export(md.PaperDocId, z.Format.Value())
	if err != nil {
		l.Debug("Unable to export data", esl.Error(err))
		return err
	}

	exportPath := filepath.Join(z.Path.Path(), md.MemberEmail)
	_, err = os.Lstat(exportPath)
	if os.IsNotExist(err) {
		if mkErr := os.MkdirAll(exportPath, 0755); mkErr != nil {
			l.Debug("Unable to make export path", esl.Error(mkErr))
			return mkErr
		}
	}
	ext := ""
	switch z.Format.Value() {
	case "html":
		ext = ".html"
	case "markdown":
		ext = ".md"
	}
	exportFilePath := filepath.Join(exportPath, md.PaperDocId+ext)
	if mvErr := es_filemove.Move(path.Path(), exportFilePath); mvErr != nil {
		l.Debug("Unable to move the file", esl.Error(mvErr))
		return mvErr
	}

	z.Paper.Row(&PaperExport{
		MemberEmail:   md.MemberEmail,
		PaperDocId:    md.PaperDocId,
		PaperOwner:    meta.Owner,
		PaperTitle:    meta.Title,
		PaperRevision: meta.Revision,
		ExportPath:    exportFilePath,
	})
	return nil
}

func (z *Export) listMemberPaper(member *mo_member.Member, c app_control.Control, s eq_sequence.Stage) error {
	l := c.Log().With(esl.String("memberEmail", member.Email))
	q := s.Get("scan_paper")
	mc := z.Peer.Client().AsMemberId(member.TeamMemberId).WithPath(dbx_client.Namespace(member.Profile().RootNamespaceId))
	err := sv_paper.NewLegacy(mc).List(z.FilterBy.Value(), func(docId string) {
		q.Enqueue(&MemberDoc{
			MemberId:              member.TeamMemberId,
			MemberEmail:           member.Email,
			MemberRootNamespaceId: member.Profile().RootNamespaceId,
			PaperDocId:            docId,
		})
	})
	if err != nil {
		l.Debug("Unable to retrieve list created", esl.Error(err))
		return err
	}
	return nil
}

func (z *Export) Exec(c app_control.Control) error {
	if err := z.Paper.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_paper", z.exportMemberDoc, c)
		s.Define("scan_member", z.listMemberPaper, c, s)
		scan := s.Get("scan_member")

		for _, member := range members {
			scan.Enqueue(member)
		}
	})

	return nil
}

func (z *Export) Test(c app_control.Control) error {
	exportPath, err := qt_file.MakeTestFolder("export", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(exportPath)
	}()

	return rc_exec.Exec(c, &Export{}, func(r rc_recipe.Recipe) {
		m := r.(*Export)
		m.Path = mo_path.NewFileSystemPath(exportPath)
	})
}
