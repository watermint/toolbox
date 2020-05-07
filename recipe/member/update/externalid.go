package update

import (
	"encoding/csv"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_resource"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
	"time"
)

type ExternalIdRow struct {
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
}

type Externalid struct {
	rc_recipe.RemarkIrreversible
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
	return qt_resource.WithResource(z, func(j es_json.Json) error {
		type Data struct {
			Email      string `path:"email"`
			ExternalId string `path:"external_id"`
		}
		pair := make(map[string]string)
		err := j.ArrayEach(func(e es_json.Json) error {
			row := &Data{}
			if err := e.Model(row); err != nil {
				return err
			}

			if !dbx_util.RegexEmail.MatchString(row.Email) {
				l.Error("invalid email address", es_log.String("email", row.Email))
				return errors.New("invalid input")
			}
			pair[row.Email] = row.ExternalId + " " + time.Now().Format("2006-01-02T15-04-05")
			return nil
		})
		if err != nil {
			return err
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

			qtr_endtoend.TestRows(c, "operation_log", func(cols map[string]string) error {
				email := cols["email"]
				extid := cols["external_id"]
				if pair[email] != extid {
					return errors.New("external id was not modified")
				}
				return nil
			})
			return lastErr
		}
	})
}
