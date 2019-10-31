package filerequest

import (
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

type CloneVO struct {
	File app_file.Data
	Peer app_conn.ConnBusinessFile
}

type Clone struct {
}

func (z *Clone) Hidden() {
}

func (z *Clone) Requirement() app_vo.ValueObject {
	return &CloneVO{}
}

func (z *Clone) Exec(k app_kitchen.Kitchen) error {
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
	rep, err := k.Report("clone",
		app_report.TransactionHeader(
			&mo_filerequest.MemberFileRequest{},
			&mo_filerequest.MemberFileRequest{}))
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
			rep.Failure(app_report.MsgInvalidData, fm, nil)
			return nil
		}
		member, ok := emailToMember[strings.ToLower(fm.Email)]
		if !ok {
			rep.Failure(app_msg.M("recipe.team.filerequest.clone.err.no_member_found_for_email"),
				fm, nil)
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
			mo_path.NewPath(fm.Destination),
			opts...,
		)
		if err != nil {
			rep.Failure(app_msg.M("recipe.team.filerequest.clone.err.cannot_create"),
				fm, nil)
			return nil
		} else {
			rep.Success(fm, req)
		}
		return nil
	})
}

func (z *Clone) Test(c app_control.Control) error {
	return qt_test.HumanInteractionRequired()
}
