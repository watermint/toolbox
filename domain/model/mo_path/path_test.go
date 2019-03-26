package mo_path

import "testing"

func TestPathImpl(t *testing.T) {
	{
		p := pathImpl{path: ""}
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
		p := pathImpl{path: "/world.txt"}
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
		p := pathImpl{path: "/hello/world.txt"}
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
		p := pathImpl{path: "id:abc123xyz"}
		if pi := p.Path(); pi != "id:abc123xyz" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "abc123xyz" || !e {
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
		p := pathImpl{path: "id:abc123xyz/world.txt"}
		if pi := p.Path(); pi != "id:abc123xyz/world.txt" {
			t.Error("invalid")
		}
		if pi, e := p.Id(); pi != "abc123xyz" || !e {
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
		p := pathImpl{path: "id:abc123xyz/hello/world.txt"}
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
		p := pathImpl{path: "ns:123456"}
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
		p := pathImpl{path: "ns:123456/hello/world.txt"}
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

	if p := NewPath("/hello//world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("//hello//world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("id:abc123xyz//hello/world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}

	if p := NewPath(""); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewPath("/"); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewPath("/hello/world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("id:abc123xyz"); p.Path() != "id:abc123xyz" {
		t.Error("invalid")
	}
	if p := NewPath("id:abc123xyz/hello/world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("ns:123456"); p.Path() != "ns:123456" {
		t.Error("invalid")
	}
	if p := NewPath("ns:123456/hello/world.txt"); p.Path() != "ns:123456/hello/world.txt" {
		t.Error("invalid")
	}

	// Windows paths

	if p := NewPath("\\"); p.Path() != "" {
		t.Error("invalid")
	}
	if p := NewPath("\\hello\\world.txt"); p.Path() != "/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("id:abc123xyz"); p.Path() != "id:abc123xyz" {
		t.Error("invalid")
	}
	if p := NewPath("id:abc123xyz\\hello\\world.txt"); p.Path() != "id:abc123xyz/hello/world.txt" {
		t.Error("invalid")
	}
	if p := NewPath("ns:123456"); p.Path() != "ns:123456" {
		t.Error("invalid")
	}
	if p := NewPath("ns:123456\\hello\\world.txt"); p.Path() != "ns:123456/hello/world.txt" {
		t.Error("invalid")
	}

}
