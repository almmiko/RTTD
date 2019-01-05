package trees

func PreOrder(node *BinaryTreeNode, cb func(value int)) {
	if node != nil {
		cb(node.Value)
		PreOrder(node.Left, cb)
		PreOrder(node.Right, cb)
	}
}

func PreOrderNonRecursive(tree *BinaryTree, cb func(value int)) {
	var stack []*BinaryTreeNode

	for {
		for tree.Root != nil {
			stack = append(stack, tree.Root)
			cb(tree.Root.Value)
			tree.Root = tree.Root.Left
		}

		if len(stack) == 0 {
			break
		}

		tree.Root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		tree.Root = tree.Root.Right
	}
}

func InOrder(node *BinaryTreeNode, cb func(value int)) {
	if node != nil {
		InOrder(node.Left, cb)
		cb(node.Value)
		InOrder(node.Right, cb)
	}
}

func InOrderNonRecursive(tree *BinaryTree, cb func(value int)) {
	var stack []*BinaryTreeNode

	for {
		for tree.Root != nil {
			stack = append(stack, tree.Root)
			tree.Root = tree.Root.Left
		}

		if len(stack) == 0 {
			break
		}

		tree.Root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		cb(tree.Root.Value)
		tree.Root = tree.Root.Right
	}
}

func PostOrder(node *BinaryTreeNode, cb func(value int)) {
	if node != nil {
		PostOrder(node.Left, cb)
		PostOrder(node.Right, cb)
		cb(node.Value)
	}
}

func PostOrderNonRecursive(tree *BinaryTree, cb func(value int)) {
	var stack []*BinaryTreeNode
	var prevNode *BinaryTreeNode

	for {
		for tree.Root != nil {
			stack = append(stack, tree.Root)
			tree.Root = tree.Root.Left
		}

		if len(stack) == 0 {
			break
		}

		for tree.Root == nil && len(stack) != 0 {
			tree.Root = stack[len(stack) - 1]

			if tree.Root.Right == nil || tree.Root.Right == prevNode {
				cb(tree.Root.Value)
				stack = stack[:len(stack)-1]
				prevNode = tree.Root
				tree.Root = nil
			} else {
				tree.Root = tree.Root.Right
			}
		}
	}
}