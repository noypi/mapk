package mapk_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/noypi/mapk"
	assertpkg "github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assertpkg.New(t)

	m := mapk.Map(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	m.Put("s1", "v1")
	assert.Equal("v1", m.Get("s1"))
	assert.Equal(1, m.Len())

	m.Put("s1", "v2")
	m.Each(func(a, b interface{}) bool {
		fmt.Println("k=", a, "v=", b)
		return true
	})
	assert.Equal("v2", m.Get("s1"))
	assert.Equal(1, m.Len())

	type _kv struct {
		k, v string
	}

	kvsexpected := []_kv{
		_kv{"s1", "v100"},
		_kv{"s3", "v3"},
		_kv{"s4", "v4"},
		_kv{"s5", "v5"},
		_kv{"s6", "v6"},
	}
	for _, v := range kvsexpected {
		m.Put(v.k, v.v)
	}

	kvs := []_kv{}
	m.Each(func(a, b interface{}) bool {
		kvs = append(kvs, _kv{a.(string), b.(string)})
		return true
	})

	assert.Equal(len(kvsexpected), m.Len())
	for i := 0; i < len(kvsexpected); i++ {
		assert.Equal(kvsexpected[i], kvs[i])
	}

	kvs = []_kv{}
	m.EachFrom("s4", func(a, b interface{}) bool {
		kvs = append(kvs, _kv{a.(string), b.(string)})
		return true
	})

	expected := kvsexpected[2:]
	for i := 0; i < len(expected); i++ {
		assert.Equal(expected[i], kvs[i])
	}
}
