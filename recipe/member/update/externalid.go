package update

import (
	"encoding/csv"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

type ExternalIdRow struct {
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
}

type Externalid struct {
	Peer         dbx_conn.ConnBusinessMgmt
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	SkipNotFound app_msg.Message
}

func (z *Externalid) Preset() {
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
	members, err := sv_member.New(z.Peer.Context()).List()
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
		updated, err := sv_member.New(z.Peer.Context()).Update(mem)
		if err != nil {
			z.OperationLog.Failure(err, row)
			return err
		}
		z.OperationLog.Success(row, updated)
		return nil
	})
}

func (z *Externalid) Test(c app_control.Control) error {
	l := c.Log()
	res, found := c.TestResource(rc_recipe.Key(z))
	if !found || !res.IsArray() {
		l.Debug("SKIP: Test resource not found")
		return qt_errors.ErrorNotEnoughResource
	}

	pair := make(map[string]string)
	for _, row := range res.Array() {
		email := row.Get("email").String()
		extid := row.Get("external_id").String() + " " + time.Now().Format("2006-01-02T15-04-05")

		if !dbx_util.RegexEmail.MatchString(email) {
			l.Error("invalid email address", zap.String("email", email))
			return errors.New("invalid input")
		}
		pair[email] = extid
	}

	// prep csv
	dataFile := filepath.Join(c.Workspace().Test(), "external_id.csv")
	{
		f, err := os.Create(dataFile)
		if err != nil {
			return err
		}
		cw := csv.NewWriter(f)
		if err := cw.Write([]string{"email", "external_id"}); err != nil {
			return err
		}
		for k, v := range pair {
			if err := cw.Write([]string{k, v}); err != nil {
				return err
			}
		}
		cw.Flush()
		f.Close()
	}

	// test
	{
		lastErr := rc_exec.Exec(c, &Externalid{}, func(r rc_recipe.Recipe) {
			rc := r.(*Externalid)
			rc.File.SetFilePath(dataFile)
		})

		qt_recipe.TestRows(c, "operation_log", func(cols map[string]string) error {
			email := cols["email"]
			extid := cols["external_id"]
			if pair[email] != extid {
				return errors.New("external id was not modified")
			}
			return nil
		})
		return lastErr
	}
}
