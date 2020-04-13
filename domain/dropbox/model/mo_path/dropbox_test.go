package mo_path

import "testing"

func TestPathImpl(t *testing.T) {
	{
		p := NewDropboxPath("")
		if pi := p.Path(); pi != "" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "" || e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/" {
			t.Error("invalid")
		}
	}

	{
		p := NewDropboxPath("/world.txt")
		if pi := p.Path(); pi != "/world.txt" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "" || e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/world.txt" {
			t.Error("invalid")
		}
	}

	{
		p := NewDropboxPath("/hello/world.txt")
		if pi := p.Path(); pi != "/hello/world.txt" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "" || e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/hello/world.txt" {
			t.Error("invalid")
		}
	}

	{
		p := NewDropboxPath("id:abc123xyz")
		if pi := p.Path(); pi != "id:abc123xyz" {
			t.Error("invalid", "["+pi+"]")
		}
		if pi, e := p.Id(); pi != "abc123xyz" || !e {
			t.Error("invalid", "["+pi+"]")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid", "["+pi+"]")
		}
		if pi := p.LogicalPath(); pi != "/" {
			t.Error("invalid", "["+pi+"]")
		}
	}

	{
		p := NewDropboxPath("id:abc123xyz/world.txt")
		if pi := p.Path(); pi != "id:abc123xyz/world.txt" {
			t.Error("invalid", "["+pi+"]")
		}
		if pi, e := p.Id(); pi != "abc123xyz" || !e {
			t.Error("invalid", "["+pi+"]")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid", "["+pi+"]")
		}
		if pi := p.LogicalPath(); pi != "/world.txt" {
			t.Error("invalid", "["+pi+"]")
		}
	}

	{
		p := NewDropboxPath("id:abc123xyz/hello/world.txt")
		if pi := p.Path(); pi != "id:abc123xyz/hello/world.txt" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "abc123xyz" || !e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "" || e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/hello/world.txt" {
			t.Error("invalid")
		}
	}

	{
		p := NewDropboxPath("ns:123456")
		if pi := p.Path(); pi != "ns:123456" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "" || e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "123456" || !e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/" {
			t.Error("invalid")
		}
	}

	{
		p := NewDropboxPath("ns:123456/hello/world.txt")
		if pi := p.Path(); pi != "ns:123456/hello/world.txt" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "" || e {
			t.Error("invalid")
		}
		if pi, e := p.Namespace(); pi != "123456" || !e {
			t.Error("invalid")
		}
		if pi := p.LogicalPath(); pi != "/hello/world.txt" {
			t.Error("invalid")
		}
	}
}

func TestNewPath(t *testing.T) {
	// multiple separators

	if p := NewDropboxPath("/hello//world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("//hello//world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("id:abc123xyz//hello/world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}

	if p := NewDropboxPath(""); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("/"); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("/hello/world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("id:abc123xyz"); p.Path() != "id:abc123xyz" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("id:abc123xyz/hello/world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("ns:123456"); p.Path() != "ns:123456" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("ns:123456/hello/world.txt"); p.Path() != "ns:123456/hello/world.txt" {
		t.Error("invalid")
	}

	// Windows paths

	if p := NewDropboxPath("\\"); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("\\hello\\world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("id:abc123xyz"); p.Path() != "id:abc123xyz" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("id:abc123xyz\\hello\\world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("ns:123456"); p.Path() != "ns:123456" {
		t.Error("invalid")
	}
	if p := NewDropboxPath("ns:123456\\hello\\world.txt"); p.Path() != "ns:123456/hello/world.txt" {
		t.Error("invalid")
	}

}
