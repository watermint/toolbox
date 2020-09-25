package nw_proxy

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"testing"
)

func TestSetup(t *testing.T) {
	if x := Setup("https://api.dropboxapi.com", Direct, esl.Default()); !x {
		t.Error(x)
	}

	// should fallback to direct proxy setting when it's not reachable to the proxy.
	if x := Setup("https://api.dropboxapi.com", "no_existent_proxy:8080", esl.Default()); !x {
		t.Error(x)
	}
}
