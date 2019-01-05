package trees

// Problem 1
// Give an algorithm for finding maximum element in binary tree.

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