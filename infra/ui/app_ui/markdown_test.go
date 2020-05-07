package app_ui

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"io/ioutil"
	"testing"
)

func TestMdImpl(t *testing.T) {
	m := app_msg.CreateMessage("ping", app_msg.P{"Message": "Pong"})
	mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{
		m.Key(): "Ping {{.Message}}",
	})
	lg := es_log.Default()
	c := NewMarkdown(mc, lg, ioutil.Discard, es_dialogue.DenyAll())
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
			t.RowRaw("hello", es_number.New(i).String())
		}
	})
}
