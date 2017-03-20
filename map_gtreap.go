package mapk

import (
	"math/rand"
	"sync/atomic"

	"github.com/steveyen/gtreap"
)

func MapGTreap(comp func(a, b interface{}) int) *_GtreapMap {
	o := new(_GtreapMap)
	o.m = gtreap.NewTreap(func(a, b interface{}) int {
		return comp(a.(*_kv).k, b.(*_kv).k)
	})
	return o
}

func (this _GtreapMap) Get(k interface{}) interface{} {
	o := this.m.Get(&_kv{k: k})
	if gtreap.Item(nil) == o {
		return nil
	}
	return o.(*_kv).v
}

func (this *_GtreapMap) Delete(k interface{}) {
	this.m = this.m.Delete(&_kv{k: k})
	atomic.AddInt64(&this.count, -1)
}

func (this *_GtreapMap) Put(k, v interface{}) {
	ok := &_kv{k: k}
	o := this.m.Get(ok)
	if gtreap.Item(nil) == o {
		atomic.AddInt64(&this.count, 1)
		ok.v = v
		this.m = this.m.Upsert(ok, rand.Int())
	} else {
		o.(*_kv).v = v
	}
}

func (this _GtreapMap) Len() int {
	return int(atomic.LoadInt64(&this.count))
}

func (this *_GtreapMap) Has(k interface{}) bool {
	o := this.m.Get(&_kv{k: k})
	return gtreap.Item(nil) != o
}

func (this _GtreapMap) EachFrom(kprefix interface{}, cb func(k, v interface{}) bool) {
	this.m.VisitAscend(&_kv{k: kprefix}, func(v gtreap.Item) bool {
		o := v.(*_kv)
		return cb(o.k, o.v)
	})
}

func (this _GtreapMap) Each(cb func(k, v interface{}) bool) {
	if 0 == atomic.LoadInt64(&this.count) {
		return
	}
	this.EachFrom(this.m.Min().(*_kv).k, cb)
}

type _GtreapMap struct {
	m     *gtreap.Treap
	count int64
}
