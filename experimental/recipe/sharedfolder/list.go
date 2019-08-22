package sharedfolder

import (
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type ListVO struct {
	PeerName app_conn.ConnUserFile
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

	k.Log().Debug("Scanning folders")
	folders, err := sv_sharedfolder.New(conn).List()
	if err != nil {
		return err
	}

	// Write report
	rep, err := k.Report("sharedfolder", &mo_sharedfolder.SharedFolder{})
	if err != nil {
		return err
	}
	defer rep.Close()

	for _, folder := range folders {
		rep.Row(folder)
	}
	return nil
}
