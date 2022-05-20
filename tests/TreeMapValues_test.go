package tests

import (
	"testing"
	"trees"
)

func fail(t *testing.T, node int) {
	t.Errorf("Incorrect value at node %d", node)
	t.FailNow()
}

func TestTreeMapKeyValuePairs(t *testing.T) {
	testValues := map[int]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
	}

	tree := gotrees.NewTreeMap[int, string]()
	for k, v := range testValues {
		tree.AddItem(k, v)
	}

	if tree.Size != 4 {
		t.Error("Incorrect list size")
		t.FailNow()
	}

	for k, v := range testValues {
		node := tree.Find(k)
		if node.GetValue() != v {
			fail(t, k)
		}
	}

	node := tree.Min()
	if node.GetKey() != 1 {
		t.Errorf("Unexpected tree map min: %d", node.GetKey())
		t.FailNow()
	}

	if node.GetValue() != "A" {
		t.Errorf("Unexpected tree map min value: %s", node.GetValue())
		t.FailNow()
	}
}
