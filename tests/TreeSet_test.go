package tests

import "testing"

func TestTreeSetLookup(test *testing.T) {
	tree := getTreeSetInOrder()
	for _, v := range testKeys {
		key := tree.Find(v)
		if key.GetKey() != v {
			test.Errorf("Unexpected lookup key: %d", key.GetKey())
			test.FailNow()
		}
	}
}
