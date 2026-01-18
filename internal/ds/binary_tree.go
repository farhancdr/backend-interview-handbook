package ds

// Why interviewers ask this:
// Binary trees are fundamental to understanding hierarchical data structures and tree traversal
// algorithms. They're the basis for BSTs, heaps, and many other structures. Interviewers test
// understanding of recursion, tree traversal (inorder, preorder, postorder, level-order), and
// tree properties (height, depth, balance).

// Common pitfalls:
// - Confusing left and right subtrees during traversal
// - Not handling nil nodes properly in recursive functions
// - Off-by-one errors in height/depth calculations
// - Forgetting base cases in recursive implementations
// - Not understanding the difference between tree traversal orders

// Key takeaway:
// Binary tree has at most two children per node. Master the three DFS traversals (inorder,
// preorder, postorder) and BFS (level-order). Recursion is natural for trees. Always handle
// nil nodes as base case.

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents a binary tree structure
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree creates a new empty binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// NewTreeNode creates a new tree node with given value
func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value, Left: nil, Right: nil}
}

// Insert adds a value to the tree using level-order insertion
// This maintains a complete binary tree property
// Time Complexity: O(n)
func (bt *BinaryTree) Insert(value int) {
	newNode := NewTreeNode(value)

	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	// Level-order insertion using queue
	queue := []*TreeNode{bt.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = newNode
			return
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = newNode
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}

// InorderTraversal returns values in inorder (Left-Root-Right)
// Time Complexity: O(n), Space Complexity: O(h) where h is height
func (bt *BinaryTree) InorderTraversal() []int {
	result := []int{}
	bt.inorderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) inorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	bt.inorderHelper(node.Left, result)
	*result = append(*result, node.Value)
	bt.inorderHelper(node.Right, result)
}

// PreorderTraversal returns values in preorder (Root-Left-Right)
// Time Complexity: O(n), Space Complexity: O(h)
func (bt *BinaryTree) PreorderTraversal() []int {
	result := []int{}
	bt.preorderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) preorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Value)
	bt.preorderHelper(node.Left, result)
	bt.preorderHelper(node.Right, result)
}

// PostorderTraversal returns values in postorder (Left-Right-Root)
// Time Complexity: O(n), Space Complexity: O(h)
func (bt *BinaryTree) PostorderTraversal() []int {
	result := []int{}
	bt.postorderHelper(bt.Root, &result)
	return result
}

func (bt *BinaryTree) postorderHelper(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	bt.postorderHelper(node.Left, result)
	bt.postorderHelper(node.Right, result)
	*result = append(*result, node.Value)
}

// LevelOrderTraversal returns values in level-order (BFS)
// Time Complexity: O(n), Space Complexity: O(w) where w is max width
func (bt *BinaryTree) LevelOrderTraversal() []int {
	result := []int{}

	if bt.Root == nil {
		return result
	}

	queue := []*TreeNode{bt.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}

// Height returns the height of the tree (longest path from root to leaf)
// Height of empty tree is -1, single node is 0
// Time Complexity: O(n)
func (bt *BinaryTree) Height() int {
	return bt.heightHelper(bt.Root)
}

func (bt *BinaryTree) heightHelper(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bt.heightHelper(node.Left)
	rightHeight := bt.heightHelper(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Size returns the total number of nodes in the tree
// Time Complexity: O(n)
func (bt *BinaryTree) Size() int {
	return bt.sizeHelper(bt.Root)
}

func (bt *BinaryTree) sizeHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + bt.sizeHelper(node.Left) + bt.sizeHelper(node.Right)
}

// Search checks if a value exists in the tree
// Time Complexity: O(n)
func (bt *BinaryTree) Search(value int) bool {
	return bt.searchHelper(bt.Root, value)
}

func (bt *BinaryTree) searchHelper(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if node.Value == value {
		return true
	}

	return bt.searchHelper(node.Left, value) || bt.searchHelper(node.Right, value)
}

// IsEmpty returns true if tree has no nodes
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root == nil
}

// Clear removes all nodes from the tree
func (bt *BinaryTree) Clear() {
	bt.Root = nil
}

// MaxValue returns the maximum value in the tree
// Returns 0 and false if tree is empty
// Time Complexity: O(n)
func (bt *BinaryTree) MaxValue() (int, bool) {
	if bt.Root == nil {
		return 0, false
	}

	return bt.maxValueHelper(bt.Root), true
}

func (bt *BinaryTree) maxValueHelper(node *TreeNode) int {
	if node == nil {
		return -1 << 31 // Min int value
	}

	max := node.Value

	leftMax := bt.maxValueHelper(node.Left)
	if leftMax > max {
		max = leftMax
	}

	rightMax := bt.maxValueHelper(node.Right)
	if rightMax > max {
		max = rightMax
	}

	return max
}

// MinValue returns the minimum value in the tree
// Returns 0 and false if tree is empty
// Time Complexity: O(n)
func (bt *BinaryTree) MinValue() (int, bool) {
	if bt.Root == nil {
		return 0, false
	}

	return bt.minValueHelper(bt.Root), true
}

func (bt *BinaryTree) minValueHelper(node *TreeNode) int {
	if node == nil {
		return 1<<31 - 1 // Max int value
	}

	min := node.Value

	leftMin := bt.minValueHelper(node.Left)
	if leftMin < min {
		min = leftMin
	}

	rightMin := bt.minValueHelper(node.Right)
	if rightMin < min {
		min = rightMin
	}

	return min
}
