package update

import (
	"encoding/csv"
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_spec"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

type ExternalIdVO struct {
	Peer rc_conn.ConnBusinessMgmt
	File fd_file.Feed
}

type ExternalIdRow struct {
	Email      string `json:"email"`
	ExternalId string `json:"external_id"`
}

const (
	reportExternalId = "external_id"
)

type ExternalId struct {
}

func (z *ExternalId) Console() {
}

func (z *ExternalId) Requirement() rc_vo.ValueObject {
	return &ExternalIdVO{}
}

func (z *ExternalId) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*ExternalIdVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	if err := vo.File.Model(k.Control(), &ExternalIdRow{}); err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportExternalId)
	if err != nil {
		return err
	}
	defer rep.Close()

	return vo.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*ExternalIdRow)

		mem, ok := emailToMember[row.Email]
		if !ok {
			rep.Skip(app_msg.M("recipe.member.update.externalid.skip.not_found"), m)
			return nil
		}

		mem.ExternalId = row.ExternalId
		updated, err := sv_member.New(ctx).Update(mem)
		if err != nil {
			rep.Failure(err, row)
			return err
		}
		rep.Success(row, updated)
		return nil
	})
}

func (z *ExternalId) Test(c app_control.Control) error {
	l := c.Log()
	res, found := c.TestResource(rc_spec.Key(z))
	if !found || !res.IsArray() {
		l.Debug("SKIP: Test resource not found")
		return qt_recipe.NotEnoughResource()
	}
	vo := &ExternalIdVO{}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		l.Debug("Skip test")
		return qt_recipe.NotEnoughResource()
	}
	pair := make(map[string]string)
	for _, row := range res.Array() {
		email := row.Get("email").String()
		extid := row.Get("external_id").String() + " " + time.Now().Format("2006-01-02T15-04-05")

		if !api_util.RegexEmail.MatchString(email) {
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
		vo.File = fd_file_impl.NewTestData(dataFile)
		lastErr := z.Exec(rc_kitchen.NewKitchen(c, vo))

		qt_recipe.TestRows(c, reportExternalId, func(cols map[string]string) error {
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

func (z *ExternalId) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportExternalId,
			rp_model.TransactionHeader(&ExternalIdRow{}, &mo_member.Member{}),
		),
	}
}
