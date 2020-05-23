package mo_content

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"testing"
)

func TestNewContentsFile(t *testing.T) {
	src := `{
  "type": "file",
  "encoding": "base64",
  "size": 5362,
  "name": "README.md",
  "path": "README.md",
  "content": "encoded content ...",
  "sha": "3d21ec53a331a6f037a91c368710b99387d012c1",
  "url": "https://api.github.com/repos/octokit/octokit.rb/contents/README.md",
  "git_url": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/3d21ec53a331a6f037a91c368710b99387d012c1",
  "html_url": "https://github.com/octokit/octokit.rb/blob/master/README.md",
  "download_url": "https://raw.githubusercontent.com/octokit/octokit.rb/master/README.md",
  "_links": {
    "git": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/3d21ec53a331a6f037a91c368710b99387d012c1",
    "self": "https://api.github.com/repos/octokit/octokit.rb/contents/README.md",
    "html": "https://github.com/octokit/octokit.rb/blob/master/README.md"
  }
}`

	j, err := es_json.ParseString(src)
	if err != nil {
		t.Error(err)
		return
	}
	cts, err := NewContents(j)
	if err != nil {
		t.Error(err)
		return
	}

	if c, found := cts.File(); !found {
		t.Error(c, found)
	} else {
		if c.Type != "file" {
			t.Error(c.Type)
		}
		if c.Size != 5362 {
			t.Error(c.Size)
		}
		if c.Sha != "3d21ec53a331a6f037a91c368710b99387d012c1" {
			t.Error(c.Sha)
		}
		if c.Name != "README.md" {
			t.Error(c.Name)
		}
	}

	if c, found := cts.Dir(); found {
		t.Error(c, found)
	}
	if c, found := cts.Submodule(); found {
		t.Error(c, found)
	}
	if c, found := cts.Symlink(); found {
		t.Error(c, found)
	}
}

func TestNewContentsDir(t *testing.T) {
	src := `[
  {
    "type": "file",
    "size": 625,
    "name": "octokit.rb",
    "path": "lib/octokit.rb",
    "sha": "fff6fe3a23bf1c8ea0692b4a883af99bee26fd3b",
    "url": "https://api.github.com/repos/octokit/octokit.rb/contents/lib/octokit.rb",
    "git_url": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/fff6fe3a23bf1c8ea0692b4a883af99bee26fd3b",
    "html_url": "https://github.com/octokit/octokit.rb/blob/master/lib/octokit.rb",
    "download_url": "https://raw.githubusercontent.com/octokit/octokit.rb/master/lib/octokit.rb",
    "_links": {
      "self": "https://api.github.com/repos/octokit/octokit.rb/contents/lib/octokit.rb",
      "git": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/fff6fe3a23bf1c8ea0692b4a883af99bee26fd3b",
      "html": "https://github.com/octokit/octokit.rb/blob/master/lib/octokit.rb"
    }
  },
  {
    "type": "dir",
    "size": 0,
    "name": "octokit",
    "path": "lib/octokit",
    "sha": "a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d",
    "url": "https://api.github.com/repos/octokit/octokit.rb/contents/lib/octokit",
    "git_url": "https://api.github.com/repos/octokit/octokit.rb/git/trees/a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d",
    "html_url": "https://github.com/octokit/octokit.rb/tree/master/lib/octokit",
    "download_url": null,
    "_links": {
      "self": "https://api.github.com/repos/octokit/octokit.rb/contents/lib/octokit",
      "git": "https://api.github.com/repos/octokit/octokit.rb/git/trees/a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d",
      "html": "https://github.com/octokit/octokit.rb/tree/master/lib/octokit"
    }
  }
]`

	j, err := es_json.ParseString(src)
	if err != nil {
		t.Error(err)
		return
	}
	cts, err := NewContents(j)
	if err != nil {
		t.Error(err)
		return
	}

	if c, found := cts.File(); found {
		t.Error(c, found)
	}
	if c, found := cts.Dir(); !found {
		t.Error(c, found)
	} else {
		if len(c) != 2 {
			t.Error(c)
		}
		if c[0].Name != "octokit.rb" {
			t.Error(c[0])
		}
		if c[1].Path != "lib/octokit" {
			t.Error(c[1])
		}
	}
	if c, found := cts.Submodule(); found {
		t.Error(c, found)
	}
	if c, found := cts.Symlink(); found {
		t.Error(c, found)
	}
}

