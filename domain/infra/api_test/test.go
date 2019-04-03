package api_test

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/model/mo_path"
)

const (
	testPeerName = "test_suite"
)

var (
	ToolboxTestSuiteFolder = mo_path.NewPath("/toolbox-testsuite")
)

func DoTestTokenFull(test func(ctx api_context.Context)) {
	doTest(api_auth.DropboxTokenFull, test)
}
func DoTestBusinessInfo(test func(ctx api_context.Context)) {
	doTest(api_auth.DropboxTokenBusinessInfo, test)
}
func DoTestBusinessFile(test func(ctx api_context.Context)) {
	doTest(api_auth.DropboxTokenBusinessFile, test)
}
func DoTestBusinessManagement(test func(ctx api_context.Context)) {
	doTest(api_auth.DropboxTokenBusinessManagement, test)
}
func DoTestBusinessAudit(test func(ctx api_context.Context)) {
	doTest(api_auth.DropboxTokenBusinessAudit, test)
}

func doTest(tokenType string, test func(ctx api_context.Context)) {
	ec := app.NewExecContextForTest()
	if !api_auth_impl.IsCacheAvailable(ec, testPeerName) {
		return
	}

	au := api_auth_impl.New(ec, api_auth_impl.PeerName(testPeerName))
	ctx, err := au.Auth(tokenType)
	if err != nil {
		return
	}
	test(ctx)
	ec.Shutdown()
}
