package qt_recipe

import (
	"encoding/csv"
	rice "github.com/GeertJohan/go.rice"
	"github.com/pkg/profile"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/control/app_run_impl"
	"github.com/watermint/toolbox/infra/network/nw_ratelimit"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_conn_impl"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/recpie/app_vo_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_memory"
	"github.com/watermint/toolbox/quality/infra/qt_control_impl"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

const (
	EndToEndPeer       = "end_to_end_test"
	TestTeamFolderName = "watermint-toolbox-test"
)

// Returns conn if v & end to end peer exist. found = true when v is an interface of app_conn.ConnDropboxAPI
func ApplyConn(v interface{}, c app_control.Control) (conn app_conn.ConnDropboxApi, found bool) {
	l := c.Log()
	a := api_auth_impl.NewCached(c, api_auth_impl.PeerName(EndToEndPeer))
	if p, found := os.LookupEnv("TOOLBOX_SKIPENDTOENDTEST"); found {
		if b, _ := strconv.ParseBool(p); b {
			l.Info("Skip end to end test")
			return nil, true
		}
	}

	if _, ok := v.(app_conn.ConnBusinessInfo); ok {
		if _, err := a.Auth(api_auth.DropboxTokenBusinessInfo); err != nil {
			l.Info("BusinessInfo: Skip end to end test")
			return nil, true
		}
		return &app_conn_impl.ConnBusinessInfo{
			PeerName: EndToEndPeer,
		}, true
	}

	if _, ok := v.(app_conn.ConnBusinessFile); ok {
		if _, err := a.Auth(api_auth.DropboxTokenBusinessFile); err != nil {
			l.Info("BusinessFile: Skip end to end test")
			return nil, true
		}
		return &app_conn_impl.ConnBusinessFile{
			PeerName: EndToEndPeer,
		}, true
	}
	if _, ok := v.(app_conn.ConnBusinessAudit); ok {
		if _, err := a.Auth(api_auth.DropboxTokenBusinessAudit); err != nil {
			l.Info("BusinessAudit: Skip end to end test")
			return nil, true
		}
		return &app_conn_impl.ConnBusinessAudit{
			PeerName: EndToEndPeer,
		}, true
	}
	if _, ok := v.(app_conn.ConnBusinessMgmt); ok {
		if _, err := a.Auth(api_auth.DropboxTokenBusinessManagement); err != nil {
			l.Info("BusinessManagement: Skip end to end test")
			return nil, true
		}
		return &app_conn_impl.ConnBusinessMgmt{
			PeerName: EndToEndPeer,
		}, true
	}
	if _, ok := v.(app_conn.ConnUserFile); ok {
		if _, err := a.Auth(api_auth.DropboxTokenFull); err != nil {
			l.Info("UserFull: Skip end to end test")
			return nil, true
		}
		return &app_conn_impl.ConnUserFile{
			PeerName: EndToEndPeer,
		}, true
	}

	return nil, false
}

func ApplyTestPeers(ctl app_control.Control, vo app_vo.ValueObject) bool {
	l := ctl.Log()
	l.Debug("Prepare for applying test peers")

	vc := app_vo_impl.NewValueContainer(vo)
	for k, v := range vc.Values {
		if conn, found := ApplyConn(v, ctl); found {
			l.Debug("Conn found for key", zap.String("k", k), zap.Bool("connApplied", conn != nil))
			if conn == nil {
				return false
			} else {
				vc.Values[k] = conn
			}
		}
	}

	l.Debug("Applying for debug")
	vc.Apply(vo)

	return true
}

func Resources(t *testing.T) (bx, web *rice.Box, mc app_msg_container.Container, ui app_ui.UI) {
	bx = rice.MustFindBox("../../../resources")
	web = rice.MustFindBox("../../../web")

	mc = app_run_impl.NewContainer(bx)
	ui = app_ui.NewConsole(mc, qt_control_impl.NewMessageTest(t), true)
	return
}

func findTestResource() (resource gjson.Result, found bool) {
	l := app_root.Log()
	p, found := os.LookupEnv("TOOLBOX_TESTRESOURCE")
	if !found {
		return gjson.Parse("{}"), false
	}
	l = l.With(zap.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("unable to read file", zap.Error(err))
		return gjson.Parse("{}"), false
	}
	if !gjson.ValidBytes(b) {
		l.Debug("invalid file content", zap.ByteString("resource", b))
		return gjson.Parse("{}"), false
	}
	return gjson.ParseBytes(b), true
}

func TestWithControl(t *testing.T, twc func(ctl app_control.Control)) {
	nw_ratelimit.SetTestMode(true)
	bx, web, mc, ui := Resources(t)

	ctl := app_control_impl.NewSingle(ui, bx, web, mc, false, make([]app_recipe.Recipe, 0))
	cs := ctl.(*app_control_impl.Single)
	if res, found := findTestResource(); found {
		var err error
		ctl, err = cs.NewTestControl(res)
		if err != nil {
			t.Error("Unable to create new test control", err)
			return
		}
	}
	err := ctl.Up(app_control.Test(), app_control.Concurrency(runtime.NumCPU()))
	if err != nil {
		os.Exit(app_control.FatalStartup)
	}
	defer ctl.Down()

	twc(ctl)
}

func TestRecipe(t *testing.T, re app_recipe.Recipe) {
	nw_ratelimit.SetTestMode(true)
	TestWithControl(t, func(ctl app_control.Control) {
		l := ctl.Log()
		l.Debug("Start testing")
		pr := profile.Start(
			profile.ProfilePath(ctl.Workspace().Log()),
			profile.MemProfile,
		)

		err := re.Test(ctl)

		pr.Stop()
		ut_memory.DumpStats(l)

		if err == nil {
			return
		}

		switch err.(type) {
		case *ErrorNoTestRequired:
			l.Info("Skip: No test required for this recipe")

		case *ErrorHumanInteractionRequired:
			l.Info("Skip: Human interaction required for this test")

		case *ErrorNotEnoughResource:
			l.Info("Skip: Not enough resource")

		case *ErrorScenarioTest:
			l.Info("Skip: Implemented as scenario test")

		case *ErrorImplementMe:
			l.Warn("Test is not implemented for this recipe")

		default:
			t.Error(err)
		}
	})
}

type RowTester func(cols map[string]string) error

func TestRows(ctl app_control.Control, reportName string, tester RowTester) error {
	l := ctl.Log().With(zap.String("reportName", reportName))
	csvFile := filepath.Join(ctl.Workspace().Report(), reportName+".csv")

	l.Debug("Start loading report", zap.String("csvFile", csvFile))

	cf, err := os.Open(csvFile)
	if err != nil {
		l.Warn("Unable to open report CSV", zap.Error(err))
		return err
	}
	defer cf.Close()
	csf := csv.NewReader(cf)
	var header []string
	isFirstLine := true

	for {
		cols, err := csf.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			l.Warn("An error occurred during read report file", zap.Error(err))
			return err
		}
		if isFirstLine {
			header = cols
			isFirstLine = false
		} else {
			colMap := make(map[string]string)
			for i, h := range header {
				colMap[h] = cols[i]
			}
			if err := tester(colMap); err != nil {
				l.Warn("Tester returned an error", zap.Error(err), zap.Any("cols", colMap))
				return err
			}
		}
	}

	return nil
}
