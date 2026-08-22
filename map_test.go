package gotrees_test

import (
	"fmt"
	"testing"

	"github.com/achedges/go-assertions"
	"github.com/achedges/gotrees"
)

func TestTreeMap_NewTreeMap(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()
	assertions.EqualUints(0, m.Size(), t)
}

func TestTreeMap_Add_Remove(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()
	
	m.Add(1, "A")
	assertions.EqualUints(1, m.Size(), t)

	m.Add(3, "C")
	assertions.EqualUints(2, m.Size(), t)

	m.Remove(1)
	assertions.EqualUints(1, m.Size(), t)

	m.Remove(3)
	assertions.EqualUints(0, m.Size(), t)
}

func TestTreeMap_Contains(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()
	m.Add(1, "A")
	m.Add(2, "B")
	m.Add(3, "C")

	assertions.True(m.Contains(1), t)
	assertions.True(m.Contains(2), t)
	assertions.True(m.Contains(3), t)

	assertions.False(m.Contains(0), t)
	assertions.False(m.Contains(4), t)
}

func TestTreeMap_Get(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()
	m.Add(1, "A")
	
	value, found := m.Get(1)
	assertions.True(found, t)
	assertions.EqualStrings("A", value, t)

	value, found = m.Get(0)
	assertions.False(found, t)
}

func TestTreeMap_Min_Max(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()

	// test empty state first
	key, found := m.Min()
	assertions.False(found, t)
	key, found = m.Max()
	assertions.False(found, t)

	m.Add(1, "A")
	m.Add(2, "B")
	m.Add(3, "C")

	key, found = m.Min()
	assertions.True(found, t)
	assertions.EqualInts(1, key, t)

	key, found = m.Max()
	assertions.True(found, t)
	assertions.EqualInts(3, key, t)
}

func TestTreeMap_Next_Prev(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()

	// test empty state first
	key, found := m.Next(0)
	assertions.False(found, t)
	key, found = m.Prev(0)
	assertions.False(found, t)

	m.Add(1, "A")
	m.Add(2, "B")
	m.Add(3, "C")

	key, found = m.Prev(2)
	assertions.True(found, t)
	assertions.EqualInts(1, key, t)

	key, found = m.Next(2)
	assertions.True(found, t)
	assertions.EqualInts(3, key, t)
}

func TestTreeMap_GetKeys(t *testing.T) {
	m := gotrees.NewTreeMap[int, string]()
	for i := range 10 {
		m.Add(i, fmt.Sprintf("%d", i))
	}

	inOrderKeys := m.GetKeys(gotrees.TreeWalkInOrder)
	assertions.EqualInts(10, len(inOrderKeys), t)
	for i, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		assertions.EqualInts(v, inOrderKeys[i], t)
	}

	preOrderKeys := m.GetKeys(gotrees.TreeWalkPreOrder)
	assertions.EqualInts(10, len(preOrderKeys), t)
	for i, v := range []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9} {
		assertions.EqualInts(v, preOrderKeys[i], t)
	}

	postOrderKeys := m.GetKeys(gotrees.TreeWalkPostOrder)
	assertions.EqualInts(10, len(postOrderKeys), t)
	for i, v := range []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3} {
		assertions.EqualInts(v, postOrderKeys[i], t)
	}

	bfsKeys := m.GetKeys(gotrees.TreeWalkBFS)
	assertions.EqualInts(10, len(bfsKeys), t)
	for i, v := range []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9} {
		assertions.EqualInts(v, bfsKeys[i], t)
	}
}