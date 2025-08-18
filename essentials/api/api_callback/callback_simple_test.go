package api_callback

import (
	"github.com/watermint/toolbox/essentials/runtime/es_open"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

func TestCallbackImpl_urlForPath_variations(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		// Test different port numbers
		tests := []struct {
			name   string
			port   int
			secure bool
			path   string
			want   string
		}{
			{"http_8080", 8080, false, "/test", "http://localhost:8080/test"},
			{"http_3000", 3000, false, "/api", "http://localhost:3000/api"},
			{"https_8443", 8443, true, "/auth", "https://localhost:8443/auth"},
			{"https_443", 443, true, "/", "https://localhost:443/"},
			{"http_root", 80, false, "", "http://localhost:80"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				cb := &callbackImpl{
					ctl:     ctl,
					service: service,
					port:    tt.port,
					secure:  tt.secure,
				}

				got := cb.urlForPath(tt.path)
				if got != tt.want {
					t.Errorf("urlForPath() = %v, want %v", got, tt.want)
				}
			})
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_openUrl(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		cb := &callbackImpl{
			ctl:     ctl,
			service: service,
			port:    8080,
			secure:  false,
			opener:  es_open.NewTestDummy(),
		}

		// Just test that it doesn't panic
		cb.openUrl("http://example.com/auth")

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
