package ds

import (
	"reflect"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := NewBST()

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(9)

	if bst.Size() != 5 {
		t.Errorf("expected size 5, got %d", bst.Size())
	}

	// Inorder should give sorted order
	expected := []int{1, 3, 5, 7, 9}
	if !reflect.DeepEqual(bst.InorderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bst.InorderTraversal())
	}
}

func TestBST_InsertDuplicates(t *testing.T) {
	bst := NewBST()

	bst.Insert(5)
	bst.Insert(5)
	bst.Insert(5)

	if bst.Size() != 1 {
		t.Errorf("duplicates should not be inserted, expected size 1, got %d", bst.Size())
	}
}

func TestBST_Search(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	if !bst.Search(7) {
		t.Error("should find value 7")
	}

	if !bst.Search(10) {
		t.Error("should find value 10 (root)")
	}

	if bst.Search(99) {
		t.Error("should not find value 99")
	}
}

func TestBST_SearchEmpty(t *testing.T) {
	bst := NewBST()

	if bst.Search(1) {
		t.Error("should not find value in empty BST")
	}
}

func TestBST_DeleteLeaf(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	if !bst.Delete(5) {
		t.Error("delete should succeed")
	}

	if bst.Search(5) {
		t.Error("value 5 should be deleted")
	}

	if bst.Size() != 2 {
		t.Errorf("expected size 2, got %d", bst.Size())
	}
}

func TestBST_DeleteNodeWithOneChild(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(3)

	if !bst.Delete(5) {
		t.Error("delete should succeed")
	}

	if bst.Search(5) {
		t.Error("value 5 should be deleted")
	}

	if !bst.Search(3) {
		t.Error("value 3 should still exist")
	}
}

func TestBST_DeleteNodeWithTwoChildren(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(12)
	bst.Insert(17)

	// Delete node with two children
	if !bst.Delete(10) {
		t.Error("delete should succeed")
	}

	if bst.Search(10) {
		t.Error("value 10 should be deleted")
	}

	// BST property should be maintained
	if !bst.IsValidBST() {
		t.Error("BST property should be maintained after deletion")
	}

	// All other values should still exist
	values := []int{3, 5, 7, 12, 15, 17}
	for _, v := range values {
		if !bst.Search(v) {
			t.Errorf("value %d should still exist", v)
		}
	}
}

func TestBST_DeleteRoot(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)

	if !bst.Delete(10) {
		t.Error("delete should succeed")
	}

	if !bst.IsEmpty() {
		t.Error("BST should be empty after deleting only node")
	}
}

func TestBST_DeleteNonExistent(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)

	if bst.Delete(99) {
		t.Error("delete of non-existent value should fail")
	}

	if bst.Size() != 2 {
		t.Errorf("size should remain 2, got %d", bst.Size())
	}
}

func TestBST_FindMin(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	min, ok := bst.FindMin()
	if !ok {
		t.Error("should find min value")
	}
	if min != 3 {
		t.Errorf("expected min 3, got %d", min)
	}
}

func TestBST_FindMinEmpty(t *testing.T) {
	bst := NewBST()

	_, ok := bst.FindMin()
	if ok {
		t.Error("should not find min in empty BST")
	}
}

func TestBST_FindMax(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(20)

	max, ok := bst.FindMax()
	if !ok {
		t.Error("should find max value")
	}
	if max != 20 {
		t.Errorf("expected max 20, got %d", max)
	}
}

func TestBST_FindMaxEmpty(t *testing.T) {
	bst := NewBST()

	_, ok := bst.FindMax()
	if ok {
		t.Error("should not find max in empty BST")
	}
}

func TestBST_InorderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range values {
		bst.Insert(v)
	}

	// Inorder should give sorted order
	expected := []int{20, 30, 40, 50, 60, 70, 80}
	if !reflect.DeepEqual(bst.InorderTraversal(), expected) {
		t.Errorf("expected %v, got %v", expected, bst.InorderTraversal())
	}
}

func TestBST_Height(t *testing.T) {
	bst := NewBST()

	if bst.Height() != -1 {
		t.Errorf("expected height -1 for empty BST, got %d", bst.Height())
	}

	bst.Insert(10)
	if bst.Height() != 0 {
		t.Errorf("expected height 0 for single node, got %d", bst.Height())
	}

	bst.Insert(5)
	bst.Insert(15)
	if bst.Height() != 1 {
		t.Errorf("expected height 1, got %d", bst.Height())
	}
}

func TestBST_IsEmpty(t *testing.T) {
	bst := NewBST()

	if !bst.IsEmpty() {
		t.Error("new BST should be empty")
	}

	bst.Insert(1)
	if bst.IsEmpty() {
		t.Error("BST with node should not be empty")
	}
}

func TestBST_Clear(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	bst.Clear()

	if !bst.IsEmpty() {
		t.Error("BST should be empty after clear")
	}

	if bst.Size() != 0 {
		t.Errorf("expected size 0 after clear, got %d", bst.Size())
	}
}

func TestBST_IsValidBST(t *testing.T) {
	bst := NewBST()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	if !bst.IsValidBST() {
		t.Error("BST should be valid")
	}
}

func TestBST_IsValidBSTInvalid(t *testing.T) {
	bst := NewBST()
	bst.Root = NewTreeNode(10)
	bst.Root.Left = NewTreeNode(5)
	bst.Root.Right = NewTreeNode(15)
	// Violate BST property: left subtree has value > root
	bst.Root.Left.Right = NewTreeNode(20)

	if bst.IsValidBST() {
		t.Error("BST should be invalid")
	}
}

func TestBST_ComplexOperations(t *testing.T) {
	bst := NewBST()

	// Insert values
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, v := range values {
		bst.Insert(v)
	}

	// Verify size
	if bst.Size() != len(values) {
		t.Errorf("expected size %d, got %d", len(values), bst.Size())
	}

	// Delete some values
	bst.Delete(20)
	bst.Delete(70)

	// Verify remaining values
	if bst.Search(20) || bst.Search(70) {
		t.Error("deleted values should not be found")
	}

	// Verify BST property maintained
	if !bst.IsValidBST() {
		t.Error("BST property should be maintained")
	}

	// Verify inorder is still sorted
	inorder := bst.InorderTraversal()
	for i := 1; i < len(inorder); i++ {
		if inorder[i] <= inorder[i-1] {
			t.Error("inorder traversal should be sorted")
		}
	}
}
