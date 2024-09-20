package binarytree

type TreeNode[K comparable] struct {
	Left  *TreeNode[K]
	Right *TreeNode[K]
	Val   K
}

type Tree[K comparable] struct {
	root *TreeNode[K]
}

func New[K comparable](rootVal K) *Tree[K] {
	return &Tree[K]{root: &TreeNode[K]{Val: rootVal}}
}

func (node *Tree[K]) Contains(key K) bool {
	return contains(node.root, key)
}

func (node *Tree[K]) Counter() int {
	return counter(node.root)
}

func contains[K comparable](root *TreeNode[K], key K) bool {
	if root == nil {
		return false
	}

	if root.Val == key {
		return true
	}

	leftContain := contains(root.Left, key)
	rightContain := contains(root.Right, key)
	return leftContain || rightContain
}

func counter[K comparable](root *TreeNode[K]) int {
	if root == nil {
		return 0
	}

	return counter(root.Left) + counter(root.Right) + 1
}
