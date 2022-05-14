package tests

import (
	"localutils/ds"
	"testing"
)

func getTreeMapReversed() *ds.Tree[int] {
	tree := ds.NewTree[int]()
	n := len(keys)
	for i := range keys {
		ds.AddKeyValue(tree, keys[n-1-i], keys[n-1-i])
	}
	return tree
}

func TestTreeMapReversedListSize(t *testing.T) {
	tree := getTreeMapReversed()
	listSizeTestHelper(t, tree)
}

func TestTreeMapReversedMinMax(t *testing.T) {
	tree := getTreeMapReversed()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeMapReversedNodeBoundaries(t *testing.T) {
	tree := getTreeMapReversed()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeMapReversedNextNodes(t *testing.T) {
	tree := getTreeMapReversed()
	nextNodesTestHelper(t, tree)
}

func TestTreeMapReversedPreviousNodes(t *testing.T) {
	tree := getTreeMapReversed()
	previousNodesTestHelper(t, tree)
}

func TestTreeMapReversedTraversal(t *testing.T) {
	tree := getTreeMapReversed()
	preOrderKeys := []int{6, 2, 1, 0, 4, 3, 5, 8, 7, 9}
	postOrderKeys := []int{0, 1, 3, 5, 4, 2, 7, 9, 8, 6}
	bfsKeys := []int{6, 2, 8, 1, 4, 7, 9, 0, 3, 5}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeMapReversedFind(t *testing.T) {
	tree := getTreeMapReversed()
	findTestHelper(t, tree, keys[:])
}
