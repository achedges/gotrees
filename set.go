package gotrees

import "cmp"

type TreeSet[K cmp.Ordered] struct {
	size uint32
	root treeNode[K]
}

func NewTreeSet[K cmp.Ordered]() *TreeSet[K] {
	return &TreeSet[K]{
		size: 0,
		root: nil,
	}
}

func (t *TreeSet[K]) Size() uint32 {
	return t.size
}

func (t *TreeSet[K]) Add(key K) {
	node := newKeyNode(key)
	newRoot := t.insertNode(t.root, node)
	t.root = newRoot
	t.root.setParent(nil)
}

func (t *TreeSet[K]) Remove(key K) {
	t.deleteNode(key)
}

func (t *TreeSet[K]) Contains(key K) bool {
	return t.find(key) != nil
}

func (t *TreeSet[K]) Min() (K, bool){
	if t.root == nil {
		var empty K
		return empty, false
	}
	subtreeMin := getSubtreeMin(t.root)
	return subtreeMin.getKey(), true
}

func (t *TreeSet[K]) Max() (K, bool) {
	if t.root == nil {
		var empty K
		return empty, false
	}
	subtreeMax := getSubtreeMax(t.root)
	return subtreeMax.getKey(), true
}

func (t *TreeSet[K]) Next(key K) (K, bool) {
	var retval K
	found := false

	n := t.find(key)
	if n == nil {
		return retval, false
	}

	if n.getRight() != nil {
		subtreeMin := getSubtreeMin(n.getRight())
		if subtreeMin != nil {
			retval = subtreeMin.getKey()
			found = true
		}
	} else {
		parent := n.getParent()
		for parent != nil && parent.getRight() != nil && n.getKey() == parent.getRight().getKey() {
			n = parent
			parent = parent.getParent()
		}
		if parent != nil {
			retval = parent.getKey()
			found = true
		}
	}

	return retval, found
}

func (t *TreeSet[K]) Prev(key K) (K, bool) {
	var retval K
	found := false

	n := t.find(key)
	if n == nil {
		return retval, false
	}

	if n.getLeft() != nil {
		subtreeMin := getSubtreeMin(n.getLeft())
		if subtreeMin != nil {
			retval = subtreeMin.getKey()
			found = true
		}
	} else {
		parent := n.getParent()
		for parent != nil && parent.getLeft() != nil && n.getKey() == parent.getLeft().getKey() {
			n = parent
			parent = parent.getParent()
		}
		if parent != nil {
			retval = parent.getKey()
			found = true
		}
	}

	return retval, found
}

func (t *TreeSet[K]) GetKeys(traversal int) []K {
	keys := make([]K, 0)
	if t.root == nil {
		return keys
	}

	if traversal == TreeWalkBFS {
		_map := make(map[int][]K, 10)
		bfs(t.root, _map, 0)
		for d := 0; d < len(_map); d++ {
			for i := range _map[d] {
				keys = append(keys, _map[d][i])
			}
		}
	} else {
		keys = t.walkKeys(t.root, keys, traversal)
	}

	return keys
}
