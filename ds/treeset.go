package ds

const (
	TreeWalkInOrder   = iota
	TreeWalkPreOrder  = iota
	TreeWalkPostOrder = iota
	TreeWalkBFS       = iota
)

// AVL tree implementation

type Tree[K Comparable] struct {
	Size uint32
	Root TreeNode[K]
}

func NewTree[K Comparable]() *Tree[K] {
	return &Tree[K]{
		Size: 0,
		Root: nil,
	}
}

func getMaxSubtreeHeight[K Comparable](node TreeNode[K]) int {
	lh := 0
	rh := 0
	if node.GetLeft() != nil {
		lh = node.GetLeft().GetHeight()
	}

	if node.GetRight() != nil {
		rh = node.GetRight().GetHeight()
	}

	return Max(lh, rh)
}

func getSubtreeBalance[K Comparable](node TreeNode[K]) int {
	if node == nil {
		return 0
	}

	lh := 0
	rh := 0

	if node.GetLeft() != nil {
		lh = node.GetLeft().GetHeight()
	}

	if node.GetRight() != nil {
		rh = node.GetRight().GetHeight()
	}

	return lh - rh
}

func getSubtreeMin[K Comparable](node TreeNode[K]) TreeNode[K] {
	n := node
	for n.GetLeft() != nil {
		n = n.GetLeft()
	}
	return n
}

func getSubtreeMax[K Comparable](node TreeNode[K]) TreeNode[K] {
	n := node
	for n.GetRight() != nil {
		n = n.GetRight()
	}
	return n
}

func (tree *Tree[K]) walkKeys(n TreeNode[K], elements []K, order int) []K {
	if order == TreeWalkInOrder {
		if n.GetLeft() != nil {
			elements = tree.walkKeys(n.GetLeft(), elements, order)
		}
		elements = append(elements, n.GetKey())
		if n.GetRight() != nil {
			elements = tree.walkKeys(n.GetRight(), elements, order)
		}
	} else if order == TreeWalkPreOrder {
		elements = append(elements, n.GetKey())
		if n.GetLeft() != nil {
			elements = tree.walkKeys(n.GetLeft(), elements, order)
		}
		if n.GetRight() != nil {
			elements = tree.walkKeys(n.GetRight(), elements, order)
		}
	} else if order == TreeWalkPostOrder {
		if n.GetLeft() != nil {
			elements = tree.walkKeys(n.GetLeft(), elements, order)
		}
		if n.GetRight() != nil {
			elements = tree.walkKeys(n.GetRight(), elements, order)
		}
		elements = append(elements, n.GetKey())
	}

	return elements
}

