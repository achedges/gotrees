package gotrees

type TreeSet[K Comparable] struct {
	Size uint32
	Root TreeNode[K]
}

func NewTreeSet[K Comparable]() *TreeSet[K] {
	return &TreeSet[K]{
		Size: 0,
		Root: nil,
	}
}

func (tree *TreeSet[K]) AddItem(key K) {
	node := &KeyNode[K]{
		height: 1,
		key:    key,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	tree.Root = tree.insertNode(tree.Root, node)
	tree.Root.SetParent(nil)
}

func (tree *TreeSet[K]) DeleteItem(key K) {
	tree.deleteNode(tree.Find(key))
}

func (tree *TreeSet[K]) Find(key K) *KeyNode[K] {
	node := tree.find(key)
	if node == nil {
		return nil
	} else {
		return node.(*KeyNode[K])
	}
}

func (tree *TreeSet[K]) Contains(key K) bool {
	return tree.find(key) != nil
}

func (tree *TreeSet[K]) Min() TreeNode[K] {
	if tree.Root == nil {
		return nil
	}
	return getSubtreeMin(tree.Root)
}

func (tree *TreeSet[K]) Max() TreeNode[K] {
	if tree.Root == nil {
		return nil
	}
	return getSubtreeMax(tree.Root)
}

func (tree *TreeSet[K]) Next(n TreeNode[K]) TreeNode[K] {
	if n.GetRight() != nil {
		return getSubtreeMin(n.GetRight())
	}

	parent := n.GetParent()
	for parent != nil && parent.GetRight() != nil && n.GetKey() == parent.GetRight().GetKey() {
		n = parent
		parent = parent.GetParent()
	}

	return parent
}

func (tree *TreeSet[K]) Prev(n TreeNode[K]) TreeNode[K] {
	if n.GetLeft() != nil {
		return getSubtreeMax(n.GetLeft())
	}

	parent := n.GetParent()
	for parent != nil && parent.GetLeft() != nil && n.GetKey() == parent.GetLeft().GetKey() {
		n = parent
		parent = parent.GetParent()
	}

	return parent
}

func (tree *TreeSet[K]) GetKeys(traversal int) []K {
	keys := make([]K, 0)
	if tree.Root == nil {
		return keys
	}

	if traversal == TreeWalkBFS {
		_map := make(map[int][]K, 10)
		bfs(tree.Root, _map, 0)
		for d := 0; d < len(_map); d++ {
			for i := range _map[d] {
				keys = append(keys, _map[d][i])
			}
		}
	} else {
		keys = tree.walkKeys(tree.Root, keys, traversal)
	}

	return keys
}