func TestNewContentsSymlink(t *testing.T) {
	src := `{
  "type": "symlink",
  "target": "/path/to/symlink/target",
  "size": 23,
  "name": "some-symlink",
  "path": "bin/some-symlink",
  "sha": "452a98979c88e093d682cab404a3ec82babebb48",
  "url": "https://api.github.com/repos/octokit/octokit.rb/contents/bin/some-symlink",
  "git_url": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/452a98979c88e093d682cab404a3ec82babebb48",
  "html_url": "https://github.com/octokit/octokit.rb/blob/master/bin/some-symlink",
  "download_url": "https://raw.githubusercontent.com/octokit/octokit.rb/master/bin/some-symlink",
  "_links": {
    "git": "https://api.github.com/repos/octokit/octokit.rb/git/blobs/452a98979c88e093d682cab404a3ec82babebb48",
    "self": "https://api.github.com/repos/octokit/octokit.rb/contents/bin/some-symlink",
    "html": "https://github.com/octokit/octokit.rb/blob/master/bin/some-symlink"
  }
}`

	j, err := es_json.ParseString(src)
	if err != nil {
		t.Error(err)
		return
	}
	cts, err := NewContents(j)
	if err != nil {
		t.Error(err)
		return
	}

	if c, found := cts.File(); found {
		t.Error(c, found)
	}
	if c, found := cts.Dir(); found {
		t.Error(c, found)
	}
	if c, found := cts.Submodule(); found {
		t.Error(c, found)
	}
	if c, found := cts.Symlink(); !found {
		t.Error(c, found)
	} else {
		if c.Type != "symlink" {
			t.Error(c.Type)
		}
		if c.Target != "/path/to/symlink/target" {
			t.Error(c.Target)
		}
		if c.Size != 23 {
			t.Error(c.Size)
		}
	}
}

func TestNewContentsSubmodule(t *testing.T) {
	src := `{
  "type": "submodule",
  "submodule_git_url": "git://github.com/jquery/qunit.git",
  "size": 0,
  "name": "qunit",
  "path": "test/qunit",
  "sha": "6ca3721222109997540bd6d9ccd396902e0ad2f9",
  "url": "https://api.github.com/repos/jquery/jquery/contents/test/qunit?ref=master",
  "git_url": "https://api.github.com/repos/jquery/qunit/git/trees/6ca3721222109997540bd6d9ccd396902e0ad2f9",
  "html_url": "https://github.com/jquery/qunit/tree/6ca3721222109997540bd6d9ccd396902e0ad2f9",
  "download_url": null,
  "_links": {
    "git": "https://api.github.com/repos/jquery/qunit/git/trees/6ca3721222109997540bd6d9ccd396902e0ad2f9",
    "self": "https://api.github.com/repos/jquery/jquery/contents/test/qunit?ref=master",
    "html": "https://github.com/jquery/qunit/tree/6ca3721222109997540bd6d9ccd396902e0ad2f9"
  }
}`

	j, err := es_json.ParseString(src)
	if err != nil {
		t.Error(err)
		return
	}
	cts, err := NewContents(j)
	if err != nil {
		t.Error(err)
		return
	}

	if c, found := cts.File(); found {
		t.Error(c, found)
	}
	if c, found := cts.Dir(); found {
		t.Error(c, found)
	}
	if c, found := cts.Submodule(); !found {
		t.Error(c, found)
	} else {
		if c.Type != "submodule" {
			t.Error(c.Type)
		}
		if c.Size != 0 {
			t.Error(c.Size)
		}
		if c.Name != "qunit" {
			t.Error(c.Name)
		}
	}
	if c, found := cts.Symlink(); found {
		t.Error(c, found)
	}
}
