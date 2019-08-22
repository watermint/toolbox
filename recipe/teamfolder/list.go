package teamfolder

import (
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
)

type ListVO struct {
	PeerName app_conn.ConnBusinessFile
}

func (z *ListVO) Validate(t app_vo.Validator) {
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	// TypeAssertionError will be handled by infra
	var vo interface{} = k.Value()
	fvo := vo.(*ListVO)

	connFile, err := fvo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	folders, err := sv_teamfolder.New(connFile).List()
	if err != nil {
		// ApiError will be reported by infra
		return err
	}

	rep, err := k.Report("teamfolder", &mo_teamfolder.TeamFolder{})
	if err != nil {
		return err
	}
	defer rep.Close()
	for _, folder := range folders {
		k.Log().Debug("Folder", zap.Any("folder", folder))
		rep.Row(folder)
	}

	return nil
}
