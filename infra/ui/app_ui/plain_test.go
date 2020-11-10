package app_ui

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"strings"
	"testing"
)

func TestMakeConsoleDemo(t *testing.T) {
	m := app_msg.CreateMessage("ping", app_msg.P{"Message": "Pong"})
	mc := app_msg_container_impl.NewSingleWithMessagesForTest(map[string]string{
		m.Key(): "Ping {{.Message}}",
	})
	s := MakeConsoleDemo(mc, func(ui UI) {
		ui.Header(m)
		ui.SubHeader(m)
		ui.Info(m)
		ui.Error(m)
		ui.Break()
		ui.AskProceed(m)
		ui.AskCont(m)
		ui.AskText(m)
		ui.AskSecure(m)
		ui.Success(m)
		ui.Failure(m)
		ui.Progress(m)
		ui.Quote(m)
		ui.Code("puts 'hello'")
		ui.IsConsole()
		ui.IsWeb()

		it := ui.InfoTable("t")
		it.Header(m)
		it.Row(m)
		it.Flush()

		ui.WithTable("s", func(t Table) {
			t.HeaderRaw("command", "index")
			for i := 0; i < consoleNumRowsThreshold+1; i++ {
				t.RowRaw("hello", es_number.New(i).String())
			}
		})
	})
	if s == "" || !strings.HasPrefix(strings.TrimSpace(s), "Ping Pong\n") {
		t.Error(s)
	}
}
