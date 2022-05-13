package tests

import (
	"fmt"
	"localutils/ds"
	"testing"
)

const listSize uint32 = 10

var keys = [listSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func listSizeTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K]) {
	if tree.Size != listSize {
		test.Errorf("Unexpected tree size: %d", tree.Size)
		test.FailNow()
	}
}

func minMaxTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K], expMin K, expMax K) {
	if tree.Min().Key != expMin {
		test.Errorf("Unexpected tree minimum: %d", any(tree.Min().Key).(int))
		test.FailNow()
	}
	if tree.Max().Key != expMax {
		test.Errorf("Unexpected tree maximum %d", any(tree.Max().Key).(int))
		test.FailNow()
	}
}

func nodeBoundariesTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K]) {
	if tree.Root.Parent != nil {
		test.Error("Invalid parent node on tree root")
		test.FailNow()
	}
	if tree.Prev(tree.Min()) != nil {
		test.Error("Invalid previous node on tree minimum")
		test.FailNow()
	}
	if tree.Next(tree.Max()) != nil {
		test.Error("Invalid next node on tree maximum")
		test.FailNow()
	}
}

func nextNodesTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K]) {
	var cur = tree.Min()
	var nex = tree.Next(cur)
	if nex == nil {
		test.Error("Initial next node is nil")
		test.FailNow()
	}

	for nex != nil {
		exp := any(cur.Key).(int) + 1
		if any(nex.Key).(int) != exp {
			test.Error("Incorrect next key")
			test.FailNow()
		}
		cur = nex
		nex = tree.Next(nex)
	}
}

func previousNodesTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K]) {
	var cur = tree.Max()
	var pre = tree.Prev(cur)
	if pre == nil {
		test.Error("Initial previous node is nil")
		test.FailNow()
	}

	for pre != nil {
		exp := any(cur.Key).(int) - 1
		if any(pre.Key).(int) != exp {
			test.Error("Incorrect previous key")
			test.FailNow()
		}
		cur = pre
		pre = tree.Prev(pre)
	}
}

func traversalTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K], preOrderKeys []K, postOrderKeys []K, bfsKeys []K) {
	var i uint32 = 0
	var treeKeys = tree.GetKeys(ds.TreeWalkInOrder)

	for i < tree.Size {
		if any(treeKeys[i]).(int) != keys[i] {
			test.Error("Incorrect in-order key traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(ds.TreeWalkPreOrder)

	for i < tree.Size {
		if treeKeys[i] != preOrderKeys[i] {
			test.Error("Incorrect pre-order traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(ds.TreeWalkPostOrder)

	for i < tree.Size {
		if treeKeys[i] != postOrderKeys[i] {
			test.Error("Incorrect post-order traversal")
			test.FailNow()
		}
		i++
	}

	i = 0
	treeKeys = tree.GetKeys(ds.TreeWalkBFS)

	for i < tree.Size {
		if treeKeys[i] != bfsKeys[i] {
			test.Error("Incorrect breadth-first traversal")
			test.FailNow()
		}
		i++
	}
}

func findTestHelper[K ds.Comparable](test *testing.T, tree *ds.TreeSet[K], testKeys []K) {
	for _, v := range testKeys {
		val := tree.Find(v).Key
		if val != v {
			test.Errorf("Incorrect value found: %s", fmt.Sprint(val))
			test.FailNow()
		}
	}
}
