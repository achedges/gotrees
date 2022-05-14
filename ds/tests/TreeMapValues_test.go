package tests

import (
	"localutils/ds"
	"testing"
)

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

	//node1 := tree.Find(1)
}
