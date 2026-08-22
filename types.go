package gotrees

import "cmp"

type treeNode[K cmp.Ordered] interface {
	getKey() K
	setKey(key K)
	getHeight() int
	setHeight(height int)
	getLeft() treeNode[K]
	getRight() treeNode[K]
	getParent() treeNode[K]
	setLeft(node treeNode[K])
	setRight(node treeNode[K])
	setParent(node treeNode[K])
}

// key node

type keyNode[K cmp.Ordered] struct {
	key    K
	height int
	left   treeNode[K]
	right  treeNode[K]
	parent treeNode[K]
}

func newKeyNode[K cmp.Ordered](key K) *keyNode[K] {
	return &keyNode[K]{
		height: 1,
		key:    key,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (n *keyNode[K]) getKey() K {
	return n.key
}

func (n *keyNode[K]) setKey(key K) {
	n.key = key
}

func (n *keyNode[K]) getHeight() int {
	return n.height
}

func (n *keyNode[K]) setHeight(height int) {
	n.height = height
}

func (n *keyNode[K]) getLeft() treeNode[K] {
	return n.left
}

func (n *keyNode[K]) getRight() treeNode[K] {
	return n.right
}

func (n *keyNode[K]) getParent() treeNode[K] {
	return n.parent
}

func (n *keyNode[K]) setLeft(left treeNode[K]) {
	n.left = left
}

func (n *keyNode[K]) setRight(right treeNode[K]) {
	n.right = right
}

func (n *keyNode[K]) setParent(parent treeNode[K]) {
	n.parent = parent
}

// key-value node

type keyValueNode[K cmp.Ordered, V any] struct {
	keyNode[K]
	value V
}

func (n *keyValueNode[K, V]) getKey() K {
	return n.key
}

func (n *keyValueNode[K, V]) setKey(key K) {
	n.key = key
}

func (n *keyValueNode[K, V]) getHeight() int {
	return n.height
}

func (n *keyValueNode[K, V]) setHeight(height int) {
	n.height = height
}

func (n *keyValueNode[K, V]) getLeft() treeNode[K] {
	return n.left
}

func (n *keyValueNode[K, V]) setLeft(left treeNode[K]) {
	n.left = left
}

func (n *keyValueNode[K, V]) getRight() treeNode[K] {
	return n.right
}

func (n *keyValueNode[K, V]) setRight(right treeNode[K]) {
	n.right = right
}

func (n *keyValueNode[K, V]) getParent() treeNode[K] {
	return n.parent
}

func (n *keyValueNode[K, V]) setParent(parent treeNode[K]) {
	n.parent = parent
}
