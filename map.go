package gotrees

import "cmp"

type TreeMap[K cmp.Ordered, V any] struct {
	TreeSet[K]
}

func NewTreeMap[K cmp.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{
		TreeSet: TreeSet[K]{
			size: 0,
			root: nil,
		},
	}
}

func (m *TreeMap[K, V]) Size() uint32 {
	return m.size
}

func (tree *TreeMap[K, V]) Add(key K, value V) {
	node := &keyValueNode[K, V]{
		keyNode: *newKeyNode(key),
		value:   value,
	}

	tree.root = tree.insertNode(tree.root, node)
	tree.root.setParent(nil)
}

func (tree *TreeMap[K, V]) Remove(key K) {
	tree.deleteNode(key)
}

func (tree *TreeMap[K, V]) Get(key K) (V, bool) {
	if node := tree.find(key); node != nil {
		return node.(*keyValueNode[K, V]).value, true
	}

	var empty V
	return empty, false
}

func (tree *TreeMap[K, V]) Contains(key K) bool {
	return tree.find(key) != nil
}
