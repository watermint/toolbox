package qt_file

import (
	"io/ioutil"
	"os"
	"time"
)

func MakeDummyFile(name string) (path string, err error) {
	d, err := ioutil.TempFile("", name)
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

func MakeTestFile(name string, content string) (path string, err error) {
	d, err := ioutil.TempFile("", name)
	if err != nil {
		return "", err
	}
	_, err = d.Write([]byte(content))
	if err != nil {
		os.Remove(d.Name())
		return "", err
	}
	d.Close()
	return d.Name(), nil
}
