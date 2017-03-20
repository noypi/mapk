package mapk_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/noypi/mapk"
	assertpkg "github.com/stretchr/testify/assert"
)

func testmap01(m mapk.IMap, t *testing.T) {
	assert := assertpkg.New(t)

	m.Put("sa1", "v1")
	assert.Equal("v1", m.Get("sa1"))
	assert.Equal(1, m.Len())

	m.Put("sa1", "v2")
	m.Each(func(a, b interface{}) bool {
		fmt.Println("k=", a, "v=", b)
		return true
	})
	assert.Equal("v2", m.Get("sa1"))
	assert.Equal(1, m.Len())

	type _kv struct {
		k, v string
	}

	kvsexpected := []_kv{
		_kv{"sa1", "v100"},
		_kv{"sb3", "v3"},
		_kv{"sc4", "v4"},
		_kv{"sd5", "v5"},
		_kv{"se6", "v6"},
	}
	for _, v := range kvsexpected {
		m.Put(v.k, v.v)
	}
	assert.Equal(len(kvsexpected), m.Len())

	kvs := []_kv{}
	m.Each(func(a, b interface{}) bool {
		kvs = append(kvs, _kv{a.(string), b.(string)})
		return true
	})
	for i := 0; i < len(kvsexpected); i++ {
		assert.Equal(kvsexpected[i], kvs[i])
	}

	kvs = []_kv{}
	m.EachFrom("sc", func(a, b interface{}) bool {
		kvs = append(kvs, _kv{a.(string), b.(string)})
		return true
	})

	expected := kvsexpected[2:]
	assert.Equal(len(expected), len(kvs))
	for i := 0; i < len(expected); i++ {
		assert.Equal(expected[i], kvs[i])
	}
}

func TestMap01_Gtreap(t *testing.T) {
	m := mapk.Map(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	testmap01(m, t)
}

func TestMap01_Slice(t *testing.T) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	testmap01(m, t)
}