func (tree *Tree[K]) Next(n TreeNode[K]) TreeNode[K] {
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

func (tree *Tree[K]) Prev(n TreeNode[K]) TreeNode[K] {
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

func (tree *Tree[K]) Min() TreeNode[K] {
	if tree.Root == nil {
		return nil
	}
	return getSubtreeMin(tree.Root)
}

func (tree *Tree[K]) Max() TreeNode[K] {
	if tree.Root == nil {
		return nil
	}
	return getSubtreeMax(tree.Root)
}

func rotateLeft[K Comparable](node TreeNode[K]) TreeNode[K] {
	newroot := node.GetRight()
	tmp := newroot.GetLeft()
	newroot.SetLeft(node)
	node.SetRight(tmp)

	if tmp != nil {
		tmp.SetParent(node)
	}

	node.SetHeight(getMaxSubtreeHeight(node) + 1)
	newroot.SetHeight(getMaxSubtreeHeight(newroot) + 1)
	return newroot
}

func rotateRight[K Comparable](node TreeNode[K]) TreeNode[K] {
	newroot := node.GetLeft()
	tmp := newroot.GetRight()
	newroot.SetRight(node)
	node.SetLeft(tmp)

	if tmp != nil {
		tmp.SetParent(node)
	}

	node.SetHeight(getMaxSubtreeHeight(node) + 1)
	newroot.SetHeight(getMaxSubtreeHeight(newroot) + 1)
	return newroot
}

func replaceNode[K Comparable](old TreeNode[K], new TreeNode[K]) {
	old.SetKey(new.GetKey())
	old.SetHeight(new.GetHeight())
	old.SetLeft(new.GetLeft())
	old.SetRight(new.GetRight())
	old.SetParent(new.GetParent())
}

func (tree *Tree[K]) insertNode(root TreeNode[K], node TreeNode[K]) TreeNode[K] {
	if root == nil {
		root = node
		tree.Size += 1
		return root
	}

	if node.GetKey() < root.GetKey() {
		root.SetLeft(tree.insertNode(root.GetLeft(), node))
		root.GetLeft().SetParent(root)
	} else if node.GetKey() > root.GetKey() {
		root.SetRight(tree.insertNode(root.GetRight(), node))
		root.GetRight().SetParent(root)
	} else {
		replaceNode(root, node)
	}

	var lheight = 0
	var rheight = 0

	if root.GetLeft() != nil {
		lheight = root.GetLeft().GetHeight()
	}
	if root.GetRight() != nil {
		rheight = root.GetRight().GetHeight()
	}

	root.SetHeight(Max(lheight, rheight) + 1)
	balance := lheight - rheight

	if balance > 1 && node.GetKey() < root.GetLeft().GetKey() {
		root = rotateRight(root)
	} else if balance < -1 && node.GetKey() > root.GetRight().GetKey() {
		root = rotateLeft(root)
	} else if balance > 1 && node.GetKey() > root.GetLeft().GetKey() {
		root.SetLeft(rotateLeft(root.GetLeft()))
		root = rotateRight(root)
	} else if balance < -1 && node.GetKey() < root.GetRight().GetKey() {
		root.SetRight(rotateRight(root.GetRight()))
		root = rotateLeft(root)
	}

	if root.GetLeft() != nil {
		root.GetLeft().SetParent(root)
	}

	if root.GetRight() != nil {
		root.GetRight().SetParent(root)
	}

	return root
}

func (tree *Tree[K]) deleteNode(root TreeNode[K], node TreeNode[K]) TreeNode[K] {
	if root == nil {
		return root
	}

	if node.GetKey() < root.GetKey() {
		root.SetLeft(tree.deleteNode(root.GetLeft(), node))
	} else if node.GetKey() > root.GetKey() {
		root.SetRight(tree.deleteNode(root.GetRight(), node))
	} else {
		if root.GetLeft() == nil || root.GetRight() == nil {
			var tmp TreeNode[K]
			if root.GetLeft() == nil {
				tmp = root.GetRight()
			} else {
				tmp = root.GetLeft()
			}

			if tmp != nil {
				root = tmp
			}
		} else {
			replaceNode(root, getSubtreeMax(root.GetLeft()))
			root.SetLeft(tree.deleteNode(root.GetLeft(), root))
		}
	}

	if root == nil {
		return root
	}

	root.SetHeight(getMaxSubtreeHeight(root) + 1)
	balance := getSubtreeBalance(root)

	if balance > 1 && getSubtreeBalance(root.GetLeft()) >= 0 {
		return rotateRight(root)
	} else if balance > 1 && getSubtreeBalance(root.GetLeft()) < 0 {
		root.SetLeft(rotateLeft(root.GetLeft()))
		return rotateRight(root)
	} else if balance < -1 && getSubtreeBalance(root.GetRight()) <= 0 {
		return rotateLeft(root)
	} else if balance < -1 && getSubtreeBalance(root.GetRight()) > 0 {
		root.SetRight(rotateRight(root.GetRight()))
		return rotateLeft(root)
	}

	root.SetParent(nil)
	return root
}

func AddKey[K Comparable](tree *Tree[K], key K) {
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

func AddKeyValue[K Comparable, V any](tree *Tree[K], key K, value V) {
	node := &KeyValueNode[K, V]{
		KeyNode: KeyNode[K]{
			key:    key,
			height: 1,
			left:   nil,
			right:  nil,
			parent: nil,
		},
		value: value,
	}

	tree.Root = tree.insertNode(tree.Root, node)
	tree.Root.SetParent(nil)
}

func (tree *Tree[K]) Find(key K) TreeNode[K] {
	if tree.Root == nil || tree.Root.GetKey() == key {
		return tree.Root
	}

	node := tree.Root
	for node != nil {
		if node.GetKey() == key {
			return node
		} else {
			if key < node.GetKey() {
				node = node.GetLeft()
			} else {
				node = node.GetRight()
			}
		}
	}

	return nil
}

func (tree *Tree[K]) Contains(key K) bool {
	return tree.Find(key) != nil
}

func bfs[K Comparable](node TreeNode[K], depthmap map[int][]K, depth int) {
	_, present := depthmap[depth]
	if !present {
		depthmap[depth] = make([]K, 0, 10)
	}

	depthmap[depth] = append(depthmap[depth], node.GetKey())
	if node.GetLeft() != nil {
		bfs(node.GetLeft(), depthmap, depth+1)
	}
	if node.GetRight() != nil {
		bfs(node.GetRight(), depthmap, depth+1)
	}
}

func (tree *Tree[K]) GetKeys(traversal int) []K {
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
