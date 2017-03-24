package toolbox

import (
	"fmt"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/integration/business"
	"github.com/watermint/toolbox/service/compare"
	"github.com/watermint/toolbox/service/members"
	"github.com/watermint/toolbox/service/report"
	"github.com/watermint/toolbox/service/sharedlink"
	"github.com/watermint/toolbox/service/upload"
	"github.com/watermint/toolbox/tools/dteammember/commands"
	"os"
	"path"
	"sync"
	"testing"
	"time"
)

func TestAvailableTools(t *testing.T) {
	AvailableTools()
}

func TestUploadAndCompare(t *testing.T) {
	infraOpts := infra.InfraOpts{}
	err := infraOpts.Startup()
	if err != nil {
		t.Skip("Skip")
		return
	}
	defer infraOpts.Shutdown()
	token := os.Getenv("TEST_TOKEN_DROPBOX_FULL")
	if token == "" {
		t.Skip("No token for test.")
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Error("Could not acquire wd", err)
		return
	}

	localBasePath := path.Join(wd, "infra")
	dbxTestPath := "/test"
	dbxTestSession := fmt.Sprintf("%x", time.Now().Unix())
	dbxBasePath := path.Join(dbxTestPath, dbxTestSession)

	uc := &upload.UploadContext{
		LocalRecursive:     true,
		LocalFollowSymlink: false,
		DropboxBasePath:    dbxBasePath,
		DropboxToken:       token,
		BandwidthLimit:     0,
	}
	upload.Upload([]string{localBasePath}, uc, 1)

	co := compare.CompareOpts{
		InfraOpts:       &infraOpts,
		ReportOpts:      &report.MultiReportOpts{},
		DropboxToken:    token,
		DropboxBasePath: dbxBasePath,
		LocalBasePath:   localBasePath,
	}
	match, err := compare.Compare(&co)
	if err != nil {
		t.Error(err)
		return
	}
	if !match {
		t.Error("Contents are not matched")
	}
}

func TestUpdateSharedLink(t *testing.T) {
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
	sharedlink.UpdateSharedLinkForTeam(token, sharedlink.UpdateSharedLinkExpireContext{
		TargetUser: "",
		Expiration: time.Duration(100) * time.Hour * 24,
		Overwrite:  false,
	})
}

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

	listOpts := commands.ListOptions{
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
	go commands.ReportSpace(token, rows, mq, wg)
	wg.Done()

	err = business.LoadTeamMembers(token, mq)
	if err != nil {
		t.Error(err)
	}

	wg.Wait()

}
