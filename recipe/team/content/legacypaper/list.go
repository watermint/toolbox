package legacypaper

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_paper"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type PaperList struct {
	MemberEmail   string `json:"member_email"`
	PaperDocId    string `json:"paper_doc_id"`
	PaperOwner    string `json:"paper_owner"`
	PaperTitle    string `json:"paper_title"`
	PaperRevision int64  `json:"paper_revision"`
}

type MemberDoc struct {
	MemberId    string `json:"member_id"`
	MemberEmail string `json:"member_email"`
	PaperDocId  string `json:"paper_doc_id"`
}

type List struct {
	Peer     dbx_conn.ConnScopedTeam
	Paper    rp_model.RowReport
	FilterBy mo_string.SelectString
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeTeamDataMember,
	)
	z.Paper.SetModel(&PaperList{})
	z.FilterBy.SetOptions("docs_created", "docs_created", "docs_accessed")
}

func (z *List) listMemberDoc(md *MemberDoc, c app_control.Control) error {
	l := c.Log().With(esl.String("memberEmail", md.MemberEmail), esl.String("docId", md.PaperDocId))
	mc := z.Peer.Client().AsMemberId(md.MemberId)
	meta, err := sv_paper.NewLegacy(mc).Metadata(md.PaperDocId, "markdown")
	if err != nil {
		l.Debug("Unable to export data", esl.Error(err))
		return err
	}

	z.Paper.Row(&PaperList{
		MemberEmail:   md.MemberEmail,
		PaperDocId:    md.PaperDocId,
		PaperOwner:    meta.Owner,
		PaperTitle:    meta.Title,
		PaperRevision: meta.Revision,
	})
	return nil
}

func (z *List) listMemberPaper(member *mo_member.Member, c app_control.Control, s eq_sequence.Stage) error {
	l := c.Log().With(esl.String("memberEmail", member.Email))
	q := s.Get("scan_paper")
	mc := z.Peer.Client().AsMemberId(member.TeamMemberId)
	err := sv_paper.NewLegacy(mc).List(z.FilterBy.Value(), func(docId string) {
		q.Enqueue(&MemberDoc{
			MemberId:    member.TeamMemberId,
			MemberEmail: member.Email,
			PaperDocId:  docId,
		})
	})
	if err != nil {
		l.Debug("Unable to retrieve list created", esl.Error(err))
		return err
	}
	return nil
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Paper.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("scan_paper", z.listMemberDoc, c)
		s.Define("scan_member", z.listMemberPaper, c, s)
		scan := s.Get("scan_member")

		for _, member := range members {
			scan.Enqueue(member)
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
