package es_env

import (
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"os"
	"testing"
)

func TestIsEnabled(t *testing.T) {
	testEnv := app_definitions.EnvNameDebugVerbose + "_OF_TEST_IS_ENABLED"
	os.Setenv(testEnv, "true")
	if e := IsEnabled(testEnv); !e {
		t.Error(e)
	}
	os.Setenv(testEnv, "false")
	if e := IsEnabled(testEnv); e {
		t.Error(e)
	}
	os.Setenv(testEnv, "unknown")
	if e := IsEnabled(testEnv); e {
		t.Error(e)
	}
}
