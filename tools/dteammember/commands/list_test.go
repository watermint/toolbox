package commands

import (
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/service/members"
	"github.com/watermint/toolbox/service/report"
	"os"
	"sync"
	"testing"
)

func TestTeamMemberList(t *testing.T) {
	infraOpts := infra.InfraOpts{}
	err := infraOpts.Startup()
	if err != nil {
		t.Skip("Skip")
		return
	}
	defer infraOpts.Shutdown()
	token := os.Getenv("TEST_TOKEN_BUSINESS_INFO")
	if token == "" {
		t.Skip("No token for test.")
		return
	}

	listOpts := ListOptions{
		Infra:  &infraOpts,
		Status: "all",
	}

	rows := make(chan report.ReportRow)
	wg := sync.WaitGroup{}
	go report.WriteCsv(os.Stdout, rows, &wg)

	err = members.ListMembers(token, rows, listOpts.Status)
	if err != nil {
		t.Error(err)
	}
	wg.Wait()
}
