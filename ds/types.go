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

type TreeNode[K Comparable, V any] interface {
	Key() K
	Value() V
	Height() int
	Left() *TreeNode[K, V]
	Right() *TreeNode[K, V]
	Parent() *TreeNode[K, V]
}

type KeyNode[K Comparable] struct {
	height int
	key    K
	left   *KeyNode[K]
	right  *KeyNode[K]
	parent *KeyNode[K]
}

func (node *KeyNode[K]) Height() int {
	return node.height
}

func (node *KeyNode[K]) Key() K {
	return node.key
}

func (node *KeyNode[K]) Value() K {
	return node.key
}

func (node *KeyNode[K]) Left() *KeyNode[K] {
	return node.left
}

func (node *KeyNode[K]) Right() *KeyNode[K] {
	return node.right
}

func (node *KeyNode[K]) Parent() *KeyNode[K] {
	return node.parent
}

type KeyValueNode[K Comparable, V any] struct {
	height int
	key    K
	value  V
	left   *KeyValueNode[K, V]
	right  *KeyValueNode[K, V]
	parent *KeyValueNode[K, V]
}

func (node *KeyValueNode[K, V]) Height() int {
	return node.height
}

func (node *KeyValueNode[K, V]) Key() K {
	return node.key
}

func (node *KeyValueNode[K, V]) Value() V {
	return node.value
}

func (node *KeyValueNode[K, V]) Left() *KeyValueNode[K, V] {
	return node.left
}

func (node *KeyValueNode[K, V]) Right() *KeyValueNode[K, V] {
	return node.right
}

func (node *KeyValueNode[K, V]) Parent() *KeyValueNode[K, V] {
	return node.parent
}
