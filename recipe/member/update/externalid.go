package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type ExternalIdRow struct {
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
}

type Externalid struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	SkipNotFound app_msg.Message
}

func (z *Externalid) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.File.SetModel(&ExternalIdRow{})
	z.OperationLog.SetModel(
		&ExternalIdRow{},
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"result.team_member_id",
			"result.familiar_name",
			"result.abbreviated_name",
			"result.member_folder_id",
			"result.external_id",
			"result.account_id",
			"result.persistent_id",
		),
	)
}

func (z *Externalid) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*ExternalIdRow)

		mem, ok := emailToMember[row.Email]
		if !ok {
			z.OperationLog.Skip(z.SkipNotFound, m)
			return nil
		}

		mem.ExternalId = row.ExternalId
		updated, err := sv_member.New(z.Peer.Client()).Update(mem)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}
		z.OperationLog.Success(row, updated)
		return nil
	})
}

func (z *Externalid) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("external_id", "john@example.com,EMAIL john@example.com\nemma@example.com EMAIL emma@example.com\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Externalid{}, func(r rc_recipe.Recipe) {
		m := r.(*Externalid)
		m.File.SetFilePath(f)
	})
}
