package mapk

import (
	"sort"
)

type _SliceMap struct {
	kvs _kvslist
	cmp func(k1, k2 interface{}) int
}

func MapSlice(cmp func(k, v interface{}) int) IMap {
	return &_SliceMap{
		cmp: cmp,
	}
}

func (this _SliceMap) find(k interface{}) int {
	return sort.Search(len(this.kvs), func(i int) bool {
		return 0 <= this.cmp(this.kvs[i].k, k)
	})
}

func (this _SliceMap) less(i, j int) bool {
	return this.cmp(this.kvs[i].k, this.kvs[j].k) < 0
}

func (this _SliceMap) Get(k interface{}) interface{} {
	i := this.find(k)
	if i < len(this.kvs) {
		return this.kvs[i].v
	}
	return nil
}

func (this *_SliceMap) Put(k, v interface{}) {
	if 0 < len(this.kvs) {
		i := this.find(k)
		if i < len(this.kvs) {
			this.kvs[i].v = v
			return
		}
	}

	this.kvs = append(this.kvs, &_kv{k: k, v: v})
	sort.Slice(this.kvs, this.less)
}

func (this *_SliceMap) Delete(k interface{}) {
	i := this.find(k)
	if i >= len(this.kvs) {
		return
	}

	this.kvs = append(this.kvs[0:i], this.kvs[i+1:]...)
}

func (this _SliceMap) Has(k interface{}) bool {
	return this.find(k) < len(this.kvs)
}

func (this _SliceMap) EachFrom(kprefix interface{}, cb func(k, v interface{}) bool) {
	i := this.find(kprefix)
	if i >= len(this.kvs) {
		i = sort.Search(len(this.kvs), func(i int) bool {
			return 0 < this.cmp(this.kvs[i].k, kprefix)
		})
	}
	if i < len(this.kvs) {
		for _, v := range this.kvs[i:] {
			if !cb(v.k, v.v) {
				break
			}
		}
	}
}

func (this _SliceMap) Each(cb func(k, v interface{}) bool) {
	for _, v := range this.kvs {
		if !cb(v.k, v.v) {
			break
		}
	}
}
func (this _SliceMap) Len() int { return len(this.kvs) }
