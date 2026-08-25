package gotrees

import "cmp"

const (
	TreeWalkInOrder = iota
	TreeWalkPreOrder
	TreeWalkPostOrder
	TreeWalkBFS
)

func getMaxSubtreeHeight[K cmp.Ordered](node treeNode[K]) int {
	lh := 0
	rh := 0
	if node.getLeft() != nil {
		lh = node.getLeft().getHeight()
	}

	if node.getRight() != nil {
		rh = node.getRight().getHeight()
	}

	return max(lh, rh)
}

func getSubtreeMin[K cmp.Ordered](node treeNode[K]) treeNode[K] {
	if node == nil {
		return nil
	}

	n := node
	for n.getLeft() != nil {
		n = n.getLeft()
	}
	return n
}

func getSubtreeMax[K cmp.Ordered](node treeNode[K]) treeNode[K] {
	if node == nil {
		return nil
	}

	n := node
	for n.getRight() != nil {
		n = n.getRight()
	}
	return n
}

func rotateLeft[K cmp.Ordered](node treeNode[K]) treeNode[K] {
	newroot := node.getRight()
	tmp := newroot.getLeft()
	newroot.setLeft(node)
	node.setRight(tmp)

	if tmp != nil {
		tmp.setParent(node)
	}

	node.setHeight(getMaxSubtreeHeight(node) + 1)
	newroot.setHeight(getMaxSubtreeHeight(newroot) + 1)
	return newroot
}

func rotateRight[K cmp.Ordered](node treeNode[K]) treeNode[K] {
	newroot := node.getLeft()
	tmp := newroot.getRight()
	newroot.setRight(node)
	node.setLeft(tmp)

	if tmp != nil {
		tmp.setParent(node)
	}

	node.setHeight(getMaxSubtreeHeight(node) + 1)
	newroot.setHeight(getMaxSubtreeHeight(newroot) + 1)
	return newroot
}

func replaceNode[K cmp.Ordered](old treeNode[K], new treeNode[K]) treeNode[K] {
	new.setKey(old.getKey())
	new.setHeight(old.getHeight())
	new.setLeft(old.getLeft())
	new.setRight(old.getRight())
	new.setParent(old.getParent())
	return new
}

func bfs[K cmp.Ordered](node treeNode[K], depthmap map[int][]K, depth int) {
	_, present := depthmap[depth]
	if !present {
		depthmap[depth] = make([]K, 0, 10)
	}

	depthmap[depth] = append(depthmap[depth], node.getKey())
	if node.getLeft() != nil {
		bfs(node.getLeft(), depthmap, depth+1)
	}
	if node.getRight() != nil {
		bfs(node.getRight(), depthmap, depth+1)
	}
}

func (t *TreeSet[K]) insertNode(root treeNode[K], node treeNode[K]) treeNode[K] {
	if root == nil {
		root = node
		t.size += 1
		return root
	}

	if node.getKey() < root.getKey() {
		root.setLeft(t.insertNode(root.getLeft(), node))
		root.getLeft().setParent(root)
	} else if node.getKey() > root.getKey() {
		root.setRight(t.insertNode(root.getRight(), node))
		root.getRight().setParent(root)
	} else {
		root = replaceNode(root, node)
	}

	var lheight = 0
	var rheight = 0

	if root.getLeft() != nil {
		lheight = root.getLeft().getHeight()
	}
	if root.getRight() != nil {
		rheight = root.getRight().getHeight()
	}

	root.setHeight(max(lheight, rheight) + 1)
	balance := lheight - rheight

	if balance > 1 && node.getKey() < root.getLeft().getKey() {
		root = rotateRight(root)
	} else if balance < -1 && node.getKey() > root.getRight().getKey() {
		root = rotateLeft(root)
	} else if balance > 1 && node.getKey() > root.getLeft().getKey() {
		root.setLeft(rotateLeft(root.getLeft()))
		root = rotateRight(root)
	} else if balance < -1 && node.getKey() < root.getRight().getKey() {
		root.setRight(rotateRight(root.getRight()))
		root = rotateLeft(root)
	}

	if root.getLeft() != nil {
		root.getLeft().setParent(root)
	}

	if root.getRight() != nil {
		root.getRight().setParent(root)
	}

	return root
}

func (t *TreeSet[K]) transplantNode(old treeNode[K], new treeNode[K]) {
	if old.getParent() == nil {
		t.root = new
	} else if old.getParent().getLeft() != nil && old.getKey() == old.getParent().getLeft().getKey() {
		old.getParent().setLeft(new)
	} else {
		old.getParent().setRight(new)
	}

	if new != nil {
		new.setParent(old.getParent())
	}
}

func (t *TreeSet[K]) find(key K) treeNode[K] {
	node := t.root
	for node != nil {
		if node.getKey() == key {
			break
		}

		if key < node.getKey() {
			node = node.getLeft()
		} else {
			node = node.getRight()
		}
	}

	return node
}

func (t *TreeSet[K]) deleteNode(key K) {
	node := t.find(key)
	if node == nil {
		return
	}

	if node.getLeft() == nil {
		t.transplantNode(node, node.getRight())
	} else if node.getRight() == nil {
		t.transplantNode(node, node.getLeft())
	} else {
		y := getSubtreeMin(node.getRight())
		if y.getParent().getKey() != node.getKey() {
			t.transplantNode(y, y.getRight())
			y.setRight(node.getRight())
			y.getRight().setParent(y)
		}
		t.transplantNode(node, y)
		y.setLeft(node.getLeft())
		y.getLeft().setParent(y)
	}

	t.size -= 1
}

func (t *TreeSet[K]) walkKeys(n treeNode[K], elements []K, order int) []K {
	if n == nil {
		return elements
	}

	if order == TreeWalkInOrder {
		if n.getLeft() != nil {
			elements = t.walkKeys(n.getLeft(), elements, order)
		}
		elements = append(elements, n.getKey())
		if n.getRight() != nil {
			elements = t.walkKeys(n.getRight(), elements, order)
		}
	} else if order == TreeWalkPreOrder {
		elements = append(elements, n.getKey())
		if n.getLeft() != nil {
			elements = t.walkKeys(n.getLeft(), elements, order)
		}
		if n.getRight() != nil {
			elements = t.walkKeys(n.getRight(), elements, order)
		}
	} else if order == TreeWalkPostOrder {
		if n.getLeft() != nil {
			elements = t.walkKeys(n.getLeft(), elements, order)
		}
		if n.getRight() != nil {
			elements = t.walkKeys(n.getRight(), elements, order)
		}
		elements = append(elements, n.getKey())
	}

	return elements
}
