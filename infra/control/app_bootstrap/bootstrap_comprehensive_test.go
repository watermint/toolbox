package app_bootstrap

import (
	"testing"
)

func TestBootstrap_PublicAPI(t *testing.T) {
	// Test that Bootstrap interface is implemented
	bs := NewBootstrap()
	if bs == nil {
		t.Fatal("NewBootstrap should not return nil")
	}
	
	// Verify it returns the correct implementation
	if _, ok := bs.(*bsImpl); !ok {
		t.Fatal("NewBootstrap should return *bsImpl")
	}
}

func TestMRun_GlobalVariable(t *testing.T) {
	// Test that MRun is initialized
	if MRun == nil {
		t.Fatal("MRun should be initialized")
	}
	
	// Test some key fields exist
	if MRun.ErrorInvalidArgument == nil {
		t.Error("ErrorInvalidArgument should be initialized")
	}
	if MRun.ErrorTooManyArguments == nil {
		t.Error("ErrorTooManyArguments should be initialized")
	}
	if MRun.ErrorInterrupted == nil {
		t.Error("ErrorInterrupted should be initialized")
	}
	if MRun.ErrorPanic == nil {
		t.Error("ErrorPanic should be initialized")
	}
	if MRun.ErrorRecipeFailed == nil {
		t.Error("ErrorRecipeFailed should be initialized")
	}
}




func TestBootstrapInterface(t *testing.T) {
	// Verify Bootstrap interface is properly implemented
	var _ Bootstrap = &bsImpl{}
	
	// Test that NewBootstrap returns correct type
	bs := NewBootstrap()
	if _, ok := bs.(*bsImpl); !ok {
		t.Error("NewBootstrap should return *bsImpl")
	}
}