package tests

import (
	"localutils/ds"
	"testing"
)

func getTreeMapInOrder() *ds.Tree[int] {
	tree := ds.NewTree[int]()
	for _, v := range keys {
		ds.AddKeyValue(tree, v, v)
	}
	return tree
}

func TestTreeMapInOrderListSize(t *testing.T) {
	tree := getTreeMapInOrder()
	listSizeTestHelper(t, tree)
}

func TestTreeMapInOrderMinMax(t *testing.T) {
	tree := getTreeMapInOrder()
	minMaxTestHelper(t, tree, 0, len(keys)-1)
}

func TestTreeMapInOrderNodeBoundaries(t *testing.T) {
	tree := getTreeMapInOrder()
	nodeBoundariesTestHelper(t, tree)
}

func TestTreeMapInOrderNextNodes(t *testing.T) {
	tree := getTreeMapInOrder()
	nextNodesTestHelper(t, tree)
}

func TestTreeMapInOrderPreviousNodes(t *testing.T) {
	tree := getTreeMapInOrder()
	previousNodesTestHelper(t, tree)
}

func TestTreeMapInOrderTraversal(t *testing.T) {
	tree := getTreeMapInOrder()
	preOrderKeys := []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9}
	postOrderKeys := []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3}
	bfsKeys := []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9}
	traversalTestHelper(t, tree, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeMapInOrderFind(t *testing.T) {
	tree := getTreeMapInOrder()
	findTestHelper(t, tree, keys[:])
}
