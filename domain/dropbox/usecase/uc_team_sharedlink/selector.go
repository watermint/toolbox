package uc_team_sharedlink

import (
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
)

type Selector interface {
	// Register the url for process.
	Register(url string) error

	// Mark the url as processed.
	Processed(url string) error

	// Check weather the url is for process.
	IsTarget(url string) (bool, error)

	// Report missing urls
	Done() error
}

const (
	selectStatusRegistered = "R"
	selectStatusProcessed  = "D"
)

type SelectorOnMissing func(url string)

func NewSelector(c app_control.Control, onMissing SelectorOnMissing) (Selector, error) {
	kv, err := c.NewKvs("select" + sc_random.MustGetSecureRandomString(6))
	if err != nil {
		return nil, err
	}
	return &selImpl{
		ctl:       c,
		kv:        kv,
		onMissing: onMissing,
	}, nil
}

type selImpl struct {
	ctl       app_control.Control
	kv        kv_storage.Storage // url -> state
	onMissing SelectorOnMissing
}

func (z *selImpl) IsTarget(url string) (bool, error) {
	shouldProcess := false
	kvErr := z.kv.View(func(kvs kv_kvs.Kvs) error {
		v, kvErr := kvs.GetString(url)
		if kvErr == kv_kvs.ErrorNotFound {
			return nil
		}
		if v == selectStatusRegistered {
			shouldProcess = true
			return nil
		}
		return kvErr
	})
	return shouldProcess, kvErr
}

func (z *selImpl) Done() error {
	return z.kv.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEach(func(url string, value []byte) error {
			s := string(value)
			if s == selectStatusRegistered {
				z.onMissing(url)
			}
			return nil
		})
	})
}

func (z *selImpl) Register(url string) error {
	return z.kv.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutString(url, selectStatusRegistered)
	})
}

func (z *selImpl) Processed(url string) error {
	return z.kv.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutString(url, selectStatusProcessed)
	})
}
