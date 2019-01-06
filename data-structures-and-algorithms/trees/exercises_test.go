package trees

import (
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