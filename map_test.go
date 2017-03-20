package mapk_test

import (
	"strings"
	"testing"

	"github.com/noypi/mapk"
	assertpkg "github.com/stretchr/testify/assert"
)

type _kv struct {
	k, v string
}

func testmap01(m mapk.IMap, t *testing.T) {
	assert := assertpkg.New(t)

	m.Put("sa1", "v1")
	assert.Equal("v1", m.Get("sa1"))
	assert.Equal(1, m.Len())

	m.Put("sa1", "v2")
	assert.Equal("v2", m.Get("sa1"))
	assert.Equal(1, m.Len())

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
	m := mapk.MapGTreap(func(a, b interface{}) int {
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

func TestMap01_ThreadSafe_Slice(t *testing.T) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	testmap01(mapk.MakeThreadSafe(m), t)
}

func TestMap01_ThreadSafe_GTreap(t *testing.T) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})

	testmap01(mapk.MakeThreadSafe(m), t)
}

var ttDataTen01 = []_kv{
	_kv{"sa1", "v100"},
	_kv{"sa2", "v100"},
	_kv{"sb3", "v3"},
	_kv{"sc4", "v4"},
	_kv{"sd5", "v5"},
	_kv{"se6", "v6"},
	_kv{"se7", "v6"},
	_kv{"se8", "v6"},
	_kv{"se9", "v6"},
	_kv{"se10", "v6"},
}

func benchmap_putten(m mapk.IMap) {
	for _, v := range ttDataTen01 {
		m.Put(v.k, v.v)
	}
}

func benchmap_getten(m mapk.IMap) {
	for _, v := range ttDataTen01 {
		m.Get(v.k)
	}
}

func benchmap_eachfrompartial7of10(m mapk.IMap) {
	m.EachFrom("sc", func(a, b interface{}) bool {
		return true
	})
}

func benchmap_eachten(m mapk.IMap) {
	m.Each(func(a, b interface{}) bool {
		return true
	})
}

func benchmap_delete5of10(m mapk.IMap) {
	m.Delete("sa1")
	m.Delete("se10")
	m.Delete("se6")
	m.Delete("sd5")
	m.Delete("sb3")
	m.Put("sa1", "1")
	m.Put("se10", "1")
	m.Put("se6", "1")
	m.Put("sd5", "1")
	m.Put("sb3", "1")
}

func BenchmarkPutTen_GTreap(b *testing.B) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_putten(m)
	}
}

func BenchmarkGetTen_GTreap(b *testing.B) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_getten(m)
	}
}

func BenchmarkEachFrom7of10_GTreap(b *testing.B) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_eachfrompartial7of10(m)
	}
}

func BenchmarkEachTen_GTreap(b *testing.B) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_eachten(m)
	}
}

func BenchmarkDeleteAdd5of10_GTreap(b *testing.B) {
	m := mapk.MapGTreap(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_delete5of10(m)
	}
}

func BenchmarkPutTen_Slice(b *testing.B) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_putten(m)
	}
}

func BenchmarkGetTen_Slice(b *testing.B) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_getten(m)
	}
}

func BenchmarkEachFrom7of10_Slice(b *testing.B) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_eachfrompartial7of10(m)
	}
}

func BenchmarkEachTen_Slice(b *testing.B) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_eachten(m)
	}
}

func BenchmarkDeleteAdd5of10_Slice(b *testing.B) {
	m := mapk.MapSlice(func(a, b interface{}) int {
		return strings.Compare(a.(string), b.(string))
	})
	benchmap_putten(m)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmap_delete5of10(m)
	}
}

func BenchmarkPutTen_Native(b *testing.B) {
	m := map[string]string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range ttDataTen01 {
			m[v.k] = v.v
		}
	}
}

func BenchmarkGetTen_Native(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < b.N; i++ {
		for _, v := range ttDataTen01 {
			m[v.k] = v.v
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range ttDataTen01 {
			_, _ = m[v.k]
		}
	}
}

func BenchmarkDeleteAdd5of10_Native(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < b.N; i++ {
		for _, v := range ttDataTen01 {
			m[v.k] = v.v
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		delete(m, "sa1")
		delete(m, "se10")
		delete(m, "se6")
		delete(m, "sd5")
		delete(m, "sb3")
		m["sa1"] = "1"
		m["se10"] = "1"
		m["se6"] = "1"
		m["sd5"] = "1"
		m["sb3"] = "1"
	}
}

func BenchmarkEachFrom_Native_NOT_SUPPORTED(b *testing.B) {
}

func BenchmarkEachTen_Native(b *testing.B) {
	m := map[string]string{}
	for i := 0; i < b.N; i++ {
		for _, v := range ttDataTen01 {
			m[v.k] = v.v
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, _ = range m {

		}
	}
}
