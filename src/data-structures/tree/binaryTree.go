package binaryTree

type Comparison = int

const (
	BIGGER  Comparison = 1
	EQUAL              = 0
	SMALLER            = -1
)

type CompareFn[T any] func(a T, b T) Comparison

type TreeNode[T any] struct {
	left  *TreeNode[T]
	right *TreeNode[T]
	value T
}

type BinaryTree[T any] struct {
	root    *TreeNode[T]
	size    int
	compare CompareFn[T]
}

func CreateBinaryTree[T any](compare CompareFn[T]) BinaryTree[T] {
	return BinaryTree[T]{root: nil, size: 0, compare: compare}
}

func FindInsertNodePtr[T any](root *TreeNode[T], value T, compare CompareFn[T]) {
	comparison := compare(value, root.value)

	if comparison == SMALLER {
		if root.left == nil {
			root.left = &TreeNode[T]{left: nil, right: nil, value: value}
		} else {
			FindInsertNodePtr(root.left, value, compare)
		}
	} else {
		if root.right == nil {
			root.right = &TreeNode[T]{left: nil, right: nil, value: value}
		} else {
			FindInsertNodePtr(root.right, value, compare)
		}
	}
}

func (tree *BinaryTree[T]) Push(value T) {
	newNodePtr := &TreeNode[T]{left: nil, right: nil, value: value}

	if tree.root == nil {
		tree.root = newNodePtr
	} else {
		FindInsertNodePtr(tree.root, value, tree.compare)
	}

	tree.size++
}

func (tree *BinaryTree[T]) ToArray() []T {
	array := make([]T, 0)

	tree.TraverseInPreOrder(func(value T) {
		array = append(array, value)
	})

	return array
}

func (tree *BinaryTree[T]) TraverseInPreOrder(callback func(value T)) {
	TraverseInPreOrderHelper(tree.root, callback)
}

func (tree *BinaryTree[T]) TraverseInOrder(callback func(value T)) {
	TraverseInOrderHelper(tree.root, callback)
}

func (tree *BinaryTree[T]) TraverseInPostOrder(callback func(value T)) {
	TraverseInPostOrderHelper(tree.root, callback)
}

func TraverseInPreOrderHelper[T any](node *TreeNode[T], callback func(value T)) {
	if node != nil {
		callback(node.value)
		TraverseInPreOrderHelper(node.left, callback)
		TraverseInPreOrderHelper(node.right, callback)
	}
}

func TraverseInOrderHelper[T any](node *TreeNode[T], callback func(value T)) {
	if node != nil {
		TraverseInPreOrderHelper(node.left, callback)
		callback(node.value)
		TraverseInPreOrderHelper(node.right, callback)
	}
}

func TraverseInPostOrderHelper[T any](node *TreeNode[T], callback func(value T)) {
	if node != nil {
		TraverseInPreOrderHelper(node.left, callback)
		TraverseInPreOrderHelper(node.right, callback)
		callback(node.value)
	}
}
