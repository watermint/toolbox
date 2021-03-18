package qt_msgusage

import (
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"sync"
	"testing"
	"time"
)

func TestMisImpl_Concurrent(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skipped()
		return
	}

	// #506
	wg := sync.WaitGroup{}
	dur := 5 * time.Second

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			start := time.Now()
			finish := start.Add(dur)

			for finish.After(time.Now()) {
				Record().Touch(sc_random.MustGetSecureRandomString(3))
				Record().NotFound(sc_random.MustGetSecureRandomString(3))
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
