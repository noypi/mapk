package mapk

import (
	"container/heap"
)

type _HeapMap struct {
	*_SliceMap
}

func MapHeap(cmp func(a, b interface{}) int) IMap {
	return &_HeapMap{MapSlice(cmp).(*_SliceMap)}
}

type _Heap _HeapMap

func (this *_HeapMap) Put(k, v interface{}) {
	i := this.find(k)
	if i < len(this._SliceMap.kvs) && 0 == this.cmp(this._SliceMap.kvs[i].k, k) {
		this._SliceMap.kvs[i].v = v
	} else {
		heap.Push((*_Heap)(this), &_kv{k: k, v: v})
	}
}

func (this _Heap) Less(i, j int) bool {
	// asc order
	return 0 > this._SliceMap.cmp(this._SliceMap.kvs[i].k, this._SliceMap.kvs[j].k)
}

func (this _Heap) Len() int {
	return len(this._SliceMap.kvs)
}

func (this *_Heap) Swap(i, j int) {
	this._SliceMap.kvs[i], this._SliceMap.kvs[j] = this._SliceMap.kvs[j], this._SliceMap.kvs[i]
}

func (this *_Heap) Push(x interface{}) {
	this._SliceMap.kvs = append(this._SliceMap.kvs, x.(*_kv))
}

func (this *_Heap) Pop() interface{} {
	panic("not supported")
	return nil
}
