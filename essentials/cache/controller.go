package cache

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	DefaultCacheTimeout = 24 * time.Hour
	CurrentVersion      = 1
	InfoPrefix          = "cache_info"
)

type Controller interface {
	// Startup cache controller. Broken or expired caches will be evicted on startup.
	Startup() error

	// Shutdown caches
	Shutdown()

	// Get or create new cache
	Cache(namespace, name string) (cache Cache)
}

type Info struct {
	Version   int       `json:"version"`
	Namespace string    `json:"namespace"`
	Name      string    `json:"name"`
	Expire    time.Time `json:"expire"`
	Last      time.Time `json:"last"`
}

func (z Info) DatabaseName() string {
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	data := fmt.Sprintf("%d-%s-%s", z.Version, z.Namespace, z.Name)
	hash := sha256.Sum256([]byte(data))
	return encoder.EncodeToString(hash[:])
}

func New(basePath string, logger esl.Logger) Controller {
	return &ctlImpl{
		basePath:   basePath,
		logger:     logger,
		cacheMutex: sync.Mutex{},
		caches:     make(map[string]Lifecycle),
		openCaches: make(map[string]Lifecycle),
		factory:    kv_storage_impl.NewFactory(basePath, logger),
	}
}

type ctlImpl struct {
	basePath   string
	logger     esl.Logger
	factory    kv_storage.Factory
	cacheMutex sync.Mutex
	caches     map[string]Lifecycle
	openCaches map[string]Lifecycle
}

func (z *ctlImpl) log() esl.Logger {
	return z.logger.With(esl.String("basePath", z.basePath))
}

func (z *ctlImpl) cacheKey(namespace, name string) string {
	return namespace + "-" + name
}

func (z *ctlImpl) noLockStartCache(path string) {
	l := z.log().With(esl.String("path", path))
	l.Debug("Check cache")

	infoData, err := ioutil.ReadFile(path)
	if err != nil {
		l.Debug("Unable to read cache file", esl.Error(err))
		return
	}

	info := Info{}
	if err := json.Unmarshal(infoData, &info); err != nil {
		l.Debug("Unable to unmarshal cache info", esl.Error(err))
		return
	}

	cache := newCache(info, path, z.basePath, z.factory, z.log(), DefaultCacheTimeout)
	cache.EvictIfRequired()
	z.caches[z.cacheKey(info.Namespace, info.Name)] = cache
}

func (z *ctlImpl) Startup() error {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	l := z.log()
	l.Debug("Startup cache controller")
	entries, err := ioutil.ReadDir(z.basePath)
	if err != nil {
		l.Debug("Unable to read the cache folder", esl.Error(err))
		return err
	}

	prefix := strings.ToLower(InfoPrefix)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if strings.ToLower(entry.Name()) != prefix {
			continue
		}
		z.noLockStartCache(filepath.Join(z.basePath, entry.Name()))
	}

	return nil
}

func (z *ctlImpl) Shutdown() {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	for _, cache := range z.openCaches {
		cache.Close()
	}
	z.factory.Close()
}

func (z *ctlImpl) Cache(namespace, name string) (cache Cache) {
	z.cacheMutex.Lock()
	defer z.cacheMutex.Unlock()

	l := z.log().With(esl.String("namespace", namespace), esl.String("name", name))
	l.Debug("Get cache")

	key := z.cacheKey(namespace, name)

	if cache, ok := z.caches[key]; ok {
		err := cache.Open()
		l.Debug("Open cache", esl.Error(err))
		return cache
	}

	info := Info{
		Version:   CurrentVersion,
		Namespace: namespace,
		Name:      name,
		Expire:    time.Now().Add(DefaultCacheTimeout),
		Last:      time.Now(),
	}

	cacheInfoPath := filepath.Join(z.basePath, InfoPrefix+info.DatabaseName()+".json")
	lifecycle := newCache(info, cacheInfoPath, z.basePath, z.factory, z.log(), DefaultCacheTimeout)
	err := lifecycle.Open()
	l.Debug("Open cache", esl.Error(err))
	if err != nil {
		l.Debug("Remark as cache opened")
		z.openCaches[key] = lifecycle
	}

	return lifecycle
}
