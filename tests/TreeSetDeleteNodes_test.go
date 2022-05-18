package tests

import (
	"testing"
)

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
	for _, v := range keys {
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
