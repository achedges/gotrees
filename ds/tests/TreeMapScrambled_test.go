package tests

import (
	"localutils/ds"
	"testing"
)

func getTreeMapScrambled() *ds.Tree[int] {
	tree := ds.NewTree[int]()
	insertionOrder := []int{0, 2, 1, 6, 4, 5, 3, 9, 7, 8}
	for _, v := range insertionOrder {
		ds.AddKeyValue(tree, keys[v], keys[v])
	}
	return tree
}

func TestTreeMapScrambledListSize(t *testing.T) {
	tree := getTreeMapScrambled()
	listSizeTestHelper(t, tree)
}

func TestTreeMapScrambledMinMax(t *testing.T) {
	tree := getTreeMapScrambled()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeMapScrambledNodeBoundaries(t *testing.T) {
	tree := getTreeMapScrambled()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeMapScrambledNextNodes(t *testing.T) {
	tree := getTreeMapScrambled()
	nextNodesTestHelper(t, tree)
}

func TestTreeMapScrambledPreviousNodes(t *testing.T) {
	tree := getTreeMapScrambled()
	previousNodesTestHelper(t, tree)
}

func TestTreeMapScrambledTraversal(t *testing.T) {
	tree := getTreeMapScrambled()
	preOrderKeys := []int{4, 1, 0, 2, 3, 6, 5, 8, 7, 9}
	postOrderKeys := []int{0, 3, 2, 1, 5, 7, 9, 8, 6, 4}
	bfsKeys := []int{4, 1, 6, 0, 2, 5, 8, 3, 7, 9}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeMapScrambledFind(t *testing.T) {
	tree := getTreeMapScrambled()
	findTestHelper(t, tree, keys[:])
}
