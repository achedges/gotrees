package tests

import (
	"fmt"
	"testing"
	"trees"
)

const listSize uint32 = 10

var keys = [listSize]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var keyValues = map[int]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
	4: "E",
	5: "F",
	6: "G",
	7: "H",
	8: "I",
	9: "J",
}

func listSizeTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K]) {
	if tree.Size != listSize {
		test.Errorf("Unexpected tree size: %d", tree.Size)
		test.FailNow()
	}
}

func listSizeTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V]) {
	if tree.Size != listSize {
		test.Errorf("Unexpected tree size: %d", tree.Size)
		test.FailNow()
	}
}

func minMaxTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K], expMin K, expMax K) {
	if tree.Min().GetKey() != expMin {
		test.Errorf("Unexpected tree minimum: %s", fmt.Sprint(tree.Min().GetKey()))
		test.FailNow()
	}
	if tree.Max().GetKey() != expMax {
		test.Errorf("Unexpected tree maximum %s", fmt.Sprint(tree.Max().GetKey()))
		test.FailNow()
	}
}

func minMaxTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V], expMin K, expMax K) {
	if tree.Min().GetKey() != expMin {
		test.Errorf("Unexpected tree minimum: %s", fmt.Sprint(tree.Min().GetKey()))
		test.FailNow()
	}
	if tree.Max().GetKey() != expMax {
		test.Errorf("Unexpected tree maximum: %s", fmt.Sprint(tree.Max().GetKey()))
		test.FailNow()
	}
}

func nodeBoundariesTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K]) {
	if tree.Root.GetParent() != nil {
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

func nodeBoundariesTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V]) {
	if tree.Root.GetParent() != nil {
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

func nextNodesTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K]) {
	var cur = tree.Min()
	var nex = tree.Next(cur)
	if nex == nil {
		test.Error("Initial next node is nil")
		test.FailNow()
	}

	for nex != nil {
		exp := any(cur.GetKey()).(int) + 1
		if any(nex.GetKey()).(int) != exp {
			test.Error("Incorrect next key")
			test.FailNow()
		}
		cur = nex
		nex = tree.Next(nex)
	}
}

func nextNodesTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V]) {
	var cur = tree.Min()
	var nex = tree.Next(cur)
	if nex == nil {
		test.Error("Initial next node is nil")
		test.FailNow()
	}

	for nex != nil {
		exp := any(cur.GetKey()).(int) + 1
		if any(nex.GetKey()).(int) != exp {
			test.Error("Incorrect next key")
			test.FailNow()
		}
		cur = nex
		nex = tree.Next(nex)
	}
}

func previousNodesTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K]) {
	var cur = tree.Max()
	var pre = tree.Prev(cur)
	if pre == nil {
		test.Error("Initial previous node is nil")
		test.FailNow()
	}

	for pre != nil {
		exp := any(cur.GetKey()).(int) - 1
		if any(pre.GetKey()).(int) != exp {
			test.Error("Incorrect previous key")
			test.FailNow()
		}
		cur = pre
		pre = tree.Prev(pre)
	}
}

func previousNodesTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V]) {
	var cur = tree.Max()
	var pre = tree.Prev(cur)
	if pre == nil {
		test.Error("Initial previous node is nil")
		test.FailNow()
	}

	for pre != nil {
		exp := any(cur.GetKey()).(int) - 1
		if any(pre.GetKey()).(int) != exp {
			test.Error("Incorrect previous key")
			test.FailNow()
		}
		cur = pre
		pre = tree.Prev(pre)
	}
}

func traversalTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K], preOrderKeys []K, postOrderKeys []K, bfsKeys []K) {
	var i uint32 = 0
	var treeKeys = tree.GetKeys(gotrees.TreeWalkInOrder)

	for i < tree.Size {
		if any(treeKeys[i]).(int) != keys[i] {
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

func findTestHelper[K gotrees.Comparable](test *testing.T, tree *gotrees.TreeSet[K], testKeys []K) {
	for _, v := range testKeys {
		val := tree.Find(v).GetKey()
		if val != v {
			test.Errorf("Incorrect value found: %s", fmt.Sprint(val))
			test.FailNow()
		}
	}
}

func findTestHelperMap[K gotrees.Comparable, V any](test *testing.T, tree *gotrees.TreeMap[K, V], testKeys []K) {
	for _, v := range testKeys {
		val := tree.Find(v).GetKey()
		if val != v {
			test.Errorf("Incorrect value found: %s", fmt.Sprint(val))
			test.FailNow()
		}
	}
}
