package app_ui

import (
	"github.com/watermint/toolbox/essentials/concurrency/es_mutex"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
)

type proxyTableImpl struct {
	lg esl.Logger
	mc app_msg_container.Container
	mx es_mutex.Mutex
	it Table
}

func (z proxyTableImpl) verifyMessages(ms []app_msg.Message) {
	for _, m := range ms {
		k := m.Key()
		if !z.mc.Exists(m) {
			z.lg.Debug("Message key not found", esl.String("key", k))
			qt_missingmsg.Record().NotFound(k)
		}
	}
}

func (z proxyTableImpl) Header(h ...app_msg.Message) {
	z.verifyMessages(h)
	z.mx.Do(func() {
		z.it.Header(h...)
	})
}

func (z proxyTableImpl) HeaderRaw(h ...string) {
	z.mx.Do(func() {
		z.it.HeaderRaw(h...)
	})
}

func (z proxyTableImpl) Row(m ...app_msg.Message) {
	z.verifyMessages(m)
	z.mx.Do(func() {
		z.it.Row(m...)
	})
}

func (z proxyTableImpl) RowRaw(m ...string) {
	z.mx.Do(func() {
		z.it.RowRaw(m...)
	})
}

func (z proxyTableImpl) Flush() {
	z.mx.Do(func() {
		z.it.Flush()
	})
}
