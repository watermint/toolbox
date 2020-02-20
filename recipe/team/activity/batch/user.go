package batch

import (
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/kvs/kv_transaction"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
	"math/rand"
)

type MsgUser struct {
	ProgressScanningUser app_msg.Message
	ErrorUserNotFound    app_msg.Message
}

var (
	MUser = app_msg.Apply(&MsgUser{}).(*MsgUser)
)

type UserEmail struct {
	Email string `json:"email"`
}

type UserWorker struct {
	Ctl        app_control.Control
	Context    api_context.Context
	StartTime  string
	EndTime    string
	Category   string
	EventCache kv_storage.Storage
	UserEmail  string
}

func (z *UserWorker) Exec() error {
	l := z.Ctl.Log().With(zap.String("UserEmail", z.UserEmail))
	ui := z.Ctl.UI()
	ui.Info(MUser.ProgressScanningUser.With("Email", z.UserEmail))

	member, err := sv_member.New(z.Context).ResolveByEmail(z.UserEmail)
	if err != nil {
		l.Debug("user not found", zap.Error(err))
		ui.Error(MUser.ErrorUserNotFound.With("Email", z.UserEmail).With("Error", err))
		return err
	}

	return sv_activity.New(z.Context).List(
		func(event *mo_activity.Event) error {
			return z.EventCache.Update(func(tx kv_transaction.Transaction) error {
				kvs, err := tx.Kvs(member.Email)
				if err != nil {
					l.Debug("Unable to create kvs", zap.Error(err))
					return err
				}
				seq, err := kvs.NextSequence()
				if err != nil {
					l.Debug("Unable to generate seq", zap.Error(err))
					// pseudo seq
					seq = rand.Uint64()
				}
				key := fmt.Sprintf("%s-%d", event.Timestamp, seq)

				if err = kvs.PutJson(key, event.Raw); err != nil {
					l.Debug("Unable to store data", zap.Error(err))
					return err
				}
				return nil
			})
		},
		sv_activity.AccountId(member.AccountId),
		sv_activity.StartTime(z.StartTime),
		sv_activity.EndTime(z.EndTime),
		sv_activity.Category(z.Category),
	)
}

type User struct {
	Peer       rc_conn.ConnBusinessAudit
	StartTime  mo_time.TimeOptional
	EndTime    mo_time.TimeOptional
	Category   string
	Combined   rp_model.RowReport
	User       rp_model.RowReport
	UserList   fd_file.RowFeed
	EventCache kv_storage.Storage
}

func (z *User) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.Combined.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	err := z.UserList.EachRow(func(m interface{}, rowIndex int) error {
		e := m.(*UserEmail)
		q.Enqueue(&UserWorker{
			Ctl:        c,
			Context:    z.Peer.Context(),
			StartTime:  z.StartTime.Iso8601(),
			EndTime:    z.EndTime.Iso8601(),
			Category:   z.Category,
			EventCache: z.EventCache,
			UserEmail:  e.Email,
		})
		return nil
	})
	l.Debug("Waiting for workers")
	q.Wait()
	if err != nil {
		l.Debug("Failure during reading model", zap.Error(err))
		return err
	}

	return z.EventCache.View(func(tx kv_transaction.Transaction) error {
		return tx.ForEach(func(name string, kvs kv_kvs.Kvs) error {
			ll := l.With(zap.String("user", name))
			suffix := ut_filepath.Escape(name)
			ur, err := z.User.OpenNew(rp_model.Suffix("_" + suffix))
			if err != nil {
				ll.Debug("Unable to create per user report", zap.Error(err))
				ur = nil
			}
			defer ur.Close()
			return kvs.ForEach(func(key string, value []byte) error {
				ev := &mo_activity.Event{}
				if err = api_parser.ParseModelRaw(ev, value); err != nil {
					ll.Debug("Unable to parse model", zap.Error(err))
					return err
				}
				ec := ev.Compatible()
				z.Combined.Row(ec)
				if ur != nil {
					ur.Row(ec)
				}
				return nil
			})
		})
	})
}

func (z *User) Test(c app_control.Control) error {
	return qt_endtoend.ImplementMe()
}

func (z *User) Preset() {
	z.Combined.SetModel(&mo_activity.Compatible{})
	z.User.SetModel(&mo_activity.Compatible{})
	z.UserList.SetModel(&UserEmail{})
}
