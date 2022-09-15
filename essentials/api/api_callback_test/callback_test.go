package api_callback_test

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_callback"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/runtime/es_open"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"net/http"
	"net/url"
	"sync"
	"testing"
)

var (
	ErrorAuthFailure = errors.New("auth failure")
)

func NewMockService(ctl app_control.Control) *MockService {
	return &MockService{
		ctl:   ctl,
		state: sc_random.MustGetSecureRandomString(8),
		code:  sc_random.MustGetSecureRandomString(8),
	}
}

type MockService struct {
	ctl      app_control.Control
	state    string
	code     string
	redirect string
}

func (z *MockService) Url(redirectUrl string) string {
	z.redirect = redirectUrl
	return "http://localhost/mock?state" + z.state + "&redirect_url" + url.QueryEscape(redirectUrl)
}

func (z *MockService) Verify(state, code string) bool {
	l := z.ctl.Log()
	if state != z.state {
		l.Debug("Wrong state", esl.String("given", state), esl.String("expected", z.state))
		return false
	}
	if code != z.code {
		l.Debug("Wrong code", esl.String("given", code), esl.String("expected", z.code))
		return false
	}
	l.Debug("verification succeed")
	return true
}

func (z *MockService) Ping(url string) error {
	l := z.ctl.Log()
	hc := http.Client{}
	l.Debug("Ping valid request", esl.String("url", url))

	res, err := hc.Get(url)
	if err != nil {
		l.Debug("Error from ping", esl.Error(err))
		return err
	}
	l.Debug("Response", esl.Int("code", res.StatusCode))
	if res.StatusCode == http.StatusOK {
		l.Debug("Return success")
		return nil
	}
	return ErrorAuthFailure
}

func (z *MockService) PingValid() error {
	return z.Ping(z.redirect + "?code=" + z.code + "&state=" + z.state)
}

func (z *MockService) PingInvalid() error {
	return z.Ping(z.redirect + "?code=XXX&state=YYY")
}

func TestCallbackImpl_SuccessScenario(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ms := NewMockService(ctl)
		cb := api_callback.NewWithOpener(ctl, ms, 7800, es_open.NewTestDummy())
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			if err := cb.Flow(); err != nil {
				t.Error(err)
			}
			wg.Done()
		}()
		if !cb.WaitServerReady() {
			t.Error("the server have trouble")
			cb.Shutdown()
			return
		}

		if err := ms.PingValid(); err != nil {
			t.Error(err)
			cb.Shutdown()
			return
		}
		wg.Wait()
	})
}

func TestCallbackImpl_FailureInvalidCode(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ms := NewMockService(ctl)
		cb := api_callback.NewWithOpener(ctl, ms, 7800, es_open.NewTestDummy())
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			if err := cb.Flow(); err != nil {
				t.Error(err)
			}
			wg.Done()
		}()
		if !cb.WaitServerReady() {
			t.Error("the server have trouble")
			cb.Shutdown()
			return
		}

		if err := ms.PingInvalid(); err != ErrorAuthFailure {
			t.Error(err)
			cb.Shutdown()
			return
		}
		wg.Wait()
	})
}

func TestCallbackImpl_FailureCantStart(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ms := NewMockService(ctl)
		cb1 := api_callback.NewWithOpener(ctl, ms, 7800, es_open.NewTestDummy())
		cb2 := api_callback.NewWithOpener(ctl, ms, 7800, es_open.NewTestDummy())
		go func() {
			if err := cb1.Flow(); err != nil {
				t.Error(err)
			}
		}()
		defer cb1.Shutdown()
		if !cb1.WaitServerReady() {
			t.Error("server looks like have a trouble")
			return
		}
		if err := cb2.Flow(); err == nil {
			t.Error("invalid")
		}
	})
}
