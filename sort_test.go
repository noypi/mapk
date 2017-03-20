package mapk

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	assertpkg "github.com/stretchr/testify/assert"
)

func testSortHundred(m IMap, t *testing.T, nsize, ranlen int) {
	assert := assertpkg.New(t)

	s := struct{}{}
	nPrevSize := 0
	sorted := []int{}
	for i := 0; i < nsize; i++ {
		n := rand.Intn(ranlen)
		has := m.Has(n)
		m.Put(n, s)
		if !has {
			assert.Equal(nPrevSize+1, m.Len())
		} else {
			assert.Equal(nPrevSize, m.Len())
		}
		if nPrevSize != m.Len() {
			nPrevSize = m.Len()
			sorted = append(sorted, n)
		}
	}

	sort.Sort(sort.IntSlice(sorted))
	assert.Equal(len(sorted), m.Len())

	i := 0
	m.Each(func(k, v interface{}) bool {
		assert.Equal(sorted[i], k.(int))
		i++
		return true
	})
}

func testSortHundred_withloop(m IMap, t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))
	assert := assertpkg.New(t)

	for i := 1; i < 50; i++ {
		testSortHundred(m, t, 200, 20*i)
		m.Clear()
		assert.Equal(0, m.Len())

		//
		testSortHundred(MakeThreadSafe(m), t, 200, 20*i)
		m.Clear()
		assert.Equal(0, m.Len())
	}
}

func TestSortHundred_Slice(t *testing.T) {
	m := MapSlice(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	testSortHundred_withloop(m, t)

}

func TestSortHundred_GTreap(t *testing.T) {
	m := MapGTreap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	testSortHundred_withloop(m, t)
}
