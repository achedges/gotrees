package ds

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Comparable interface {
	Integer | Float | ~string
}

type TreeSet[K Comparable] struct {
	Size uint32
	Root TreeNode[K]
}

type TreeMap[K Comparable, V any] struct {
	TreeSet[K]
}

type TreeNode[K Comparable] interface {
	GetKey() K
	SetKey(key K)
	GetHeight() int
	SetHeight(height int)
	GetLeft() TreeNode[K]
	SetLeft(node TreeNode[K])
	GetRight() TreeNode[K]
	SetRight(node TreeNode[K])
	GetParent() TreeNode[K]
	SetParent(node TreeNode[K])
}

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
