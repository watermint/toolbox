package cmd_member

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/model/dbx_member"
	"go.uber.org/zap"
)

type MemberProvision struct {
	Email     string
	GivenName string
	Surname   string
}

func (z *MemberProvision) InviteMember(silent bool) *dbx_member.InviteMember {
	return &dbx_member.InviteMember{
		MemberEmail:      z.Email,
		MemberGivenName:  z.GivenName,
		MemberSurname:    z.Surname,
		SendWelcomeEmail: !silent,
	}
}

func (z *MemberProvision) UpdateMember() (email string, update *dbx_member.UpdateMember) {
	email = z.Email
	update = &dbx_member.UpdateMember{
		NewGivenName: z.GivenName,
		NewSurname:   z.Surname,
	}
	return
}

type MembersProvision struct {
	ec      *app.ExecContext
	optCsv  string
	Members []*MemberProvision
	Logger  *zap.Logger
}

func (z *MembersProvision) FlagConfig(f *flag.FlagSet) {
	descCsv := z.ec.Msg("cmd.member.provision.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)
}

func (z *MembersProvision) Usage() func(cmd.CommandUsage) {
	return func(u cmd.CommandUsage) {
		z.ec.Msg("cmd.member.provision.usage").WithData(u).Tell()
	}
}

func (z *MembersProvision) Load(args []string) error {
	if z.optCsv != "" {
		return z.loadCsv(z.optCsv)
	}
	if len(args) > 0 {
		return z.loadArgs(args)
	}
	z.ec.Msg("cmd.member.provision.err.nodata").TellError()
	z.Logger.Warn("no csv or argument provided")
	return errors.New("please specify member data")
}

func (z *MembersProvision) loadArgs(args []string) error {
	z.Members = make([]*MemberProvision, 0)
	for _, e := range args {
		z.Members = append(z.Members, &MemberProvision{
			Email: e,
		})
	}
	return nil
}

func (z *MembersProvision) loadCsv(filePath string) error {
	z.Members = make([]*MemberProvision, 0)
	loader := app_io.NewCsvLoader(z.ec, filePath).OnRow(func(cols []string) error {
		mp := &MemberProvision{}

		if len(cols) >= 1 {
			mp.Email = cols[0]
		}
		if len(cols) >= 2 {
			mp.GivenName = cols[1]
		}
		if len(cols) >= 3 {
			mp.Surname = cols[2]
		}

		// skip
		if !api_util.RegexEmail.MatchString(mp.Email) {
			z.ec.Log().Debug("Skip: Email field does not match to the pattern", zap.Strings("cols", cols))
			return nil
		}

		z.Members = append(z.Members, mp)
		return nil
	})
	return loader.Load()
}
