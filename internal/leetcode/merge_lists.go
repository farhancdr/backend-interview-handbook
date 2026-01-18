package leetcode

// Problem 21: Merge Two Sorted Lists
// Difficulty: Easy
// Link: https://leetcode.com/problems/merge-two-sorted-lists/
//
// Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//
// Key Takeaway: Use a dummy head node to simplify edge cases. Iterate while both lists are non-nil.
// Time: O(n+m)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Append whatever is left
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return dummyHead.Next
}
