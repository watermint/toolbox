package filerequest

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type List struct {
}

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	lvo := vo.(*ListVO)
	conn, err := lvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("file_request", &mo_filerequest.MemberFileRequest{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, member := range members {
		mc := conn.AsMemberId(member.TeamMemberId)
		reqs, err := sv_filerequest.New(mc).List()
		if err != nil {
			return err
		}
		for _, req := range reqs {
			fm := mo_filerequest.NewMemberFileRequest(req, member)
			rep.Row(fm)
		}
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "file_request", func(cols map[string]string) error {
		if _, ok := cols["file_request_id"]; !ok {
			return errors.New("`file_request_id` is not found")
		}
		if _, ok := cols["team_member_id"]; !ok {
			return errors.New("`team_member_id` is not found")
		}
		return nil
	})
}
