package trees

import (
	"reflect"
	"testing"
)

func getTree() *BinaryTree {
	tree := &BinaryTree{&BinaryTreeNode{Value: 1, Key: 100}}
	tree.Insert(2, 90)
	tree.Insert(3, 110)
	tree.Insert(4, 80)
	tree.Insert(5, 95)
	tree.Insert(6, 105)
	tree.Insert(7, 120)

	return tree
}

func TestPreOrder(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{1, 2, 4, 5, 3, 6, 7}

	PreOrder(tree.Root, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestPreOrderNonRecursive(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{1, 2, 4, 5, 3, 6, 7}

	PreOrderNonRecursive(tree, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestInOrder(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{4, 2, 5, 1, 6, 3, 7}

	InOrder(tree.Root, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestInOrderNonRecursive(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{4, 2, 5, 1, 6, 3, 7}

	InOrderNonRecursive(tree, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestPostOrder(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{4, 5, 2, 6, 7, 3, 1}

	PostOrder(tree.Root, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestPostOrderNonRecursive(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{4, 5, 2, 6, 7, 3, 1}

	PostOrderNonRecursive(tree, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}