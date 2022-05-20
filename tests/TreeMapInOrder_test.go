package tests

import (
	"testing"
	"trees"
)

func getTreeMapInOrder() *gotrees.TreeMap[int, string] {
	tree := gotrees.NewTreeMap[int, string]()
	for _, v := range keys { // have to add nodes in the same order to match expected tree
		tree.AddItem(v, keyValues[v])
	}
	return tree
}

func TestTreeMapInOrderListSize(t *testing.T) {
	tree := getTreeMapInOrder()
	listSizeTestHelperMap(t, tree)
}

func TestTreeMapInOrderMinMax(t *testing.T) {
	tree := getTreeMapInOrder()
	minMaxTestHelperMap(t, tree, 0, len(keys)-1)
}

func TestTreeMapInOrderNodeBoundaries(t *testing.T) {
	tree := getTreeMapInOrder()
	nodeBoundariesTestHelperMap(t, tree)
}

func TestTreeMapInOrderNextNodes(t *testing.T) {
	tree := getTreeMapInOrder()
	nextNodesTestHelperMap(t, tree)
}

func TestTreeMapInOrderPreviousNodes(t *testing.T) {
	tree := getTreeMapInOrder()
	previousNodesTestHelperMap(t, tree)
}

func TestTreeMapInOrderTraversal(t *testing.T) {
	tree := getTreeMapInOrder()
	preOrderKeys := []int{3, 1, 0, 2, 7, 5, 4, 6, 8, 9}
	postOrderKeys := []int{0, 2, 1, 4, 6, 5, 9, 8, 7, 3}
	bfsKeys := []int{3, 1, 7, 0, 2, 5, 8, 4, 6, 9}
	traversalTestHelper(t, &tree.TreeSet, preOrderKeys, postOrderKeys, bfsKeys)
}

func TestTreeMapInOrderFind(t *testing.T) {
	tree := getTreeMapInOrder()
	findTestHelperMap(t, tree, keys[:])
}
