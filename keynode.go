package gotrees

type KeyNode[K Comparable] struct {
	key    K
	height int
	left   TreeNode[K]
	right  TreeNode[K]
	parent TreeNode[K]
}

func (node *KeyNode[K]) GetKey() K {
	return node.key
}

func (node *KeyNode[K]) SetKey(key K) {
	node.key = key
}

func (node *KeyNode[K]) GetValue() any {
	return node.key
}

func (node *KeyNode[K]) SetValue(_ any) {
	return
}

func (node *KeyNode[K]) GetHeight() int {
	return node.height
}

func (node *KeyNode[K]) SetHeight(height int) {
	node.height = height
}

func (node *KeyNode[K]) GetLeft() TreeNode[K] {
	return node.left
}

func (node *KeyNode[K]) SetLeft(left TreeNode[K]) {
	node.left = left
}

func (node *KeyNode[K]) GetRight() TreeNode[K] {
	return node.right
}

func (node *KeyNode[K]) SetRight(right TreeNode[K]) {
	node.right = right
}

func (node *KeyNode[K]) GetParent() TreeNode[K] {
	return node.parent
}

func (node *KeyNode[K]) SetParent(parent TreeNode[K]) {
	node.parent = parent
}
