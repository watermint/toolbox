package cmd_member

import (
	"github.com/watermint/toolbox/app/util"
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
	Members []*MemberProvision
	Logger  *zap.Logger
}

func (z *MembersProvision) LoadCsv(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		z.Logger.Warn(
			"Unable to open file",
			zap.String("file", filePath),
			zap.Error(err),
		)
		return err
	}
	csv := util.NewBomAwareCsvReader(f)

	z.Members = make([]*MemberProvision, 0)
	for {
		cols, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			z.Logger.Warn(
				"Unable to read CSV file",
				zap.String("file", filePath),
				zap.Error(err),
			)
			return err
		}
		if len(cols) < 1 {
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
