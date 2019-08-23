package sharedlink

import (
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

func (*ListVO) Validate(t app_vo.Validator) {
}

type List struct {
}

func (*List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (*List) Exec(k app_kitchen.Kitchen) error {
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
	rep, err := k.Report("sharedlink", &mo_sharedlink.SharedLinkMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, member := range members {
		mc := conn.AsMemberId(member.TeamMemberId)
		links, err := sv_sharedlink.New(mc).List()
		if err != nil {
			return err
		}
		for _, link := range links {
			lm := mo_sharedlink.NewSharedLinkMember(link, member)
			rep.Row(lm)
		}
	}
	return nil
}
