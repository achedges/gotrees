package gotrees

const (
	TreeWalkInOrder   = iota
	TreeWalkPreOrder  = iota
	TreeWalkPostOrder = iota
	TreeWalkBFS       = iota
)

// AVL tree implementation

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

func (tree *TreeSet[K]) walkKeys(n TreeNode[K], elements []K, order int) []K {
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

func replaceNode[K Comparable](old TreeNode[K], new TreeNode[K]) TreeNode[K] {
	new.SetKey(old.GetKey())
	new.SetHeight(old.GetHeight())
	new.SetLeft(old.GetLeft())
	new.SetRight(old.GetRight())
	new.SetParent(old.GetParent())
	return new
}

func (tree *TreeSet[K]) insertNode(root TreeNode[K], node TreeNode[K]) TreeNode[K] {
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
		root = replaceNode(root, node)
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

func (tree *TreeSet[K]) transplantNode(old TreeNode[K], new TreeNode[K]) {
	if old.GetParent() == nil {
		tree.Root = new
	} else if old.GetParent().GetLeft() != nil && old.GetKey() == old.GetParent().GetLeft().GetKey() {
		old.GetParent().SetLeft(new)
	} else {
		old.GetParent().SetRight(new)
	}

	if new != nil {
		new.SetParent(old.GetParent())
	}
}

func (tree *TreeSet[K]) deleteNode(node TreeNode[K]) {
	if node == nil {
		return
	}

	if node.GetLeft() == nil {
		tree.transplantNode(node, node.GetRight())
	} else if node.GetRight() == nil {
		tree.transplantNode(node, node.GetLeft())
	} else {
		y := getSubtreeMin(node.GetRight())
		if y.GetParent().GetKey() != node.GetKey() {
			tree.transplantNode(y, y.GetRight())
			y.SetRight(node.GetRight())
			y.GetRight().SetParent(y)
		}
		tree.transplantNode(node, y)
		y.SetLeft(node.GetLeft())
		y.GetLeft().SetParent(y)
	}

	tree.Size -= 1
}

func (tree *TreeSet[K]) find(key K) TreeNode[K] {
	node := tree.Root
	for node != nil {
		if node.GetKey() == key {
			break
		}

		if key < node.GetKey() {
			node = node.GetLeft()
		} else {
			node = node.GetRight()
		}
	}

	return node
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
