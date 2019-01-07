package trees

import (
	"reflect"
	"testing"
)

func TestFindMaxValue(t *testing.T) {
	tree := getTree()

	var result int
	var expected = 7

	result = FindMaxValue(tree.Root)

	if result != expected {
		t.Errorf("Value is incorrect, got %v instead of %v", result, expected)
	}
}

func TestFindMaxValueNonRecursive(t *testing.T) {
	tree := getTree()

	var result int
	var expected = 7

	result = FindMaxValueNonRecursive(tree.Root)

	if result != expected {
		t.Errorf("Value is incorrect, got %v instead of %v", result, expected)
	}
}

func TestContains(t *testing.T) {
	tree := getTree()

	var result bool

	result = Contains(tree.Root, 5)

	if !result {
		t.Errorf("Value is incorrect, got %v instead of %v", result, true)
	}

	result = Contains(tree.Root, 55)

	if result {
		t.Errorf("Value is incorrect, got %v instead of %v", result, false)
	}
}

func TestSize(t *testing.T) {
	tree := getTree()

	var result int
	var expected = 7

	result = Size(tree.Root)

	if result != expected {
		t.Errorf("Value is incorrect, got %v instead of %v", result, expected)
	}
}

func TestSizeNonRecursive(t *testing.T) {
	tree := getTree()

	var result int
	var expected = 7

	result = SizeNonRecursive(tree.Root)

	if result != expected {
		t.Errorf("Value is incorrect, got %v instead of %v", result, expected)
	}
}

func TestContainsNonRecursive(t *testing.T) {
	tree := getTree()

	var result bool

	result = ContainsNonRecursive(tree.Root, 5)

	if !result {
		t.Errorf("Value is incorrect, got %v instead of %v", result, true)
	}

	result = Contains(tree.Root, 55)

	if result {
		t.Errorf("Value is incorrect, got %v instead of %v", result, false)
	}
}

func TestDepth(t *testing.T) {
	tree := &BinaryTree{&BinaryTreeNode{Value: 1, Key: 100}}
	tree.Insert(2, 90)
	tree.Insert(3, 110)
	tree.Insert(4, 80)
	tree.Insert(5, 95)

	var result int
	var expected = 3

	result = Depth(tree.Root)

	if result != expected {
		t.Errorf("Value is incorrect, got %v instead of %v", result, expected)
	}
}

func TestLevelOrderTraversalInReverse(t *testing.T) {
	tree := getTree()

	var result []int
	var expected = []int{4, 5, 6, 7, 2, 3, 1}

	LevelOrderTraversalInReverse(tree.Root, func(value int) {
		result = append(result, value)
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Traversal order incorrect, got %v instead of %v", result, expected)
	}
}

func TestDeleteTree(t *testing.T) {
	tree := getTree()
	tree = DeleteTree()

	if tree != nil {
		t.Errorf("Delete tree is incorrect, got %v instead of %v", tree, nil)
	}
}