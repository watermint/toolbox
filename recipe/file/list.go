package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type ListVO struct {
	PeerName         app_conn.ConnUserFile
	Path             string
	Recursive        bool
	IncludeDeleted   bool
	IncludeMediaInfo bool
}

type List struct {
}

func (z *List) Requirement() app_vo.ValueObject {
	return &ListVO{}
}

func (z *List) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*ListVO)

	ctx, err := vo.PeerName.Connect(k.Control())
	if err != nil {
		return err
	}

	opts := make([]sv_file.ListOpt, 0)
	if vo.IncludeDeleted {
		opts = append(opts, sv_file.IncludeDeleted())
	}
	if vo.IncludeMediaInfo {
		opts = append(opts, sv_file.IncludeMediaInfo())
	}
	if vo.Recursive {
		opts = append(opts, sv_file.Recursive())
	}
	opts = append(opts, sv_file.IncludeHasExplicitSharedMembers())

	rep, err := k.Report("file", &mo_file.ConcreteEntry{})
	if err != nil {
		return err
	}
	defer rep.Close()

	err = sv_file.NewFiles(ctx).ListChunked(mo_path.NewPath(vo.Path), func(entry mo_file.Entry) {
		rep.Row(entry.Concrete())
	}, opts...)
	if err != nil {
		k.Log().Debug("Failed to list files")
		return err
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	lvo := &ListVO{
		Path:      "",
		Recursive: false,
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return app_test.TestRows(c, "file", func(cols map[string]string) error {
		if _, ok := cols["id"]; !ok {
			return errors.New("`id` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
