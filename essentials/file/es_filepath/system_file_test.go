package es_filepath

import "testing"

func TestIsSystemFile(t *testing.T) {
	sysFiles := []string{
		"/tmp/.DS_Store",
		"/Users/toolbox/Icon\r",
		"/Users/toolbox/Documents/thumbs.db",
		"/Users/toolbox/Dropbox/.dropbox",
		"/Users/toolbox/.dropbox",
	}
	otherFiles := []string{
		"/tmp/a.txt",
		"/tmp/.DS_StoreStore",
		"/tmp/thumbnails.db",
		"/tmp/dropbox",
	}

	for _, f := range sysFiles {
		if x := IsSystemFile(f); !x {
			t.Error(f, x)
		}
	}
	for _, f := range otherFiles {
		if x := IsSystemFile(f); x {
			t.Error(f, x)
		}
	}
}
