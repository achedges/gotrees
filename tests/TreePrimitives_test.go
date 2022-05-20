package tests

import (
	"testing"
)

func TestTreeSize(test *testing.T) {
	tree := getTreeSetInOrder()
	if tree.Size != listSize {
		test.Errorf("Unexpected tree size: %d", tree.Size)
		test.FailNow()
	}
}

func TestTreeMinMax(test *testing.T) {
	tree := getTreeSetInOrder()
	if tree.Min().GetKey() != 0 {
		test.Errorf("Unexpected tree minimum: %d", tree.Min().GetKey())
		test.FailNow()
	}
	if tree.Max().GetKey() != 9 {
		test.Errorf("Unexpected tree maximum: %d", tree.Max().GetKey())
		test.FailNow()
	}
	if tree.Prev(tree.Min()) != nil {
		test.Error("Unexpected previous node from tree minimum")
		test.FailNow()
	}
	if tree.Next(tree.Max()) != nil {
		test.Error("Unexpected next node from tree maximum")
		test.FailNow()
	}
}

func TestTreeNextPrevNodes(test *testing.T) {
	tree := getTreeSetInOrder()

	var node = tree.Min()
	var next = tree.Next(node)
	var nodeKey = 0
	var nextKey = 0
	var prevKey = 0

	if next == nil {
		test.Error("Initial next node is nil")
		test.FailNow()
	}

	for next != nil {
		nodeKey = any(node.GetKey()).(int)
		nextKey = any(next.GetKey()).(int)
		if nextKey != nodeKey+1 {
			test.Errorf("Unexpected next key: %d", nextKey)
			test.FailNow()
		}
		node = next
		next = tree.Next(next)
	}

	node = tree.Max()
	var prev = tree.Prev(node)

	if prev == nil {
		test.Error("Initial previous node is nil")
		test.FailNow()
	}

	for prev != nil {
		nodeKey = any(node.GetKey()).(int)
		prevKey = any(prev.GetKey()).(int)
		if prevKey != nodeKey-1 {
			test.Errorf("Unexpected previous key: %d", prevKey)
			test.FailNow()
		}
		node = prev
		prev = tree.Prev(prev)
	}
}

func TestTreeInOrderTraversal(test *testing.T) {
	tree := getTreeSetInOrder()
	preOrderKeys := []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9}
	postOrderKeys := []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3}
	bfsKeys := []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9}
	traversalTestHelper(test, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeReversedTraversal(test *testing.T) {
	tree := getTreeSetReversed()
	preOrderKeys := []int{6, 2, 1, 0, 4, 3, 5, 8, 7, 9}
	postOrderKeys := []int{0, 1, 3, 5, 4, 2, 7, 9, 8, 6}
	bfsKeys := []int{6, 2, 8, 1, 4, 7, 9, 0, 3, 5}
	traversalTestHelper(test, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeScrambledTraversal(test *testing.T) {
	tree := getTreeSetScrambled()
	preOrderKeys := []int{4, 1, 0, 2, 3, 6, 5, 8, 7, 9}
	postOrderKeys := []int{0, 3, 2, 1, 5, 7, 9, 8, 6, 4}
	bfsKeys := []int{4, 1, 6, 0, 2, 5, 8, 3, 7, 9}
	traversalTestHelper(test, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeSetDeleteInnerNode(t *testing.T) {
	tree := getTreeSetInOrder()

	tree.DeleteItem(5)
	if tree.Size != 9 {
		t.Error("Unexpected tree size after delete")
		t.FailNow()
	}

	node := tree.Find(5)
	if node != nil {
		t.Error("Node found after delete")
		t.FailNow()
	}
}

func TestTreeSetDeleteMinNode(t *testing.T) {
	tree := getTreeSetInOrder()
	tree.DeleteItem(tree.Min().GetKey())
	min := tree.Min().GetKey()
	if min != 1 {
		t.Errorf("Unexpected minimum node after delete: %d", min)
		t.FailNow()
	}
}

func TestTreeSetDeleteMaxNode(t *testing.T) {
	tree := getTreeSetInOrder()
	tree.DeleteItem(tree.Max().GetKey())
	max := tree.Max().GetKey()
	if max != 8 {
		t.Errorf("Unexpected maximum node after delete: %d", max)
		t.FailNow()
	}
}

func TestTreeSetDeleteRootNode(t *testing.T) {
	tree := getTreeSetInOrder()
	tree.DeleteItem(tree.Root.GetKey())
	root := tree.Root.GetKey()
	if root != 4 {
		t.Errorf("Unexpected root node after delete: %d", root)
		t.FailNow()
	}
}

func TestTreeSetDeleteAllNodes(t *testing.T) {
	tree := getTreeSetInOrder()
	for _, v := range testKeys {
		tree.DeleteItem(v)
	}

	if tree.Size != 0 {
		t.Error("Tree is not empty after deleting all nodes")
		t.FailNow()
	}

	if tree.Root != nil {
		t.Error("Tree root is not nil after deleting all nodes")
		t.FailNow()
	}
}
