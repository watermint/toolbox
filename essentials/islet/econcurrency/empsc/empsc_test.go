package empsc

import (
	"github.com/watermint/toolbox/essentials/islet/econcurrency/ethread"
	"testing"
)

type channelImplTester struct {
	t      *testing.T
	mainRn string
}

func (z *channelImplTester) Consumer(v interface{}) (err error) {
	consumerRn := ethread.CurrentRoutineName()
	if z.mainRn == consumerRn {
		z.t.Error(consumerRn, z.mainRn)
	}

	z.t.Log(v)
	return nil
}

func TestChannelImpl_Send(t *testing.T) {
	ct := &channelImplTester{
		t:      t,
		mainRn: ethread.CurrentRoutineName(),
	}
	ch := New(ct.Consumer)
	ch.Producer().Send("hello")
	ch.Producer().Send("world")
	if err := ch.Close(); err != nil {
		t.Error(err)
	}
	ch.Producer().Send("again")
}
