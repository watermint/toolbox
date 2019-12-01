package ut_filepath

import "testing"

func TestRel(t *testing.T) {
	// Success
	{
		if p, err := Rel("/a/b/c", "/a/b/c"); p != "." || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/c", "/a/b/c/"); p != "." || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/c/", "/a/b/c/"); p != "." || err != nil {
			t.Error(p, err)
		}

		if p, err := Rel("/a/b/c", "/a/b/c/d"); p != "d" || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/c/", "/a/b/c/d"); p != "d" || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/B/c/", "/a/b/c/d"); p != "d" || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("/東京都/港区/", "/東京都/港区/六本木"); p != "六本木" || err != nil {
			t.Error(p, err)
		}

		isWindows = true
		if p, err := Rel("C:\\a\\b\\c", "C:\\a\\b\\c"); p != "." || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("C:\\a\\b\\c", "C:\\a\\b\\c\\d"); p != "d" || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("\\\\10.1.1.1\\a\\b\\c", "\\\\10.1.1.1\\a\\b\\c\\d"); p != "d" || err != nil {
			t.Error(p, err)
		}
		if p, err := Rel("\\\\10.1.1.1\\東京都\\港区\\六本木", "\\\\10.1.1.1\\東京都\\港区\\六本木\\六本木交差点"); p != "六本木交差点" || err != nil {
			t.Error(p, err)
		}
		isWindows = false
	}

	// Error case
	{
		if p, err := Rel("/a/b/c/d", "/a/b/c"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("", "/a/b/c"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/c", ""); p != "" || err == nil {
			t.Error(p, err)
		}

		if p, err := Rel("/a/b/d", "/a/b/c/d"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/cd", "/a/b/c/d"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("/a/b/X", "/a/b/c/d"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("/x/y/z", "/a/b/c/d"); p != "" || err == nil {
			t.Error(p, err)
		}

		isWindows = true
		if p, err := Rel("C:\\a\\b\\d", "C:\\a\\b\\c\\d"); p != "" || err == nil {
			t.Error(p, err)
		}
		if p, err := Rel("\\\\10.1.1.1\\a\\b\\d", "\\\\10.1.1.1\\a\\b\\c\\d"); p != "" || err == nil {
			t.Error(p, err)
		}
		isWindows = false
	}

}