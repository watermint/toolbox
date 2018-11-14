package cmd_member

import (
	"github.com/watermint/toolbox/model/dbx_member"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/report"
)

type MemberReport struct {
	Report report.Report
}

func (z *MemberReport) HandleFailure(email string, reason string) bool {
	z.Report.Report(
		dbx_member.InviteReport{
			Result: "failure",
			Failure: &dbx_member.FailureReport{
				Email:  email,
				Reason: reason,
			},
		},
	)
	return true
}

func (z *MemberReport) HandleSuccess(member *dbx_profile.Member) bool {
	z.Report.Report(
		dbx_member.InviteReport{
			Result:  "success",
			Success: member,
		},
	)
	return true
}
