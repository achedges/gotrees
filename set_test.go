package gotrees_test

import (
	"testing"

	"github.com/achedges/go-assertions"
	"github.com/achedges/gotrees"
)

func TestTreeSet_NewTreeSet(t *testing.T) {
	set := gotrees.NewTreeSet[int]()
	assertions.EqualUints(0, set.Size(), t)
}

func TestTreeSet_Add_Remove(t *testing.T) {
	set := gotrees.NewTreeSet[int]()

	set.Add(3)
	assertions.EqualUints(1, set.Size(), t)

	set.Add(7)
	assertions.EqualUints(2, set.Size(), t)

	set.Remove(7)
	assertions.EqualUints(1, set.Size(), t)

	set.Remove(3)
	assertions.EqualUints(0, set.Size(), t)
}

func TestTreeSet_Contains(t *testing.T) {
	set := gotrees.NewTreeSet[int]()
	set.Add(3)
	set.Add(8)
	set.Add(12)

	assertions.True(set.Contains(3), t)
	assertions.True(set.Contains(8), t)
	assertions.True(set.Contains(12), t)

	assertions.False(set.Contains(1), t)
	assertions.False(set.Contains(10), t)
	assertions.False(set.Contains(9), t)
}

func TestTreeSet_Min_Max(t *testing.T) {
	set := gotrees.NewTreeSet[int]()

	// test empty state first
	key, found := set.Min()
	assertions.False(found, t)
	key, found = set.Max()
	assertions.False(found, t)

	for i := range 10 {
		set.Add(i)
	}

	key, found = set.Min()
	assertions.True(found, t)
	assertions.EqualInts(0, key, t)

	key, found = set.Max()
	assertions.True(found, t)
	assertions.EqualInts(9, key, t)
}

func TestTreeSet_Next_Prev(t *testing.T) {
	set := gotrees.NewTreeSet[int]()
	for i := range 10 {
		set.Add(i)
	}

	var nextkey int

	i, found := set.Min()
	for found {
		nextkey, found = set.Next(i)
		if !found {
			break
		}
		assertions.EqualInts(i+1, nextkey, t)
		i += 1
	}

	_, found = set.Next(9)
	assertions.False(found, t)
	assertions.EqualInts(9, i, t)

	var prevkey int

	i, found = set.Max()
	for found {
		prevkey, found = set.Prev(i)
		if !found {
			break
		}
		assertions.EqualInts(i-1, prevkey, t)
		i -= 1
	}

	_, found = set.Prev(0)
	assertions.False(found, t)
	assertions.EqualInts(0, i, t)
}

func TestTreeSet_GetKeys(t *testing.T) {
	set := gotrees.NewTreeSet[int]()
	for i := range 10 {
		set.Add(i)
	}

	inOrderKeys := set.GetKeys(gotrees.TreeWalkInOrder)
	assertions.EqualInts(10, len(inOrderKeys), t)
	for i, v := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} {
		assertions.EqualInts(v, inOrderKeys[i], t)
	}

	preOrderKeys := set.GetKeys(gotrees.TreeWalkPreOrder)
	assertions.EqualInts(10, len(preOrderKeys), t)
	for i, v := range []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9} {
		assertions.EqualInts(v, preOrderKeys[i], t)
	}

	postOrderKeys := set.GetKeys(gotrees.TreeWalkPostOrder)
	assertions.EqualInts(10, len(postOrderKeys), t)
	for i, v := range []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3} {
		assertions.EqualInts(v, postOrderKeys[i], t)
	}

	bfsKeys := set.GetKeys(gotrees.TreeWalkBFS)
	assertions.EqualInts(10, len(bfsKeys), t)
	for i, v := range []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9} {
		assertions.EqualInts(v, bfsKeys[i], t)
	}
}
