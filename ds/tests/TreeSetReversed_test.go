package tests

import (
	"localutils/ds"
	"testing"
)

func getTreeSetReversed() *ds.TreeSet[int] {
	tree := ds.NewTreeSet[int]()
	n := len(keys)
	for i := range keys {
		tree.AddItem(keys[n-1-i])
	}
	return tree
}

func TestTreeSetReversedListSize(t *testing.T) {
	tree := getTreeSetReversed()
	listSizeTestHelper(t, tree)
}

func TestTreeSetReversedMinMax(t *testing.T) {
	tree := getTreeSetReversed()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeSetReversedNodeBoundaries(t *testing.T) {
	tree := getTreeSetReversed()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeSetReversedNextNodes(t *testing.T) {
	tree := getTreeSetReversed()
	nextNodesTestHelper(t, tree)
}

func TestTreeSetReversedPreviousNodes(t *testing.T) {
	tree := getTreeSetReversed()
	previousNodesTestHelper(t, tree)
}

func TestTreeSetReversedTraversal(t *testing.T) {
	tree := getTreeSetReversed()
	preOrderKeys := []int{6, 2, 1, 0, 4, 3, 5, 8, 7, 9}
	postOrderKeys := []int{0, 1, 3, 5, 4, 2, 7, 9, 8, 6}
	bfsKeys := []int{6, 2, 8, 1, 4, 7, 9, 0, 3, 5}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeSetReversedFind(t *testing.T) {
	tree := getTreeSetReversed()
	findTestHelper(t, tree, keys[:])
}
