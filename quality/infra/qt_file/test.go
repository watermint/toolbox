package qt_file

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func MakeDummyFile(name string) (path string, err error) {
	d, err := os.CreateTemp("", name)
	if err != nil {
		return "", err
	}
	_, err = d.Write([]byte(time.Now().Format(time.RFC3339)))
	if err != nil {
		os.Remove(d.Name())
		return "", err
	}
	d.Close()
	return d.Name(), nil
}

func TestWithTestFile(t *testing.T, name, content string, f func(path string)) {
	tf, err := MakeTestFile(name, content)
	if err != nil {
		t.Error(err)
		return
	}
	f(tf)
	_ = os.Remove(tf)
}

func MakeTestCsv(name string) (path string, err error) {
	return MakeTestFile(name, "alex@example.com,Alex\ndavid@example.com,David\nkevin@example.com,Kevin\n")
}

func MakeTestFile(name string, content string) (path string, err error) {
	d, err := os.CreateTemp("", name)
	if err != nil {
		return "", err
	}
	_, err = d.Write([]byte(content))
	if err != nil {
		_ = os.Remove(d.Name())
		return "", err
	}
	_ = d.Close()
	return d.Name(), nil
}

func TestWithTestFolder(t *testing.T, name string, withContent bool, f func(path string)) {
	path, err := MakeTestFolder(name, withContent)
	if err != nil {
		t.Error(err)
		return
	}
	f(path)
	if err = os.RemoveAll(path); err != nil {
		t.Error(err)
	}
}

func MakeTestFolder(name string, withContent bool) (path string, err error) {
	path, err = os.MkdirTemp("", name)
	if err != nil {
		return "", err
	}
	if withContent {
		err := os.WriteFile(filepath.Join(path, "test.dat"), []byte(time.Now().String()), 0644)
		if err != nil {
			os.RemoveAll(path)
			return "", err
		}
	}
	return path, nil
}
