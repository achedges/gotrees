package tests

import (
	"github.com/achedges/gotrees"
	"testing"
)

const listSize uint32 = 10

var testKeys = [listSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var testKeyValues = map[int]string{
	1: "A",
	2: "B",
	3: "C",
	4: "D",
}

func getTreeSetInOrder() *gotrees.TreeSet[int] {
	tree := gotrees.NewTreeSet[int]()
	for _, v := range testKeys {
		tree.AddItem(v)
	}
	return tree
}

func getTreeSetReversed() *gotrees.TreeSet[int] {
	tree := gotrees.NewTreeSet[int]()
	n := len(testKeys)
	for i := range testKeys {
		tree.AddItem(testKeys[n-1-i])
	}
	return tree
}

func getTreeSetScrambled() *gotrees.TreeSet[int] {
	tree := gotrees.NewTreeSet[int]()
	insertionOrder := []int{0, 2, 1, 6, 4, 5, 3, 9, 7, 8}
	for _, v := range insertionOrder {
		tree.AddItem(testKeys[v])
	}
	return tree
}

func traversalTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K], preOrderKeys []K, postOrderKeys []K, bfsKeys []K) {
	var i uint32 = 0
	var treeKeys = tree.GetKeys(gotrees.TreeWalkInOrder)

	for i < tree.Size {
		if any(treeKeys[i]).(int) != testKeys[i] {
			test.Error("Incorrect in-order key traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(gotrees.TreeWalkPreOrder)

	for i < tree.Size {
		if treeKeys[i] != preOrderKeys[i] {
			test.Error("Incorrect pre-order traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(gotrees.TreeWalkPostOrder)

	for i < tree.Size {
		if treeKeys[i] != postOrderKeys[i] {
			test.Error("Incorrect post-order traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(gotrees.TreeWalkBFS)

	for i < tree.Size {
		if treeKeys[i] != bfsKeys[i] {
			test.Error("Incorrect breadth-first traversal")
			test.FailNow()
		}
		i++
	}
}
