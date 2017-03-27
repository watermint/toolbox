package sharedlink

import (
	"github.com/watermint/toolbox/infra"
	"os"
	"testing"
	"time"
)

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
	UpdateSharedLinkForTeam(token, UpdateSharedLinkExpireContext{
		TargetUser: "",
		Expiration: time.Duration(100) * time.Hour * 24,
		Overwrite:  false,
	})
}
