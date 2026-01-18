# 💻 LeetCode Practice

> **Real interview problems with complete solutions**

This package contains popular LeetCode problems implemented in Go with comprehensive tests. Perfect for interview preparation and pattern practice.

---

## 📖 Problems

| Problem | File | Difficulty | Pattern | Key Concepts |
|:--------|:-----|:-----------|:--------|:-------------|
| **Two Sum** | [two_sum.go](two_sum.go) | Easy | Hash Map | O(n) lookup, space-time tradeoff |
| **Valid Parentheses** | [valid_parentheses.go](valid_parentheses.go) | Easy | Stack | LIFO, matching pairs |
| **Best Time to Buy/Sell Stock** | [stock.go](stock.go) | Easy | Greedy | Single pass, track minimum |
| **Merge Two Sorted Lists** | [merge_lists.go](merge_lists.go) | Easy | Two Pointers | Linked list, merge algorithm |

---

## 🚀 Quick Start

```bash
# Run all LeetCode tests
go test -v ./leetcode/

# Run specific problem
go test -v ./leetcode/ -run TestTwoSum

# Check test coverage
go test -cover ./leetcode/
```

---

## 🎓 Problem-Solving Framework

### Step-by-Step Approach

1. **Understand the Problem**
   - Read carefully
   - Identify inputs and outputs
   - Ask clarifying questions
   - Consider edge cases

2. **Plan Your Approach**
   - Brute force first
   - Identify patterns
   - Optimize time/space complexity
   - Choose data structures

3. **Implement**
   - Write clean code
   - Handle edge cases
   - Use meaningful variable names
   - Add comments for complex logic

4. **Test**
   - Test normal cases
   - Test edge cases
   - Test large inputs
   - Verify time/space complexity

5. **Optimize**
   - Can you do better?
   - Trade-offs?
   - Alternative approaches?

---

## 💡 Problem Breakdown

### Two Sum
```go
// Problem: Find two numbers that add up to target
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1] (nums[0] + nums[1] = 9)

// Approach 1: Brute Force - O(n²)
// Approach 2: Hash Map - O(n) ✅

// Key insight: For each number x, check if (target - x) exists
```

**Pattern:** Hash Map for O(1) lookup  
**Time:** O(n)  
**Space:** O(n)

---

### Valid Parentheses
```go
// Problem: Check if parentheses are balanced
// Input: "({[]})"
// Output: true

// Approach: Use stack
// 1. Push opening brackets
// 2. Pop and match closing brackets
// 3. Stack should be empty at end
```

**Pattern:** Stack for matching pairs  
**Time:** O(n)  
**Space:** O(n)

---

### Best Time to Buy/Sell Stock
```go
// Problem: Maximize profit with single buy/sell
// Input: prices = [7,1,5,3,6,4]
// Output: 5 (buy at 1, sell at 6)

// Approach: Track minimum price seen so far
// For each price, calculate profit if sold today
```

**Pattern:** Greedy, single pass  
**Time:** O(n)  
**Space:** O(1)

---

### Merge Two Sorted Lists
```go
// Problem: Merge two sorted linked lists
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

// Approach: Two pointers
// Compare heads, append smaller, move pointer
```

**Pattern:** Two pointers, merge algorithm  
**Time:** O(n + m)  
**Space:** O(1)

---

## 🧪 Testing Strategy

Each problem includes:
- **Normal cases** - Typical inputs
- **Edge cases** - Empty, single element, duplicates
- **Large inputs** - Performance validation
- **Invalid inputs** - Error handling

```go
func TestTwoSum(t *testing.T) {
    tests := []struct {
        name   string
        nums   []int
        target int
        want   []int
    }{
        {"normal case", []int{2, 7, 11, 15}, 9, []int{0, 1}},
        {"duplicates", []int{3, 3}, 6, []int{0, 1}},
        {"negative numbers", []int{-1, -2, -3, -4}, -6, []int{2, 3}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := twoSum(tt.nums, tt.target)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## 📊 Complexity Analysis

| Problem | Time | Space | Pattern |
|:--------|:-----|:------|:--------|
| Two Sum | O(n) | O(n) | Hash Map |
| Valid Parentheses | O(n) | O(n) | Stack |
| Stock | O(n) | O(1) | Greedy |
| Merge Lists | O(n+m) | O(1) | Two Pointers |

---

## ⚠️ Common Mistakes

1. **Two Sum**
   - Using same element twice
   - Not handling duplicates
   - Returning values instead of indices

2. **Valid Parentheses**
   - Not checking stack empty before pop
   - Not checking stack empty at end
   - Wrong bracket matching

3. **Stock**
   - Buying and selling on same day
   - Not handling decreasing prices
   - Negative profit

4. **Merge Lists**
   - Forgetting to append remaining nodes
   - Not handling nil lists
   - Modifying original lists

---

## 🔗 Related Topics

- **[Data Structures](../ds/)** - Stack, linked list implementations
- **[Algorithms](../algo/)** - Algorithm patterns used in solutions
- **[Basics](../basics/)** - Go fundamentals

---

## 📚 Interview Tips

1. **Communicate** - Talk through your thought process
2. **Start Simple** - Brute force first, then optimize
3. **Test** - Walk through examples before coding
4. **Edge Cases** - Always consider empty, single, duplicate inputs
5. **Complexity** - State time and space complexity
6. **Clean Code** - Readable variable names, proper structure

---

[← Back to Internal](../) | [↑ Back to Main](../../README.md)
