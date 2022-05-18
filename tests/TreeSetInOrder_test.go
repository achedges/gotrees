package tests

import (
	"testing"
	"trees"
)

func getTreeSetInOrder() *trees.TreeSet[int] {
	tree := trees.NewTreeSet[int]()
	for _, v := range keys {
		tree.AddItem(v)
	}
	return tree
}

func TestTreeSetInOrderListSize(t *testing.T) {
	tree := getTreeSetInOrder()
	listSizeTestHelper(t, tree)
}

func TestTreeSetInOrderMinMax(t *testing.T) {
	tree := getTreeSetInOrder()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeSetInOrderNodeBoundaries(t *testing.T) {
	tree := getTreeSetInOrder()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeSetInOrderNextNodes(t *testing.T) {
	tree := getTreeSetInOrder()
	nextNodesTestHelper(t, tree)
}

func TestTreeSetInOrderPreviousNodes(t *testing.T) {
	tree := getTreeSetInOrder()
	previousNodesTestHelper(t, tree)
}

func TestTreeSetInOrderTraversal(t *testing.T) {
	tree := getTreeSetInOrder()
	preOrderKeys := []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9}
	postOrderKeys := []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3}
	bfsKeys := []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeSetInOrderFind(t *testing.T) {
	tree := getTreeSetInOrder()
	findTestHelper(t, tree, keys[:])
}
