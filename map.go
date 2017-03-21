package mapk

import (
	"bytes"
	"fmt"
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
	Clear()
}

func Map(comp func(a, b interface{}) int) IMap {
	return MapSlice(comp)
}

func CmpString(a, b interface{}) int {
	return strings.Compare(a.(string), b.(string))
}

type _kv struct {
	k, v interface{}
}

type _kvslist []*_kv

func (this _kvslist) String() string {
	buf := bytes.NewBufferString("[")
	for _, v := range this {
		buf.WriteString(fmt.Sprintf("(%v,%v)", v.k, v.v))
	}
	buf.WriteString("]")
	return buf.String()
}
