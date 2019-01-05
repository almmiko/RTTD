package trees

// Problem-1 Give an algorithm for finding maximum element in binary tree.
func FindMaxValue(root *BinaryTreeNode) int {
	var maxValue int

	if root != nil {
		rootValue := root.Value

		leftValue := FindMaxValue(root.Left)
		rightValue := FindMaxValue(root.Right)

		if leftValue > rightValue {
			maxValue = leftValue
		} else {
			maxValue = rightValue
		}

		if rootValue > maxValue {
			maxValue = rootValue
		}
	}

	return maxValue
}

// Problem-2 Give an algorithm for finding the maximum element in binary tree without recursion.
func FindMaxValueNonRecursive(node *BinaryTreeNode) int {
	var tempNode *BinaryTreeNode
	var queue []*BinaryTreeNode
	var maxValue int

	if node == nil {
		return 0
	}

	queue = append(queue, node)

	for len(queue) != 0 {
		tempNode, queue = queue[0], queue[1:]

		if maxValue < tempNode.Value {
			maxValue = tempNode.Value
		}

		if tempNode.Left != nil {
			queue = append(queue, tempNode.Left)
		}

		if tempNode.Right != nil {
			queue = append(queue, tempNode.Right)
		}

	}

	return maxValue
}

// Problem-3 Give an algorithm for searching an element in binary tree.
func Contains(root *BinaryTreeNode, data int) bool {
	if root == nil {
		return false
	}

	if root.Value == data {
		return true
	} else {
		valueLeft := Contains(root.Left, data)

		if valueLeft {
			return valueLeft
		} else {
			return 	Contains(root.Right, data)
		}

	}
}

// Problem-4 Give an algorithm for searching an element in binary tree without recursion.

// Problem-6 Give an algorithm for finding the size of binary tree.

// Problem-7 Can we solve Problem-6 without recursion?

// Problem-8 Give an algorithm for printing the level order data in reverse order. For example,
// the output for the below tree should be: 4 5 6 7 2 3 1

// Problem-9 Give an algorithm for deleting the tree.

// Problem-10 Give an algorithm for finding the height (or depth) of the binary tree.

// Problem-11 Can we solve Problem-10 without recursion?

// Problem-12 Give an algorithm for finding the deepest node of the binary tree.

// Problem-13 Give an algorithm for deleting an element (assuming data is given) from binary
// tree.

// Problem-14 Give an algorithm for finding the number of leaves in the binary tree without
// using recursion.

// Problem-15 Give an algorithm for finding the number of full nodes in the binary tree without
// using recursion.

// Problem-16 Give an algorithm for finding the number of half nodes (nodes with only one
// child) in the binary tree without using recursion.

// Problem-17 Given two binary trees, return true if they are structurally identical.

// Problem-18 Give an algorithm for finding the diameter of the binary tree. The diameter of a
// tree (sometimes called the width) is the number of nodes on the longest path between two
// leaves in the tree.

// Problem-19 Give an algorithm for finding the level that has the maximum sum in the binary
// tree.

// Problem-20 Given a binary tree, print out all its root-to-leaf paths.

// Problem-21 Give an algorithm for checking the existence of path with given sum. That
// means, given a sum, check whether there exists a path from root to any of the nodes.

// Problem-22 Give an algorithm for finding the sum of all elements in binary tree.

// Problem-23 Can we solve Problem-22 without recursion?

// Problem-24 Give an algorithm for converting a tree to its mirror. Mirror of a tree is another
// tree with left and right children of all non-leaf nodes interchanged. The trees below are
// mirrors to each other.

// Problem-25 Given two trees, give an algorithm for checking whether they are mirrors of
// each other.

// Problem-26 Give an algorithm for finding LCA (Least Common Ancestor) of two nodes in a
// Binary Tree.

// Problem-27 Give an algorithm for constructing binary tree from given Inorder and Preorder
// traversals.

// Problem-28 If we are given two traversal sequences, can we construct the binary tree
// uniquely?

// Problem-29 Give an algorithm for printing all the ancestors of a node in a Binary tree. For
// the tree below, for 7 the ancestors are 1 3 7.

// Problem-30 Zigzag Tree Traversal: Give an algorithm to traverse a binary tree in Zigzag
// order. For example, the output for the tree below should be: 1 3 2 4 5 6 7

// Problem-31 Give an algorithm for finding the vertical sum of a binary tree. For example, The
// tree has 5 vertical lines
// Vertical-1: nodes-4 => vertical sum is 4
// Vertical-2: nodes-2 => vertical sum is 2
// Vertical-3: nodes-1,5,6 => vertical sum is 1 + 5 + 6 = 12
// Vertical-4: nodes-3 => vertical sum is 3
// Vertical-5: nodes-7 => vertical sum is 7
// We need to output: 4 2 12 3 7

// Problem-33 Given a tree with a special property where leaves are represented with ‘L’ and
// internal node with ‘I’. Also, assume that each node has either 0 or 2 children. Given
// preorder traversal of this tree, construct the tree.
// Example: Given preorder string => ILILL

// Problem-34 Given a binary tree with three pointers (left, right and nextSibling), give an
// algorithm for filling the nextSibling pointers assuming they are NULL initially.

// Problem-35 Is there any other way of solving Problem-34?