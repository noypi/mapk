package mapk

import (
	"sync"
)

type _threadsafe struct {
	m IMap
	l sync.RWMutex
}

func MakeThreadSafe(m IMap) IMap {
	return &_threadsafe{m: m}
}

func (this *_threadsafe) Get(k interface{}) interface{} {
	this.l.RLock()
	defer this.l.RUnlock()

	return this.m.Get(k)
}

func (this *_threadsafe) Put(k, v interface{}) {
	this.l.Lock()
	defer this.l.Unlock()

	this.m.Put(k, v)
}

func (this *_threadsafe) Delete(k interface{}) {
	this.l.Lock()
	defer this.l.Unlock()

	this.m.Delete(k)
}

func (this *_threadsafe) Has(k interface{}) bool {
	this.l.RLock()
	defer this.l.RUnlock()

	return this.m.Has(k)
}

func (this *_threadsafe) EachFrom(kprefix interface{}, cb func(k, v interface{}) bool) {
	this.l.RLock()
	defer this.l.RUnlock()

	this.m.EachFrom(kprefix, cb)
}

func (this *_threadsafe) Each(cb func(k, v interface{}) bool) {
	this.l.RLock()
	defer this.l.RUnlock()

	this.m.Each(cb)
}

func (this *_threadsafe) Len() int {
	this.l.RLock()
	defer this.l.RUnlock()

	return this.m.Len()
}
