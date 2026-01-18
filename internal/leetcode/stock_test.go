package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"standard case", []int{7, 1, 5, 3, 6, 4}, 5}, // Buy at 1, sell at 6
		{"descending prices", []int{7, 6, 4, 3, 1}, 0},
		{"ascending prices", []int{1, 2, 3, 4, 5}, 4},
		{"empty", []int{}, 0},
		{"single element", []int{5}, 0},
		{"random peaks", []int{2, 4, 1, 7}, 6}, // Buy at 1, sell at 7
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
