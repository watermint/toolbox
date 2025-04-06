package app_ui

import (
	"strconv"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
)

func TestConImpl(t *testing.T) {
	m := app_msg.CreateMessage("ping", app_msg.P{"Message": "Pong"})
	mc := app_msg_container_impl.NewSingleWithMessagesForTest(map[string]string{
		m.Key(): "Ping {{.Message}}",
	})
	lg := esl.Default()
	c := NewDiscard(mc, lg)
	c.Header(m)
	c.SubHeader(m)
	c.Info(m)
	c.Error(m)
	c.Break()
	c.AskProceed(m)
	c.AskCont(m)
	c.AskText(m)
	c.AskSecure(m)
	c.Success(m)
	c.Failure(m)
	c.Progress(m)
	c.Quote(m)
	c.Code("puts 'hello'")
	c.IsConsole()
	c.IsWeb()

	it := c.InfoTable("t")
	it.Header(m)
	it.Row(m)
	it.Flush()

	c.WithTable("s", func(t Table) {
		t.HeaderRaw("command", "index")
		for i := 0; i < consoleNumRowsThreshold+1; i++ {
			t.RowRaw("hello", strconv.Itoa(i))
		}
	})
}
