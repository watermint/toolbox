package batch

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"go.uber.org/atomic"
	"strings"
)

type MsgUser struct {
	ProgressScanningUser      app_msg.Message
	ProgressScanningUserEvent app_msg.Message
	ErrorUserNotFound         app_msg.Message
}

var (
	MUser = app_msg.Apply(&MsgUser{}).(*MsgUser)
)

type UserEmail struct {
	Email string `json:"email"`
}

const (
	keySeparator = "/"
	keySeqPrefix = "seq"
)

type UserWorker struct {
	Ctl        app_control.Control
	Context    dbx_context.Context
	StartTime  string
	EndTime    string
	Category   mo_string.OptionalString
	EventCache kv_storage.Storage
	UserEmail  string
}

func (z *UserWorker) Exec() error {
	l := z.Ctl.Log().With(esl.String("UserEmail", z.UserEmail))
	ui := z.Ctl.UI()
	ui.Info(MUser.ProgressScanningUser.With("Email", z.UserEmail))

	member, err := sv_member.New(z.Context).ResolveByEmail(z.UserEmail)
	if err != nil {
		l.Debug("user not found", esl.Error(err))
		ui.Error(MUser.ErrorUserNotFound.With("Email", z.UserEmail).With("Error", err))
		return err
	}

	opts := make([]sv_activity.ListOpt, 0)
	opts = append(opts, sv_activity.AccountId(member.AccountId))
	opts = append(opts, sv_activity.StartTime(z.StartTime))
	opts = append(opts, sv_activity.EndTime(z.EndTime))
	if z.Category.IsExists() {
		opts = append(opts, sv_activity.Category(z.Category.Value()))
	}

	eventSeq := atomic.Int64{}

	return sv_activity.New(z.Context).List(
		func(event *mo_activity.Event) error {
			return z.EventCache.Update(func(kvs kv_kvs.Kvs) error {
				seq := eventSeq.Inc()
				key := strings.Join([]string{event.Timestamp, z.UserEmail, fmt.Sprintf("%d", seq)}, keySeparator)
				app_ui.ShowProgressWithMessage(ui, MUser.ProgressScanningUserEvent)

				if err = kvs.PutJson(key, event.Raw); err != nil {
					l.Debug("Unable to store data", esl.Error(err))
					return err
				}
				return nil
			})
		},
		opts...,
	)
}

type User struct {
	Peer       dbx_conn.ConnBusinessAudit
	StartTime  mo_time.TimeOptional
	EndTime    mo_time.TimeOptional
	Category   mo_string.OptionalString
	Combined   rp_model.RowReport
	User       rp_model.RowReport
	File       fd_file.RowFeed
	EventCache kv_storage.Storage
}

func (z *User) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.Combined.Open(); err != nil {
		return err
	}

	userReps := make(map[string]rp_model.RowReport)

	q := c.NewLegacyQueue()
	err := z.File.EachRow(func(m interface{}, rowIndex int) error {
		e := m.(*UserEmail)

		suffix := es_filepath.Escape(e.Email)
		ur, err := z.User.OpenNew(rp_model.Suffix("_"+suffix), rp_model.NoConsoleOutput())
		if err != nil {
			return err
		}
		userReps[e.Email] = ur

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

	defer func() {
		for _, ur := range userReps {
			ur.Close()
		}
	}()

	if err != nil {
		l.Debug("Failure during reading model", esl.Error(err))
		return err
	}

	return z.EventCache.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEach(func(key string, value []byte) error {
			ks := strings.Split(key, keySeparator)
			switch {
			case len(ks) == 2 && ks[0] == keySeqPrefix:
				return nil
			case len(ks) != 3:
				l.Debug("Invalid key format", esl.String("key", key), esl.Strings("ks", ks))
				return errors.New("invalid key format")
			}
			//ts := ks[0]
			email := ks[1]
			//seq := ks[2]
			ur := userReps[email]

			ll := l.With(esl.String("email", email))
			ev := &mo_activity.Event{}
			if err = api_parser.ParseModelRaw(ev, value); err != nil {
				ll.Debug("Unable to parse model", esl.Error(err))
				return err
			}
			ec := ev.Compatible()
			z.Combined.Row(ec)
			ur.Row(ec)

			return nil
		})
	})
}

func (z *User) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &User{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("user", "john@example.com\nsmith@example.net\n")
		if err != nil {
			return
		}
		m := r.(*User)
		m.File.SetFilePath(f)
	})
}

func (z *User) Preset() {
	z.Combined.SetModel(&mo_activity.Compatible{})
	z.User.SetModel(&mo_activity.Compatible{})
	z.File.SetModel(&UserEmail{})
}
