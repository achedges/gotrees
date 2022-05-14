package tests

import (
	"localutils/ds"
	"testing"
)

func fail(t *testing.T, node int) {
	t.Errorf("Incorrect value at node %d", node)
	t.FailNow()
}

func TestTreeMapKeyValuePairs(t *testing.T) {
	tree := ds.NewTree[int]()
	ds.AddKeyValue(tree, 1, "A")
	ds.AddKeyValue(tree, 2, "B")
	ds.AddKeyValue(tree, 3, "C")
	ds.AddKeyValue(tree, 4, "D")

	if tree.Size != 4 {
		t.Error("Incorrect list size")
		t.FailNow()
	}

	node1 := tree.Find(1).(*ds.KeyValueNode[int, string])
	if node1.GetValue() != "A" {
		fail(t, 1)
	}

	node2 := tree.Find(2).(*ds.KeyValueNode[int, string])
	if node2.GetValue() != "B" {
		fail(t, 2)
	}

	node3 := tree.Find(3).(*ds.KeyValueNode[int, string])
	if node3.GetValue() != "C" {
		fail(t, 3)
	}

	node4 := tree.Find(4).(*ds.KeyValueNode[int, string])
	if node4.GetValue() != "D" {
		fail(t, 4)
	}
}
