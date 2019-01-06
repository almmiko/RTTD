package trees

type BinaryTreeNode struct {
	Value int
	Key   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}