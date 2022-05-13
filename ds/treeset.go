package ds

const (
	TreeWalkInOrder   = iota
	TreeWalkPreOrder  = iota
	TreeWalkPostOrder = iota
	TreeWalkBFS       = iota
)

type Node[K Comparable] struct {
	height int
	Key    K
	Left   *Node[K]
	Right  *Node[K]
	Parent *Node[K]
}

// AVL tree implementation

type TreeSet[K Comparable] struct {
	Size uint32
	Root *Node[K]
}

func NewTreeSet[K Comparable]() *TreeSet[K] {
	return &TreeSet[K]{
		Size: 0,
		Root: nil,
	}
}

func (node *Node[K]) copyTo(target *Node[K]) {
	target.Key = node.Key
}

func (node *Node[K]) getMaxSubtreeHeight() int {
	lh := 0
	rh := 0

	if node.Left != nil {
		lh = node.Left.height
	}

	if node.Right != nil {
		rh = node.Right.height
	}

	return Max(lh, rh)
}

func (node *Node[K]) getSubtreeBalance() int {
	if node == nil {
		return 0
	}

	lh := 0
	rh := 0

	if node.Left != nil {
		lh = node.Left.height
	}

	if node.Right != nil {
		rh = node.Right.height
	}

	return lh - rh
}

