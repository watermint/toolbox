package batch

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type MsgEmail struct {
	ProgressUpdate app_msg.Message
}

var (
	MEmail = app_msg.Apply(&MsgEmail{}).(*MsgEmail)
)

type EmailRow struct {
	FromEmail string `json:"from_email"`
	ToEmail   string `json:"to_email"`
}

type EmailWorker struct {
	transaction *EmailRow
	member      *mo_member.Member
	ctx         dbx_client.Client
	rep         rp_model.TransactionReport
	ctl         app_control.Control
}

func (z *EmailWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Progress(MEmail.ProgressUpdate.With("EmailFrom", z.transaction.FromEmail).With("EmailTo", z.transaction.ToEmail))

	l := z.ctl.Log().With(esl.Any("beforeMember", z.member))

	newEmail := &mo_member.Member{}
	if err := api_parser.ParseModelRaw(newEmail, z.member.Raw); err != nil {
		l.Debug("Unable to clone member data", esl.Error(err))
		z.rep.Failure(err, z.transaction)
		return err
	}

	newEmail.Email = z.transaction.ToEmail
	newMember, err := sv_member.New(z.ctx).Update(newEmail)
	if err != nil {
		l.Debug("API returned an error", esl.Error(err))
		z.rep.Failure(err, z.transaction)
		return err
	}

	z.rep.Success(z.transaction, newMember)
	return nil
}

type EmailUpdate struct {
	Row    *EmailRow
	Member *mo_member.Member
}

type Email struct {
	rc_recipe.RemarkIrreversible
	Peer                dbx_conn.ConnScopedTeam
	File                fd_file.RowFeed
	UpdateUnverified    bool
	OperationLog        rp_model.TransactionReport
	SkipSameFromToEmail app_msg.Message
	SkipUnverifiedEmail app_msg.Message
}

func (z *Email) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeMembersRead,
		dbx_auth.ScopeMembersWrite,
	)
	z.File.SetModel(&EmailRow{})
	z.OperationLog.SetModel(
		&EmailRow{},
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

func (z *Email) update(u *EmailUpdate, c app_control.Control) error {
	l := c.Log().With(esl.Any("beforeMember", u.Member))

	newEmail := &mo_member.Member{}
	if err := api_parser.ParseModelRaw(newEmail, u.Member.Raw); err != nil {
		l.Debug("Unable to clone member data", esl.Error(err))
		z.OperationLog.Failure(err, u.Row)
		return err
	}

	newEmail.Email = u.Row.ToEmail
	newMember, err := sv_member.New(z.Peer.Client()).Update(newEmail)
	if err != nil {
		l.Debug("API returned an error", esl.Error(err))
		z.OperationLog.Failure(err, u.Row)
		return err
	}

	z.OperationLog.Success(u.Row, newMember)
	return nil
}

func (z *Email) Exec(c app_control.Control) error {
	l := c.Log()
	ctx := z.Peer.Client()

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	err = z.OperationLog.Open()
	if err != nil {
		return err
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("update", z.update, c)
		q := s.Get("update")
		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			row := m.(*EmailRow)
			ll := l.With(esl.Any("row", row))

			if row.FromEmail == row.ToEmail {
				ll.Debug("Skip")
				z.OperationLog.Skip(z.SkipSameFromToEmail, row)
				return nil
			}

			member, ok := emailToMember[row.FromEmail]
			if !ok {
				ll.Debug("Member not found for email")
				z.OperationLog.Failure(errors.New("member not found for email"), row)
				return nil
			}

			if !member.EmailVerified && !z.UpdateUnverified {
				ll.Debug("Do not update unverified email")
				z.OperationLog.Skip(z.SkipUnverifiedEmail, row)
				return nil
			}

			q.Enqueue(&EmailUpdate{
				Row:    row,
				Member: member,
			})
			return nil
		})

	})
	return lastErr
}

func (z *Email) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("update", "john@example.com,john@example.net\nemma@example.com,emma@example.net\n")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	return rc_exec.ExecMock(c, &Email{}, func(r rc_recipe.Recipe) {
		m := r.(*Email)
		m.File.SetFilePath(f)
	})
}
