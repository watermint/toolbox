package es_mutex

import "sync"

type Mutex interface {
	Do(f func())
}

func New() Mutex {
	return &mutexImpl{}
}

type mutexImpl struct {
	m sync.Mutex
}

func (z *mutexImpl) Do(f func()) {
	z.m.Lock()
	defer z.m.Unlock()
	f()
}
