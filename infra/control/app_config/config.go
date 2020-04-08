package app_config

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
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
		path:     path,
		values:   make(map[string]interface{}),
		mutex:    sync.Mutex{},
		loadTime: time.Time{},
	}
}

type configImpl struct {
	path     string
	values   map[string]interface{}
	mutex    sync.Mutex
	loadTime time.Time
}

func (z *configImpl) load() (err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := app_root.Log()
	p := filepath.Join(z.path, ConfigFileName)

	s, err := os.Lstat(p)
	if err != nil {
		l.Debug("No file information; skip loading", zap.Error(err))
		return nil
	}

	if z.loadTime.After(s.ModTime()) {
		l.Debug("Skip loading")
		return nil
	}

	l.Debug("load config", zap.String("path", p))
	b, err := ioutil.ReadFile(p)
	if err != nil {
		l.Debug("Unable to read config", zap.Error(err))
		return err
	}
	if err := json.Unmarshal(b, &z.values); err != nil {
		l.Debug("unable to unmarshal", zap.Error(err))
		return err
	}
	z.loadTime = time.Now()
	return nil
}

func (z *configImpl) save() (err error) {
	z.mutex.Lock()
	defer z.mutex.Unlock()

	l := app_root.Log()
	p := filepath.Join(z.path, ConfigFileName)
	l.Debug("load config", zap.String("path", p))
	b, err := json.Marshal(z.values)
	if err != nil {
		l.Debug("Unable to marshal", zap.Error(err))
		return err
	}
	if err := ioutil.WriteFile(p, b, 0644); err != nil {
		l.Debug("Unable to write config", zap.Error(err))
		return err
	}
	return nil
}

func (z *configImpl) Get(key string) (v interface{}, err error) {
	if err := z.load(); err != nil {
		return nil, err
	}
	if v, ok := z.values[key]; ok {
		return v, nil
	} else {
		return nil, ErrorValueNotFound
	}
}

func (z *configImpl) Put(key string, v interface{}) (err error) {
	z.mutex.Lock()
	if z.values == nil {
		z.values = make(map[string]interface{})
	}
	z.values[key] = v
	z.mutex.Unlock()
	return z.save()
}

func (z *configImpl) List() (settings map[string]interface{}, err error) {
	if err := z.load(); err != nil {
		return nil, err
	}
	return z.values, nil
}
