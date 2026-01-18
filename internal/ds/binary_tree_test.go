package ds

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Insert(t *testing.T) {
	bt := NewBinaryTree()

	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)
	bt.Insert(4)

	if bt.Size() != 4 {
		t.Errorf("expected size 4, got %d", bt.Size())
	}

	// Level-order should be [1, 2, 3, 4]
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(bt.LevelOrderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bt.LevelOrderTraversal())
	}
}

func TestBinaryTree_InorderTraversal(t *testing.T) {
	bt := NewBinaryTree()
	bt.Root = NewTreeNode(2)
	bt.Root.Left = NewTreeNode(1)
	bt.Root.Right = NewTreeNode(3)

	// Inorder: Left-Root-Right = [1, 2, 3]
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(bt.InorderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bt.InorderTraversal())
	}
}

func TestBinaryTree_PreorderTraversal(t *testing.T) {
	bt := NewBinaryTree()
	bt.Root = NewTreeNode(2)
	bt.Root.Left = NewTreeNode(1)
	bt.Root.Right = NewTreeNode(3)

	// Preorder: Root-Left-Right = [2, 1, 3]
	expected := []int{2, 1, 3}
	if !reflect.DeepEqual(bt.PreorderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bt.PreorderTraversal())
	}
}

func TestBinaryTree_PostorderTraversal(t *testing.T) {
	bt := NewBinaryTree()
	bt.Root = NewTreeNode(2)
	bt.Root.Left = NewTreeNode(1)
	bt.Root.Right = NewTreeNode(3)

	// Postorder: Left-Right-Root = [1, 3, 2]
	expected := []int{1, 3, 2}
	if !reflect.DeepEqual(bt.PostorderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bt.PostorderTraversal())
	}
}

func TestBinaryTree_LevelOrderTraversal(t *testing.T) {
	bt := NewBinaryTree()
	bt.Root = NewTreeNode(1)
	bt.Root.Left = NewTreeNode(2)
	bt.Root.Right = NewTreeNode(3)
	bt.Root.Left.Left = NewTreeNode(4)
	bt.Root.Left.Right = NewTreeNode(5)

	// Level-order: [1, 2, 3, 4, 5]
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(bt.LevelOrderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bt.LevelOrderTraversal())
	}
}

func TestBinaryTree_Height(t *testing.T) {
	bt := NewBinaryTree()

	// Empty tree height is -1
	if bt.Height() != -1 {
		t.Errorf("expected height -1 for empty tree, got %d", bt.Height())
	}

	// Single node height is 0
	bt.Insert(1)
	if bt.Height() != 0 {
		t.Errorf("expected height 0 for single node, got %d", bt.Height())
	}

	// Add more nodes
	bt.Insert(2)
	bt.Insert(3)
	if bt.Height() != 1 {
		t.Errorf("expected height 1, got %d", bt.Height())
	}

	bt.Insert(4)
	if bt.Height() != 2 {
		t.Errorf("expected height 2, got %d", bt.Height())
	}
}

func TestBinaryTree_Size(t *testing.T) {
	bt := NewBinaryTree()

	if bt.Size() != 0 {
		t.Errorf("expected size 0 for empty tree, got %d", bt.Size())
	}

	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)

	if bt.Size() != 3 {
		t.Errorf("expected size 3, got %d", bt.Size())
	}
}

func TestBinaryTree_Search(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(30)
	bt.Insert(40)

	if !bt.Search(20) {
		t.Error("should find value 20")
	}

	if !bt.Search(40) {
		t.Error("should find value 40")
	}

	if bt.Search(99) {
		t.Error("should not find value 99")
	}
}

func TestBinaryTree_SearchEmpty(t *testing.T) {
	bt := NewBinaryTree()

	if bt.Search(1) {
		t.Error("should not find value in empty tree")
	}
}

func TestBinaryTree_IsEmpty(t *testing.T) {
	bt := NewBinaryTree()

	if !bt.IsEmpty() {
		t.Error("new tree should be empty")
	}

	bt.Insert(1)
	if bt.IsEmpty() {
		t.Error("tree with node should not be empty")
	}
}

func TestBinaryTree_Clear(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(1)
	bt.Insert(2)
	bt.Insert(3)

	bt.Clear()

	if !bt.IsEmpty() {
		t.Error("tree should be empty after clear")
	}

	if bt.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", bt.Size())
	}
}

func TestBinaryTree_MaxValue(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(50)
	bt.Insert(30)
	bt.Insert(20)

	max, ok := bt.MaxValue()
	if !ok {
		t.Error("should find max value")
	}
	if max != 50 {
		t.Errorf("expected max 50, got %d", max)
	}
}

func TestBinaryTree_MaxValueEmpty(t *testing.T) {
	bt := NewBinaryTree()

	_, ok := bt.MaxValue()
	if ok {
		t.Error("should not find max in empty tree")
	}
}

func TestBinaryTree_MinValue(t *testing.T) {
	bt := NewBinaryTree()
	bt.Insert(10)
	bt.Insert(50)
	bt.Insert(30)
	bt.Insert(5)

	min, ok := bt.MinValue()
	if !ok {
		t.Error("should find min value")
	}
	if min != 5 {
		t.Errorf("expected min 5, got %d", min)
	}
}

func TestBinaryTree_MinValueEmpty(t *testing.T) {
	bt := NewBinaryTree()

	_, ok := bt.MinValue()
	if ok {
		t.Error("should not find min in empty tree")
	}
}

func TestBinaryTree_TraversalEmpty(t *testing.T) {
	bt := NewBinaryTree()

	if len(bt.InorderTraversal()) != 0 {
		t.Error("inorder of empty tree should be empty")
	}

	if len(bt.PreorderTraversal()) != 0 {
		t.Error("preorder of empty tree should be empty")
	}

	if len(bt.PostorderTraversal()) != 0 {
		t.Error("postorder of empty tree should be empty")
	}

	if len(bt.LevelOrderTraversal()) != 0 {
		t.Error("level-order of empty tree should be empty")
	}
}

func TestBinaryTree_ComplexTree(t *testing.T) {
	bt := NewBinaryTree()
	// Build tree:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	bt.Root = NewTreeNode(1)
	bt.Root.Left = NewTreeNode(2)
	bt.Root.Right = NewTreeNode(3)
	bt.Root.Left.Left = NewTreeNode(4)
	bt.Root.Left.Right = NewTreeNode(5)

	// Verify traversals
	inorder := []int{4, 2, 5, 1, 3}
	if !reflect.DeepEqual(bt.InorderTraversal(), inorder) {
		t.Errorf("expected inorder %v, got %v", inorder, bt.InorderTraversal())
	}

	preorder := []int{1, 2, 4, 5, 3}
	if !reflect.DeepEqual(bt.PreorderTraversal(), preorder) {
		t.Errorf("expected preorder %v, got %v", preorder, bt.PreorderTraversal())
	}

	postorder := []int{4, 5, 2, 3, 1}
	if !reflect.DeepEqual(bt.PostorderTraversal(), postorder) {
		t.Errorf("expected postorder %v, got %v", postorder, bt.PostorderTraversal())
	}

	levelorder := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(bt.LevelOrderTraversal(), levelorder) {
		t.Errorf("expected level-order %v, got %v", levelorder, bt.LevelOrderTraversal())
	}

	if bt.Height() != 2 {
		t.Errorf("expected height 2, got %d", bt.Height())
	}

	if bt.Size() != 5 {
		t.Errorf("expected size 5, got %d", bt.Size())
	}
}
