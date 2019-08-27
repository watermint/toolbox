package api_test

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	app2 "github.com/watermint/toolbox/legacy/app"
)

const (
	TestPeerName = "test_suite"
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
	ec := app2.NewExecContextForTest()
	if !api_auth_impl.IsCacheAvailable(ec, TestPeerName) {
		return
	}

	au := api_auth_impl.NewLegacy(ec, api_auth_impl.PeerName(TestPeerName))
	ctx, err := au.Auth(tokenType)
	if err != nil {
		return
	}
	test(ctx)
	ec.Shutdown()
}
