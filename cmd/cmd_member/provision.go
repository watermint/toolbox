package cmd_member

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_member"
	"go.uber.org/zap"
	"io"
	"os"
	"strings"
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
	descCsv := z.ec.Msg("cmd.member.provision.flag.csv").Text()
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
	f, err := os.Open(filePath)
	if err != nil {
		z.ec.Msg("cmd.member.provision.err.open_file").WithData(struct {
			File string
		}{
			File: filePath,
		}).TellError()
		z.Logger.Warn(
			"Unable to open file",
			zap.String("file", filePath),
			zap.Error(err),
		)
		return err
	}
	csv := app_util.NewBomAwareCsvReader(f)

	z.Members = make([]*MemberProvision, 0)
	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			z.ec.Msg("cmd.member.provision.err.cant_read").WithData(struct {
				File string
			}{
				File: filePath,
			}).TellError()
			z.Logger.Warn(
				"Unable to read CSV file",
				zap.String("file", filePath),
				zap.Error(err),
			)
			return err
		}
		if len(cols) < 1 {
			z.ec.Msg("cmd.member.provision.err.no_row").Tell()
			z.Logger.Warn("No column found in the row. Skip")
			continue
		}

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
		if !strings.Contains(mp.Email, "@") {
			continue
		}

		z.Members = append(z.Members, mp)
	}
	return nil
}
