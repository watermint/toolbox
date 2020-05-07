package app_config

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	ConfigFileName = "config.json"
)

var (
	ErrorValueNotFound = errors.New("value not found")
)

type Config interface {
	Get(key string) (v interface{}, err error)
	Put(key string, v interface{}) (err error)
	List() (settings map[string]interface{}, err error)
}

func NewConfig(path string) Config {
	return &configImpl{
		path: path,
	}
}

type configImpl struct {
	path string
}

func (z *configImpl) load() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	l := es_log.Default()
	p := filepath.Join(z.path, ConfigFileName)

	_, err = os.Lstat(p)
	if err != nil {
		l.Debug("No file information; skip loading", es_log.Error(err))
		return values, nil
	}

	l.Debug("load config", es_log.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("Unable to read config", es_log.Error(err))
		return
	}
	if err := json.Unmarshal(b, &values); err != nil {
		l.Debug("unable to unmarshal", es_log.Error(err))
		return values, err
	}
	return
}

func (z *configImpl) save(key string, v interface{}) (err error) {
	l := es_log.Default()
	p := filepath.Join(z.path, ConfigFileName)
	l.Debug("load config", es_log.String("path", p))
	values, err := z.load()
	if err != nil {
		return err
	}
	values[key] = v

	b, err := json.Marshal(values)
	if err != nil {
		l.Debug("Unable to marshal", es_log.Error(err))
		return err
	}
	if err := ioutil.WriteFile(p, b, 0644); err != nil {
		l.Debug("Unable to write config", es_log.Error(err))
		return err
	}
	return nil
}

func (z *configImpl) Get(key string) (v interface{}, err error) {
	if values, err := z.load(); err != nil {
		return nil, err
	} else if v, ok := values[key]; ok {
		return v, nil
	} else {
		return nil, ErrorValueNotFound
	}
}

func (z *configImpl) Put(key string, v interface{}) (err error) {
	return z.save(key, v)
}

func (z *configImpl) List() (settings map[string]interface{}, err error) {
	return z.load()
}
