package api_test

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_context_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/model/dbx_auth"
)

const (
	testPeerName = "test_suite"
)

var (
	ToolboxTestSuiteFolder = mo_path.NewPath("/toolbox-testsuite")
)

func DoTestTokenFull(test func(ctx api_context.Context), opts ...api_context.Option) {
	doTest(dbx_auth.DropboxTokenFull, test, opts...)
}

func doTest(tokenType string, test func(ctx api_context.Context), opts ...api_context.Option) {
	ec := app.NewExecContextForTest()
	if !dbx_auth.IsCacheAvailable(ec, testPeerName) {
		return
	}
	au := dbx_auth.NewAuth(ec, testPeerName)
	dt, err := au.Auth(tokenType)
	if err != nil {
		return
	}
	ctx := api_context_impl.New(ec, api_auth_impl.NewCompatible(dt.Token), opts...)
	test(ctx)

	ec.Shutdown()
}