func (node *Node[K]) getSubtreeMin() *Node[K] {
	n := node
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (node *Node[K]) getSubtreeMax() *Node[K] {
	n := node
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func getSubtreeBalance[K Comparable](node *Node[K]) int {
	if node == nil {
		return 0
	} else {
		var lh = 0
		var rh = 0
		if node.Left != nil {
			lh = node.Left.height
		}
		if node.Right != nil {
			rh = node.Right.height
		}
		return lh - rh
	}
}

func (tree *TreeSet[K]) walkKeys(n *Node[K], elements []K, order int) []K {
	if order == TreeWalkInOrder {
		if n.Left != nil {
			elements = tree.walkKeys(n.Left, elements, order)
		}
		elements = append(elements, n.Key)
		if n.Right != nil {
			elements = tree.walkKeys(n.Right, elements, order)
		}
	} else if order == TreeWalkPreOrder {
		elements = append(elements, n.Key)
		if n.Left != nil {
			elements = tree.walkKeys(n.Left, elements, order)
		}
		if n.Right != nil {
			elements = tree.walkKeys(n.Right, elements, order)
		}
	} else if order == TreeWalkPostOrder {
		if n.Left != nil {
			elements = tree.walkKeys(n.Left, elements, order)
		}
		if n.Right != nil {
			elements = tree.walkKeys(n.Right, elements, order)
		}
		elements = append(elements, n.Key)
	}

	return elements
}

func (tree *TreeSet[K]) Next(n *Node[K]) *Node[K] {
	if n.Right != nil {
		return n.Right.getSubtreeMin()
	}

	parent := n.Parent
	for parent != nil && parent.Right != nil && n.Key == parent.Right.Key {
		n = parent
		parent = parent.Parent
	}

	return parent
}

func (tree *TreeSet[K]) Prev(n *Node[K]) *Node[K] {
	if n.Left != nil {
		return n.Left.getSubtreeMax()
	}

	parent := n.Parent
	for parent != nil && parent.Left != nil && n.Key == parent.Left.Key {
		n = parent
		parent = parent.Parent
	}

	return parent
}

func (tree *TreeSet[K]) Min() *Node[K] {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.getSubtreeMin()
}

func (tree *TreeSet[K]) Max() *Node[K] {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.getSubtreeMax()
}

func (node *Node[K]) rotateLeft() *Node[K] {
	newroot := node.Right
	tmp := newroot.Left
	newroot.Left = node
	node.Right = tmp

	if tmp != nil {
		tmp.Parent = node
	}

	node.height = node.getMaxSubtreeHeight() + 1
	newroot.height = newroot.getMaxSubtreeHeight() + 1
	return newroot
}

func (node *Node[K]) rotateRight() *Node[K] {
	newroot := node.Left
	tmp := newroot.Right
	newroot.Right = node
	node.Left = tmp

	if tmp != nil {
		tmp.Parent = node
	}

	node.height = node.getMaxSubtreeHeight() + 1
	newroot.height = newroot.getMaxSubtreeHeight() + 1
	return newroot
}

func (tree *TreeSet[K]) insertNode(root *Node[K], node *Node[K]) *Node[K] {
	if root == nil {
		root = node
		tree.Size += 1
		return root
	}

	if node.Key < root.Key {
		root.Left = tree.insertNode(root.Left, node)
		root.Left.Parent = root
	} else if node.Key > root.Key {
		root.Right = tree.insertNode(root.Right, node)
		root.Right.Parent = root
	} else {
		node.copyTo(root) // replace if Key found
	}

	var lheight = 0
	var rheight = 0

	if root.Left != nil {
		lheight = root.Left.height
	}
	if root.Right != nil {
		rheight = root.Right.height
	}

	root.height = Max(lheight, rheight) + 1
	balance := lheight - rheight

	if balance > 1 && node.Key < root.Left.Key {
		root = root.rotateRight()
	} else if balance < -1 && node.Key > root.Right.Key {
		root = root.rotateLeft()
	} else if balance > 1 && node.Key > root.Left.Key {
		root.Left = root.Left.rotateLeft()
		root = root.rotateRight()
	} else if balance < -1 && node.Key < root.Right.Key {
		root.Right = root.Right.rotateRight()
		root = root.rotateLeft()
	}

	if root.Left != nil {
		root.Left.Parent = root
	}

	if root.Right != nil {
		root.Right.Parent = root
	}

	return root
}

func (tree *TreeSet[K]) deleteNode(root *Node[K], node *Node[K]) *Node[K] {
	if root == nil {
		return root
	}

	if node.Key < root.Key {
		root.Left = tree.deleteNode(root.Left, node)
	} else if node.Key > root.Key {
		root.Right = tree.deleteNode(root.Right, node)
	} else {
		if root.Left == nil || root.Right == nil {
			var tmp *Node[K]
			if root.Left == nil {
				tmp = root.Right
			} else {
				tmp = root.Left
			}

			if tmp != nil {
				root = tmp
			}
		} else {
			root.Left.getSubtreeMax().copyTo(root)
			root.Left = tree.deleteNode(root.Left, root)
		}
	}

	if root == nil {
		return root
	}

	root.height = root.getMaxSubtreeHeight() + 1
	balance := getSubtreeBalance(root)

	if balance > 1 && root.Left.getSubtreeBalance() >= 0 {
		return root.rotateRight()
	} else if balance > 1 && getSubtreeBalance(root.Left) < 0 {
		root.Left = root.Left.rotateLeft()
		return root.rotateRight()
	} else if balance < -1 && getSubtreeBalance(root.Right) <= 0 {
		return root.rotateLeft()
	} else if balance < -1 && getSubtreeBalance(root.Right) > 0 {
		root.Right = root.Right.rotateRight()
		return root.rotateLeft()
	}

	root.Parent = nil
	return root
}

func (tree *TreeSet[K]) Add(key K) {
	node := &Node[K]{
		height: 1,
		Key:    key,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}

	tree.Root = tree.insertNode(tree.Root, node)
	tree.Root.Parent = nil
}

func (tree *TreeSet[K]) Find(key K) *Node[K] {
	if tree.Root == nil || tree.Root.Key == key {
		return tree.Root
	}

	node := tree.Root
	for node != nil {
		if node.Key == key {
			return node
		} else {
			if key < node.Key {
				node = node.Left
			} else {
				node = node.Right
			}
		}
	}

	return nil
}

func (tree *TreeSet[K]) Contains(key K) bool {
	return tree.Find(key) != nil
}

func (node *Node[K]) bfs(depthmap map[int][]K, depth int) {
	_, present := depthmap[depth]
	if !present {
		depthmap[depth] = make([]K, 0, 10)
	}

	depthmap[depth] = append(depthmap[depth], node.Key)
	if node.Left != nil {
		node.Left.bfs(depthmap, depth+1)
	}
	if node.Right != nil {
		node.Right.bfs(depthmap, depth+1)
	}
}

func (tree *TreeSet[K]) GetKeys(traversal int) []K {
	keys := make([]K, 0)
	if tree.Root == nil {
		return keys
	}

	if traversal == TreeWalkBFS {
		_map := make(map[int][]K, 10)
		tree.Root.bfs(_map, 0)
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
