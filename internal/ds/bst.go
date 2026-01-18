package ds

// Why interviewers ask this:
// BST is crucial for understanding ordered data structures and efficient search operations.
// It demonstrates understanding of binary tree properties, recursion, and the trade-offs
// between different operations. BST operations are frequently asked in coding interviews.

// Common pitfalls:
// - Not maintaining BST property during insertion/deletion
// - Forgetting to handle duplicate values
// - Incorrect deletion logic (especially for nodes with two children)
// - Not understanding that BST degrades to O(n) when unbalanced
// - Confusing inorder traversal with sorted order

// Key takeaway:
// BST maintains left < root < right property. Inorder traversal gives sorted order.
// Average case O(log n) for search/insert/delete, worst case O(n) for skewed tree.
// Deletion with two children: replace with inorder successor or predecessor.

// BST represents a Binary Search Tree
// Time Complexity: Average O(log n), Worst O(n) for search/insert/delete
// Space Complexity: O(n) for n nodes
type BST struct {
	Root *TreeNode
}

// NewBST creates a new empty binary search tree
func NewBST() *BST {
	return &BST{Root: nil}
}

// Insert adds a value to the BST maintaining BST property
// Duplicates are not inserted
// Time Complexity: O(log n) average, O(n) worst case
func (bst *BST) Insert(value int) {
	bst.Root = bst.insertHelper(bst.Root, value)
}

func (bst *BST) insertHelper(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return NewTreeNode(value)
	}

	if value < node.Value {
		node.Left = bst.insertHelper(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertHelper(node.Right, value)
	}
	// If value == node.Value, don't insert (no duplicates)

	return node
}

// Search checks if a value exists in the BST
// Time Complexity: O(log n) average, O(n) worst case
func (bst *BST) Search(value int) bool {
	return bst.searchHelper(bst.Root, value)
}

func (bst *BST) searchHelper(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	} else if value < node.Value {
		return bst.searchHelper(node.Left, value)
	} else {
		return bst.searchHelper(node.Right, value)
	}
}

// Delete removes a value from the BST
// Returns true if value was found and deleted
// Time Complexity: O(log n) average, O(n) worst case
func (bst *BST) Delete(value int) bool {
	if !bst.Search(value) {
		return false
	}
	bst.Root = bst.deleteHelper(bst.Root, value)
	return true
}

func (bst *BST) deleteHelper(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = bst.deleteHelper(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.deleteHelper(node.Right, value)
	} else {
		// Found the node to delete

		// Case 1: No children (leaf node)
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// Case 2: One child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Case 3: Two children
		// Find inorder successor (minimum in right subtree)
		successor := bst.findMin(node.Right)
		node.Value = successor.Value
		node.Right = bst.deleteHelper(node.Right, successor.Value)
	}

	return node
}

// FindMin returns the minimum value in the BST
// Returns 0 and false if tree is empty
// Time Complexity: O(log n) average, O(n) worst case
func (bst *BST) FindMin() (int, bool) {
	if bst.Root == nil {
		return 0, false
	}

	return bst.findMin(bst.Root).Value, true
}

func (bst *BST) findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// FindMax returns the maximum value in the BST
// Returns 0 and false if tree is empty
// Time Complexity: O(log n) average, O(n) worst case
func (bst *BST) FindMax() (int, bool) {
	if bst.Root == nil {
		return 0, false
	}

	return bst.findMax(bst.Root).Value, true
}

func (bst *BST) findMax(node *TreeNode) *TreeNode {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// InorderTraversal returns values in sorted order
// Time Complexity: O(n)
func (bst *BST) InorderTraversal() []int {
	result := []int{}
	bst.inorderHelper(bst.Root, &result)
	return result
}

func (bst *BST) inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	bst.inorderHelper(node.Left, result)
	*result = append(*result, node.Value)
	bst.inorderHelper(node.Right, result)
}

// Height returns the height of the BST
// Time Complexity: O(n)
func (bst *BST) Height() int {
	return bst.heightHelper(bst.Root)
}

func (bst *BST) heightHelper(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bst.heightHelper(node.Left)
	rightHeight := bst.heightHelper(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Size returns the total number of nodes
// Time Complexity: O(n)
func (bst *BST) Size() int {
	return bst.sizeHelper(bst.Root)
}

func (bst *BST) sizeHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + bst.sizeHelper(node.Left) + bst.sizeHelper(node.Right)
}

// IsEmpty returns true if BST has no nodes
func (bst *BST) IsEmpty() bool {
	return bst.Root == nil
}

// Clear removes all nodes from the BST
func (bst *BST) Clear() {
	bst.Root = nil
}

// IsValidBST checks if the tree maintains BST property
// Time Complexity: O(n)
func (bst *BST) IsValidBST() bool {
	return bst.isValidBSTHelper(bst.Root, nil, nil)
}

func (bst *BST) isValidBSTHelper(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}

	if min != nil && node.Value <= *min {
		return false
	}
	if max != nil && node.Value >= *max {
		return false
	}

	return bst.isValidBSTHelper(node.Left, min, &node.Value) &&
		bst.isValidBSTHelper(node.Right, &node.Value, max)
}
