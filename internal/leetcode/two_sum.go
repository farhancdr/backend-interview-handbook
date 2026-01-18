package leetcode

import (
	"errors"
)

// Problem 1: Two Sum
// Difficulty: Easy
// Link: https://leetcode.com/problems/two-sum/
//
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// Key Takeaway: Use a Hash Map for O(1) lookups to achieve O(n) time complexity.
// Brute force is O(n^2).

func TwoSum(nums []int, target int) ([]int, error) {
	// Map to store value -> index
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, found := seen[complement]; found {
			return []int{idx, i}, nil
		}
		seen[num] = i
	}

	return nil, errors.New("no solution found")
}
