package gotrees

type KeyValueNode[K Comparable, V any] struct {
	KeyNode[K]
	value V
}

func (node *KeyValueNode[K, V]) GetKey() K {
	return node.key
}

func (node *KeyValueNode[K, V]) SetKey(key K) {
	node.key = key
}

func (node *KeyValueNode[K, V]) GetValue() V {
	return node.value
}

func (node *KeyValueNode[K, V]) SetValue(value V) {
	node.value = value
}

func (node *KeyValueNode[K, V]) GetHeight() int {
	return node.height
}

func (node *KeyValueNode[K, V]) SetHeight(height int) {
	node.height = height
}

func (node *KeyValueNode[K, V]) GetLeft() TreeNode[K] {
	return node.left
}

func (node *KeyValueNode[K, V]) SetLeft(left TreeNode[K]) {
	node.left = left
}

func (node *KeyValueNode[K, V]) GetRight() TreeNode[K] {
	return node.right
}

func (node *KeyValueNode[K, V]) SetRight(right TreeNode[K]) {
	node.right = right
}

func (node *KeyValueNode[K, V]) GetParent() TreeNode[K] {
	return node.parent
}

func (node *KeyValueNode[K, V]) SetParent(parent TreeNode[K]) {
	node.parent = parent
}
