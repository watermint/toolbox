package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
	"strings"
)

type CmdMemberSync struct {
	*cmd.SimpleCommandlet
	optRemove string
	optWipe   bool
	optSilent bool
	optCsv    string
	report    app_report.Factory
}

func (z *CmdMemberSync) Name() string {
	return "sync"
}

func (z *CmdMemberSync) Desc() string {
	return "cmd.member.sync.desc"
}

func (z *CmdMemberSync) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdMemberSync) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descSilent := z.ExecContext.Msg("cmd.member.sync.flag.silent").T()
	f.BoolVar(&z.optSilent, "silent", false, descSilent)

	// first release includes only invite/update
	//descRemove := "Action for missing member (none|remove|detach)"
	//f.StringVar(&z.optRemove, "remove-action", "none", descRemove)
	//
	//descWipe := "Wipe data on remove user"
	//f.BoolVar(&z.optWipe, "wipe", false, descWipe)
}

func (z *CmdMemberSync) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}
	svm := sv_member.New(ctx)

	invite := func(email string, cols []string) error {
		opts := make([]sv_member.AddOpt, 0)
		if len(cols) >= 2 {
			givenName := cols[1]
			opts = append(opts, sv_member.AddWithGivenName(givenName))
		}
		if len(cols) >= 3 {
			surname := cols[2]
			opts = append(opts, sv_member.AddWithSurname(surname))
		}
		if z.optSilent {
			opts = append(opts, sv_member.AddWithoutSendWelcomeEmail())
		}
		member, err := svm.Add(email, opts...)
		if err != nil {
			ctx.ErrorMsg(err).TellError()
			return err
		}
		z.report.Report(member)
		return nil
	}
	update := func(member *mo_member.Member, cols []string) error {
		if len(cols) >= 2 {
			givenName := cols[1]
			member.GivenName = givenName
		}
		if len(cols) >= 3 {
			surname := cols[2]
			member.Surname = surname
		}
		updated, err := svm.Update(member)
		if err != nil {
			ctx.ErrorMsg(err).TellError()
			return err
		}
		z.report.Report(updated)
		return nil
	}

	err = app_io.NewCsvLoader(z.ExecContext, z.optCsv).
		OnRow(func(cols []string) error {
			if len(cols) < 1 {
				return nil
			}
			email := strings.TrimSpace(cols[0])
			if !api_util.RegexEmail.MatchString(email) {
				z.Log().Debug("skip: the data is not looking alike an email address", zap.String("email", email))
				return nil
			}
			member, err := svm.ResolveByEmail(email)
			if err != nil {
				return invite(email, cols)
			} else {
				return update(member, cols)
			}
		}).
		Load()

	if err != nil {
		z.Log().Debug("Unable to load", zap.Error(err))
	}
}
