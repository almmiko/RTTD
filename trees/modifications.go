package trees

func (tree *BinaryTree) Insert(value, key int) {
	node := &BinaryTreeNode{Value: value, Key: key}

	if tree.Root == nil {
		tree.Root = node
	} else {
		InsertNode(tree.Root, node)
	}
}

func InsertNode(root *BinaryTreeNode, node *BinaryTreeNode) {
	if node.Key < root.Key {
		if root.Left == nil {
			root.Left = node
		} else {
			InsertNode(root.Left, node)
		}
	} else {
		if root.Right == nil {
			root.Right = node
		} else {
			InsertNode(root.Right, node)
		}
	}
}
