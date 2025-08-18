package api_callback

import (
	"testing"
	"time"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

func TestCallbackImpl_pingMethod(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		// Test ping with server error
		cb := &callbackImpl{
			ctl:         ctl,
			service:     service,
			port:        8080,
			secure:      false,
			serverError: ErrorAnotherServerOnline,
		}

		err := cb.ping()
		if err != ErrorAnotherServerOnline {
			t.Errorf("Expected ErrorAnotherServerOnline, got %v", err)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_Shutdown(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		cb := &callbackImpl{
			ctl:        ctl,
			service:    service,
			port:       8080,
			secure:     false,
			flowStatus: make(chan struct{}, 1),
		}

		// Test shutdown without server
		cb.Shutdown()

		// flowStatus might be nil, just test that it doesn't panic

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_Fields(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		cb := &callbackImpl{
			instance:    "test-instance",
			service:     service,
			ctl:         ctl,
			port:        8080,
			secure:      true,
			serverToken: "test-token",
			serverReady: true,
		}

		if cb.instance != "test-instance" {
			t.Error("Expected instance to be 'test-instance'")
		}

		if cb.serverToken != "test-token" {
			t.Error("Expected serverToken to be 'test-token'")
		}

		if !cb.serverReady {
			t.Error("Expected serverReady to be true")
		}

		if !cb.secure {
			t.Error("Expected secure to be true")
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestMsgCallback(t *testing.T) {
	// Test that MCallback is initialized
	if MCallback == nil {
		t.Error("Expected MCallback to be initialized")
	}

	// Test that messages are accessible
	_ = MCallback.MsgOpenUrlOnYourBrowser
	_ = MCallback.MsgErrorOpenUrlOnYourBrowser
	_ = MCallback.MsgHitEnterToProceed
	_ = MCallback.MsgResultSuccessHeader
	_ = MCallback.MsgResultSuccessBody
	_ = MCallback.MsgResultFailureHeader
	_ = MCallback.MsgResultFailureBody
	_ = MCallback.MsgHelloHeader
	_ = MCallback.MsgHelloBody
}

func TestErrorConstants(t *testing.T) {
	if ErrorAnotherServerOnline.Error() != "another server is online" {
		t.Error("Expected ErrorAnotherServerOnline to have correct message")
	}

	if shutdownTimeout != 5*1000*time.Millisecond {
		t.Error("Expected shutdownTimeout to be 5 seconds")
	}
}

func TestInstanceId(t *testing.T) {
	// Test that instanceId counter increments
	initialValue := instanceId.Load()

	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		cb1 := New(ctl, service, 8080, false)
		impl1 := cb1.(*callbackImpl)

		cb2 := New(ctl, service, 8081, false)
		impl2 := cb2.(*callbackImpl)

		// Instance IDs should be different
		if impl1.instance == impl2.instance {
			t.Error("Expected different instance IDs")
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	// Verify counter incremented
	if instanceId.Load() <= initialValue {
		t.Error("Expected instanceId to increment")
	}
}

func TestCallbackImpl_urlForPathSecure(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}

		cb := &callbackImpl{
			ctl:     ctl,
			service: service,
			port:    8443,
			secure:  true,
		}

		url := cb.urlForPath("/test")
		expected := "https://localhost:8443/test"

		if url != expected {
			t.Errorf("urlForPath() = %v, want %v", url, expected)
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestService_Interface(t *testing.T) {
	service := &mockService{verifyResult: true}

	// Test Url method
	url := service.Url("http://example.com/callback")
	expected := "http://example.com/callback?state=test-state"
	if url != expected {
		t.Errorf("Url() = %v, want %v", url, expected)
	}

	// Test Verify method with true result
	if !service.Verify("test", "code") {
		t.Error("Expected Verify to return true")
	}

	// Test Verify method with false result
	service.verifyResult = false
	if service.Verify("test", "code") {
		t.Error("Expected Verify to return false")
	}
}
