package gotrees

type TreeMap[K Comparable, V any] struct {
	TreeSet[K]
}

func NewTreeMap[K Comparable, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{
		TreeSet: TreeSet[K]{
			Size: 0,
			Root: nil,
		},
	}
}

func (tree *TreeMap[K, V]) AddItem(key K, value V) {
	node := &KeyValueNode[K, V]{
		KeyNode: KeyNode[K]{
			height: 1,
			key:    key,
			left:   nil,
			right:  nil,
			parent: nil,
		},
		value: value,
	}

	tree.Root = tree.insertNode(tree.Root, node)
	tree.Root.SetParent(nil)
}

func (tree *TreeMap[K, V]) DeleteItem(key K) {
	tree.deleteNode(tree.Find(key))
}

func (tree *TreeMap[K, V]) Find(key K) *KeyValueNode[K, V] {
	node := tree.find(key)
	if node == nil {
		return nil
	} else {
		return node.(*KeyValueNode[K, V])
	}
}

func (tree *TreeMap[K, V]) Contains(key K) bool {
	return tree.find(key) != nil
}

func (tree *TreeMap[K, V]) Min() *KeyValueNode[K, V] {
	return tree.TreeSet.Min().(*KeyValueNode[K, V])
}

func (tree *TreeMap[K, V]) Max() *KeyValueNode[K, V] {
	return tree.TreeSet.Max().(*KeyValueNode[K, V])
}

func (tree *TreeMap[K, V]) Next(n TreeNode[K]) *KeyValueNode[K, V] {
	next := tree.TreeSet.Next(n)
	if next != nil {
		return next.(*KeyValueNode[K, V])
	} else {
		return nil
	}
}

func (tree *TreeMap[K, V]) Prev(n TreeNode[K]) *KeyValueNode[K, V] {
	prev := tree.TreeSet.Prev(n)
	if prev != nil {
		return prev.(*KeyValueNode[K, V])
	} else {
		return nil
	}
}

func (tree *TreeMap[K, V]) GetKeys(traversal int) []K {
	return tree.TreeSet.GetKeys(traversal)
}
