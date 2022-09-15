package mo_desktop

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestDesktop(t *testing.T) {
	j := `{
        "path": "/home/toolbox/Dropbox",
        "host": 123456789,
        "is_team": false,
        "subscription_type": "Basic"
    }`

	entry := &Desktop{}
	if err := api_parser.ParseModelString(entry, j); err != nil {
		t.Error(err)
	}
	if entry.Path != "/home/toolbox/Dropbox" {
		t.Error("path")
	}
}
