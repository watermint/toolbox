package update

import (
	"bufio"
	"encoding/csv"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_file_impl"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type EmailVO struct {
	Peer             app_conn.ConnBusinessMgmt
	File             app_file.Data
	UpdateUnverified bool
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
	rep         rp_model.Report
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
		z.rep.Failure(err, z.transaction)
		return err
	}

	newEmail.Email = z.transaction.ToEmail
	newMember, err := sv_member.New(z.ctx).Update(newEmail)
	if err != nil {
		l.Debug("API returned an error", zap.Error(err))
		z.rep.Failure(err, z.transaction)
		return err
	}

	z.rep.Success(z.transaction, newMember)
	return nil
}

const (
	reportEmail = "update"
)

type Email struct {
}

func (z *Email) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportEmail, rp_model.TransactionHeader(&EmailRow{}, &mo_member.Member{})),
	}
}

func (z *Email) Requirement() app_vo.ValueObject {
	return &EmailVO{
		UpdateUnverified: false,
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

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportEmail)
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
			rep.Failure(&rp_model.NotFound{Id: row.FromEmail}, row)
			return nil
		}

		if !member.EmailVerified && !vo.UpdateUnverified {
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
	l := c.Log()
	res, found := c.TestResource(app_recipe.Key(z))
	if !found || !res.IsArray() {
		l.Debug("SKIP: Test resource not found")
		return qt_test.NotEnoughResource()
	}
	vo := &EmailVO{}
	if !app_test.ApplyTestPeers(c, vo) {
		l.Debug("Skip test")
		return qt_test.NotEnoughResource()
	}

	pair := make(map[string]string)
	noExist := make(map[string]bool)

	for _, row := range res.Array() {
		from := row.Get("from").String()
		to := row.Get("to").String()
		exists := row.Get("exists").Bool()

		if !api_util.RegexEmail.MatchString(from) || !api_util.RegexEmail.MatchString(to) {
			l.Error("from or to email address unmatched to email address format", zap.String("from", from), zap.String("to", to))
			return errors.New("invalid input")
		}
		pair[from] = to
		noExist[from] = !exists
	}

	createCsv := func(path string, reverse bool) error {
		l.Info("Create test file", zap.String("path", path))
		f, err := os.Create(path)
		if err != nil {
			l.Debug("Unable to create test file", zap.Error(err))
			return err
		}
		cw := csv.NewWriter(f)
		if err := cw.Write([]string{"from_email", "to_email"}); err != nil {
			return err
		}

		for k, v := range pair {
			if reverse {
				if err := cw.Write([]string{v, k}); err != nil {
					return err
				}
			} else {
				if err := cw.Write([]string{k, v}); err != nil {
					return err
				}
			}
		}
		cw.Flush()
		return f.Close()
	}

	pathForward := filepath.Join(c.Workspace().Test(), "testdata_forward.csv")
	pathBackward := filepath.Join(c.Workspace().Test(), "testdata_backward.csv")

	if err := createCsv(pathForward, false); err != nil {
		l.Error("Unable to create test file", zap.String("pathForward", pathForward), zap.Error(err))
		return err
	}
	if err := createCsv(pathBackward, true); err != nil {
		l.Error("Unable to create test file", zap.String("pathForward", pathForward), zap.Error(err))
		return err
	}

	var lastErr error

	preserveReport := func(suffix string) error {
		repPath := c.Workspace().Report() + suffix
		err := os.Rename(c.Workspace().Report(), repPath)
		if err != nil {
			l.Warn("Unable to preserve forward report", zap.Error(err))
			repPath = c.Workspace().Report()
		}

		// create alt report folder
		err = os.MkdirAll(c.Workspace().Report(), 0701)
		if err != nil {
			l.Error("Unable to create workspace path", zap.Error(err))
			return err
		}
		return nil
	}

	scanReport := func() {
		resultPath := filepath.Join(c.Workspace().Report(), "update.json")
		resultFile, err := os.Open(resultPath)
		if err != nil {
			l.Warn("Unable to open", zap.Error(err))
		} else {
			scanner := bufio.NewScanner(resultFile)
			for scanner.Scan() {
				row := gjson.Parse(scanner.Text())

				status := row.Get("status").String()
				reason := row.Get("reason").String()
				inputFrom := row.Get("input.from_email").String()
				inputTo := row.Get("input.to_email").String()
				resultEmail := row.Get("result.email").String()

				ll := l.With(
					zap.String("status", status),
					zap.String("inputFrom", inputFrom),
					zap.String("inputTo", inputTo),
					zap.String("resultEmail", resultEmail),
					zap.String("reason", reason),
				)
				isNonExistent := noExist[inputFrom] || noExist[inputTo]

				ll.Info("Data file row", zap.Bool("isNonExist", isNonExistent))

				switch {
				case status == "Failure" && isNonExistent:
					ll.Info("Successfully failed for non existent")
				case status == "Failure":
					ll.Warn("Unexpected failure")
					lastErr = errors.New("unexpected failure")
				case status == "Success" && isNonExistent:
					ll.Warn("Unexpected failure")
					lastErr = errors.New("unexpected failure")
				case status == "Success":
					if inputTo == resultEmail {
						ll.Info("Successfully changed for non existent")
					} else {
						ll.Warn("Email address unchanged")
						lastErr = errors.New("email address unchanged")
					}
				default:
					ll.Warn("Unexpected status")
					lastErr = errors.New("unexpected status")
				}
			}
		}
	}

	// forward
	{
		vo.UpdateUnverified = true
		vo.File = app_file_impl.NewTestData(pathForward)

		lastErr = z.Exec(app_kitchen.NewKitchen(c, vo))
		if lastErr != nil {
			l.Warn("Error in forward operation")
		}
		scanReport()
		if err := preserveReport("_forward"); err != nil {
			return err
		}
	}

	// backward
	{
		vo.UpdateUnverified = true
		vo.File = app_file_impl.NewTestData(pathBackward)

		lastErr = z.Exec(app_kitchen.NewKitchen(c, vo))
		if lastErr != nil {
			l.Warn("Error in backward operation")
		}
		scanReport()
		if err := preserveReport("_backward"); err != nil {
			return err
		}
	}

	return lastErr
}
