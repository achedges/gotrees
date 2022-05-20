package gotrees

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
