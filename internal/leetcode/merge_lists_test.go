package leetcode

import (
	"reflect"
	"testing"
)

// Helper to convert slice to linked list
func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// Helper to convert linked list to slice
func listToSlice(head *ListNode) []int {
	nums := []int{}
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "both populated",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "one empty",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},
		{
			name:  "both empty",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "negative values",
			list1: []int{-5, -2, 0},
			list2: []int{-10, -5, 5},
			want:  []int{-10, -5, -5, -2, 0, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.list1)
			l2 := sliceToList(tt.list2)
			merged := MergeTwoLists(l1, l2)
			got := listToSlice(merged)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
