package gotrees

import (
	"testing"

	"github.com/achedges/go-assertions"
)

func TestBase_getMaxSubtreeHeight(t *testing.T) {
	root := newKeyNode(2)
	root.setLeft(newKeyNode(1))
	root.left.setHeight(4)
	root.setRight(newKeyNode(3))
	root.right.setHeight(7)
	assertions.EqualInts(7, getMaxSubtreeHeight(root), t)
}

func TestBase_getSubtreeMinMax(t *testing.T) {
	root := newKeyNode(4)
	
	root.left = newKeyNode(2)
	root.left.setLeft(newKeyNode(1))
	root.left.setRight(newKeyNode(3))

	root.right = newKeyNode(6)
	root.right.setLeft(newKeyNode(5))
	root.right.setRight(newKeyNode(7))

	subtreeMin := getSubtreeMin(root)
	if subtreeMin == nil {
		t.Error("Expected subtree min node")
		t.FailNow()
	}
	assertions.EqualInts(1, subtreeMin.getKey(), t)

	subtreeMax := getSubtreeMax(root)
	if subtreeMax == nil {
		t.Error("Expected subtree max node")
		t.FailNow()
	}
	assertions.EqualInts(7, subtreeMax.getKey(), t)
}

func TestBase_replaceNode(t *testing.T) {
	a := newKeyNode(1)
	a.setHeight(3)

	b := newKeyNode(4)
	b.setHeight(7)

	c := replaceNode(a, b)
	assertions.EqualInts(b.getKey(), c.getKey(), t)
	assertions.EqualInts(b.getHeight(), c.getHeight(), t)
}