package tests

import (
	"testing"
	"trees"
)

func getTreeSetScrambled() *trees.TreeSet[int] {
	tree := trees.NewTreeSet[int]()
	insertionOrder := []int{0, 2, 1, 6, 4, 5, 3, 9, 7, 8}
	for _, v := range insertionOrder {
		tree.AddItem(keys[v])
	}
	return tree
}

func TestTreeSetScrambledListSize(t *testing.T) {
	tree := getTreeSetScrambled()
	listSizeTestHelper(t, tree)
}

func TestTreeSetScrambledMinMax(t *testing.T) {
	tree := getTreeSetScrambled()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeSetScrambledNodeBoundaries(t *testing.T) {
	tree := getTreeSetScrambled()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeSetScrambledNextNodes(t *testing.T) {
	tree := getTreeSetScrambled()
	nextNodesTestHelper(t, tree)
}

func TestTreeSetScrambledPreviousNodes(t *testing.T) {
	tree := getTreeSetScrambled()
	previousNodesTestHelper(t, tree)
}

func TestTreeSetScrambledTraversal(t *testing.T) {
	tree := getTreeSetScrambled()
	preOrderKeys := []int{4, 1, 0, 2, 3, 6, 5, 8, 7, 9}
	postOrderKeys := []int{0, 3, 2, 1, 5, 7, 9, 8, 6, 4}
	bfsKeys := []int{4, 1, 6, 0, 2, 5, 8, 3, 7, 9}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeSetScrambledFind(t *testing.T) {
	tree := getTreeSetScrambled()
	findTestHelper(t, tree, keys[:])
}
