package commands

import (
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/integration/business"
	"github.com/watermint/toolbox/service/report"
	"os"
	"sync"
	"testing"
)

func TestTeamMemberSpace(t *testing.T) {
	infraOpts := infra.InfraOpts{}
	err := infraOpts.Startup()
	if err != nil {
		t.Skip("Skip")
		return
	}
	defer infraOpts.Shutdown()
	token := os.Getenv("TEST_TOKEN_BUSINESS_FILE")
	if token == "" {
		t.Skip("No token for test.")
		return
	}

	wg := &sync.WaitGroup{}
	mq := make(chan *team.TeamMemberInfo)
	rows := make(chan report.ReportRow)
	wg.Add(1)
	go report.WriteCsv(os.Stdout, rows, wg)
	go ReportSpace(token, rows, mq, wg)
	wg.Done()

	err = business.LoadTeamMembers(token, mq)
	if err != nil {
		t.Error(err)
	}

	wg.Wait()

}
