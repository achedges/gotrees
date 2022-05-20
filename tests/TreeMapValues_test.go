package tests

import (
	"gotrees"
	"testing"
)

func TestTreeMapKeyValuePairs(test *testing.T) {
	tree := gotrees.NewTreeMap[int, string]()
	for k, v := range testKeyValues {
		tree.AddItem(k, v)
	}

	if tree.Size != 4 {
		test.Error("Incorrect list size")
		test.FailNow()
	}

	for k, v := range testKeyValues {
		node := tree.Find(k)
		if node.GetValue() != v {
			test.Errorf("Incorrect value at node %d", k)
			test.FailNow()
		}
	}

	var node = tree.Min()
	if node.GetKey() != 1 {
		test.Errorf("Unexpected tree map min: %d", node.GetKey())
		test.FailNow()
	}

	if node.GetValue() != "A" {
		test.Errorf("Unexpected tree map min value: %s", node.GetValue())
		test.FailNow()
	}

	node = tree.Max()
	if node.GetKey() != 4 {
		test.Errorf("Unexpected tree map max: %d", node.GetKey())
		test.FailNow()
	}

	if node.GetValue() != "D" {
		test.Errorf("Unexpected tree map max value: %s", node.GetValue())
		test.FailNow()
	}
}
