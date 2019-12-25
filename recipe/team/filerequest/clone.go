package filerequest

import (
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"strings"
)

type CloneVO struct {
	File fd_file.ModelFile
	Peer rc_conn.OldConnBusinessFile
}

const (
	reportClone = "clone"
)

type Clone struct {
}

func (z *Clone) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportClone, rp_model.TransactionHeader(
			&mo_filerequest.MemberFileRequest{},
			&mo_filerequest.MemberFileRequest{})),
	}
}

func (z *Clone) Hidden() {
}

func (z *Clone) Requirement() rc_vo.ValueObject {
	return &CloneVO{}
}

func (z *Clone) Exec(k rc_kitchen.Kitchen) error {
	cvo := k.Value().(*CloneVO)

	conn, err := cvo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	// Write report
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportClone)
	if err != nil {
		return err
	}
	defer rep.Close()

	if err := cvo.File.Model(k.Control(), &mo_filerequest.MemberFileRequest{}); err != nil {
		return err
	}

	return cvo.File.EachRow(func(m interface{}, rowIndex int) error {
		fm := m.(*mo_filerequest.MemberFileRequest)
		if fm.Email == "" || fm.Destination == "" || fm.Title == "" {
			rep.Failure(&rp_model.InvalidData{}, fm)
			return nil
		}
		member, ok := emailToMember[strings.ToLower(fm.Email)]
		if !ok {
			rep.Failure(&rp_model.NotFound{Id: fm.Email}, fm)
			return nil
		}

		opts := make([]sv_filerequest.UpdateOpt, 0)
		if fm.Deadline != "" {
			opts = append(opts, sv_filerequest.OptDeadline(fm.Deadline))
		}
		if fm.DeadlineAllowLateUploads != "" {
			opts = append(opts, sv_filerequest.OptAllowLateUploads(fm.DeadlineAllowLateUploads))
		}
		req, err := sv_filerequest.New(conn.AsMemberId(member.TeamMemberId)).Create(
			fm.Title,
			mo_path.NewDropboxPath(fm.Destination),
			opts...,
		)
		if err != nil {
			rep.Failure(err, fm)
		} else {
			rep.Success(fm, req)
		}
		return nil
	})
}

func (z *Clone) Test(c app_control.Control) error {
	return qt_recipe.HumanInteractionRequired()
}
