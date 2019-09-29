package update

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_report"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type EmailVO struct {
	Peer                 app_conn.ConnBusinessMgmt
	File                 app_file.Data
	DontUpdateUnverified bool
}

type EmailRow struct {
	FromEmail string `json:"from_email"`
	ToEmail   string `json:"to_email"`
}

type EmailWorker struct {
	transaction *EmailRow
	vo          *EmailVO
	member      *mo_member.Member
	ctx         api_context.Context
	rep         app_report.Report
	ctl         app_control.Control
}

func (z *EmailWorker) Exec() error {
	ui := z.ctl.UI()
	ui.Info("recipe.member.update.email.progress.updating",
		app_msg.P{
			"EmailFrom": z.transaction.FromEmail,
			"EmailTo":   z.transaction.ToEmail,
		})

	l := z.ctl.Log().With(zap.Any("beforeMember", z.member))

	newEmail := &mo_member.Member{}
	if err := api_parser.ParseModelRaw(newEmail, z.member.Raw); err != nil {
		l.Debug("Unable to clone member data", zap.Error(err))
		z.rep.Failure(app_msg.M("recipe.member.update.email.err.internal_error_clone"), z.transaction, nil)
		return err
	}

	newEmail.Email = z.transaction.ToEmail
	newMember, err := sv_member.New(z.ctx).Update(newEmail)
	if err != nil {
		l.Debug("API returned an error", zap.Error(err))
		z.rep.Failure(app_msg.M("recipe.member.update.email.err.api_error",
			app_msg.P{
				"Error": err.Error(),
			}),
			z.transaction,
			nil)
		return err
	}

	z.rep.Success(z.transaction, newMember)
	return nil
}

type Email struct {
}

func (z *Email) Requirement() app_vo.ValueObject {
	return &EmailVO{
		DontUpdateUnverified: true,
	}
}

func (z *Email) Exec(k app_kitchen.Kitchen) error {
	l := k.Log()
	vo := k.Value().(*EmailVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	err = vo.File.Model(k.Control(), &EmailRow{})
	if err != nil {
		return err
	}

	rep, err := k.Report("update", app_report.TransactionHeader(&EmailRow{}, &mo_member.Member{}))
	if err != nil {
		return err
	}
	defer rep.Close()

	q := k.NewQueue()
	err = vo.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*EmailRow)
		ll := l.With(zap.Any("row", row))

		if row.FromEmail == row.ToEmail {
			ll.Debug("Skip")
			rep.Skip(app_msg.M("recipe.member.quota.update.skip.same_from_to_email"), row, nil)
			return nil
		}

		member, ok := emailToMember[row.FromEmail]
		if !ok {
			ll.Debug("Member not found for email")
			rep.Failure(app_msg.M("recipe.member.quota.update.err.not_found", app_msg.P{
				"Email": row.FromEmail,
			}), row, nil)
			return errors.New("member not found for given email address")
		}

		if !member.EmailVerified && vo.DontUpdateUnverified {
			ll.Debug("Do not update unverified email")
			rep.Skip(app_msg.M("recipe.member.quota.update.skip.unverified_email"), row, nil)
			return nil
		}

		q.Enqueue(&EmailWorker{
			transaction: row,
			vo:          vo,
			member:      member,
			ctx:         ctx,
			rep:         rep,
			ctl:         k.Control(),
		})

		return nil
	})
	q.Wait()
	return err
}

func (z *Email) Test(c app_control.Control) error {
	return nil
}
