package api_callback

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/essentials/runtime/es_open"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
)

// Mock Service implementation
type mockService struct {
	verifyResult bool
}

func (m *mockService) Url(redirectUrl string) string {
	return redirectUrl + "?state=test-state"
}

func (m *mockService) Verify(state, code string) bool {
	return m.verifyResult
}

func TestNew(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		
		cb := New(ctl, service, 8080, false)
		
		if cb == nil {
			t.Fatal("Expected callback instance, got nil")
		}
		
		impl, ok := cb.(*callbackImpl)
		if !ok {
			t.Fatal("Expected callbackImpl type")
		}
		
		if impl.port != 8080 {
			t.Errorf("Expected port 8080, got %d", impl.port)
		}
		
		if impl.secure != false {
			t.Error("Expected secure to be false")
		}
		
		if impl.service != service {
			t.Error("Service not set correctly")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewWithOpener(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		opener := es_open.NewTestDummy()
		
		cb := NewWithOpener(ctl, service, 8080, false, opener)
		
		if cb == nil {
			t.Fatal("Expected callback instance, got nil")
		}
		
		impl, ok := cb.(*callbackImpl)
		if !ok {
			t.Fatal("Expected callbackImpl type")
		}
		
		if impl.opener != opener {
			t.Error("Opener not set correctly")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_urlForPath(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		
		tests := []struct {
			name   string
			secure bool
			port   int
			path   string
			want   string
		}{
			{
				name:   "http",
				secure: false,
				port:   8080,
				path:   "/test",
				want:   "http://localhost:8080/test",
			},
			{
				name:   "https",
				secure: true,
				port:   8443,
				path:   "/test",
				want:   "https://localhost:8443/test",
			},
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

func TestCallbackImpl_Url(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		
		cb := &callbackImpl{
			ctl:     ctl,
			service: service,
			port:    8080,
			secure:  false,
		}
		
		url := cb.Url()
		expected := "http://localhost:8080/connect/auth"
		
		if url != expected {
			t.Errorf("Url() = %v, want %v", url, expected)
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_Ping(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		
		cb := &callbackImpl{
			ctl:         ctl,
			service:     service,
			port:        8080,
			secure:      false,
			serverToken: "test-token",
		}
		
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		
		cb.Ping(c)
		
		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}
		
		// Check response contains expected fields
		body := w.Body.String()
		if !strings.Contains(body, "test-token") {
			t.Error("Response should contain server token")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_Connect(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		tests := []struct {
			name         string
			verifyResult bool
			expectedPath string
		}{
			{
				name:         "success",
				verifyResult: true,
				expectedPath: PathSuccess,
			},
			{
				name:         "failure",
				verifyResult: false,
				expectedPath: PathFailure,
			},
		}
		
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				service := &mockService{verifyResult: tt.verifyResult}
				cb := &callbackImpl{
					ctl:     ctl,
					service: service,
					port:    8080,
					secure:  false,
				}
				
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest("GET", "/connect/auth?state=test&code=test", nil)
				
				cb.Connect(c)
				
				if w.Code != http.StatusTemporaryRedirect {
					t.Errorf("Expected status %d, got %d", http.StatusTemporaryRedirect, w.Code)
				}
				
				location := w.Header().Get("Location")
				if location != tt.expectedPath {
					t.Errorf("Expected redirect to %s, got %s", tt.expectedPath, location)
				}
			})
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestCallbackImpl_WaitServerReady(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		service := &mockService{}
		
		cb := &callbackImpl{
			ctl:         ctl,
			service:     service,
			port:        8080,
			secure:      false,
			serverReady: false,
		}
		
		// Test when server becomes ready
		go func() {
			time.Sleep(50 * time.Millisecond)
			cb.serverReady = true
		}()
		
		if !cb.WaitServerReady() {
			t.Error("Expected server to be ready")
		}
		
		// Test when server has error
		cb2 := &callbackImpl{
			ctl:         ctl,
			service:     service,
			port:        8080,
			secure:      false,
			serverReady: false,
			serverError: ErrorAnotherServerOnline,
		}
		
		if cb2.WaitServerReady() {
			t.Error("Expected server to not be ready due to error")
		}
		
		return nil
	})
	
	if err != nil {
		t.Fatal(err)
	}
}

func TestServerStatus(t *testing.T) {
	status := ServerStatus{
		Name:    "test-server",
		Version: "1.0.0",
		Token:   "test-token",
	}
	
	if status.Name != "test-server" {
		t.Errorf("Expected name test-server, got %s", status.Name)
	}
	if status.Version != "1.0.0" {
		t.Errorf("Expected version 1.0.0, got %s", status.Version)
	}
	if status.Token != "test-token" {
		t.Errorf("Expected token test-token, got %s", status.Token)
	}
}

func TestConstants(t *testing.T) {
	// Test that constants have expected values
	if PathPing != "/ping" {
		t.Errorf("PathPing = %s, want /ping", PathPing)
	}
	if PathConnect != "/connect/auth" {
		t.Errorf("PathConnect = %s, want /connect/auth", PathConnect)
	}
	if PathSuccess != "/success" {
		t.Errorf("PathSuccess = %s, want /success", PathSuccess)
	}
	if PathFailure != "/failure" {
		t.Errorf("PathFailure = %s, want /failure", PathFailure)
	}
	if PathHello != "/hello" {
		t.Errorf("PathHello = %s, want /hello", PathHello)
	}
	if DataUriImagePng != "data:image/png;base64," {
		t.Errorf("DataUriImagePng = %s, want data:image/png;base64,", DataUriImagePng)
	}
}