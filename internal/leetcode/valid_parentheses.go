package leetcode

// Problem 20: Valid Parentheses
// Difficulty: Easy
// Link: https://leetcode.com/problems/valid-parentheses/
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// Key Takeaway: Use a Stack (LIFO) data structure. Push open brackets, pop and match when encountering a close bracket.
// Space Complexity: O(n) for the stack.

func IsValidParentheses(s string) bool {
	// If length is odd, it can't be valid
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// If it's a closing bracket
		if open, isClose := pairs[char]; isClose {
			if len(stack) == 0 {
				return false // No open bracket to match
			}
			// Pop the top
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if last != open {
				return false // Mismatched
			}
		} else {
			// It's an opening bracket, push to stack
			stack = append(stack, char)
		}
	}

	// Valid only if stack is empty (all opened brackets are closed)
	return len(stack) == 0
}
