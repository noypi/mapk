package mapk

import (
	"strings"
)

type IMap interface {
	Get(k interface{}) interface{}
	Put(k, v interface{})
	Delete(k interface{})
	Has(k interface{}) bool
	EachFrom(kprefix interface{}, cb func(k, v interface{}) bool)
	Each(cb func(k, v interface{}) bool)
	Len() int
}

func Map(comp func(a, b interface{}) int) IMap {
	return MapGtreap(comp)
}

func CmpString(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

type _kv struct {
	k, v interface{}
}

type _kvslist []*_kv
