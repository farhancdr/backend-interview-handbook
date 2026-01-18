package leetcode

import "math"

// Problem 121: Best Time to Buy and Sell Stock
// Difficulty: Easy
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
// Key Takeaway: "Sliding Window" / Single Pass. Track the `minPrice` seen so far and calculate `currentPrice - minPrice`.
// Time: O(n), Space: O(1).

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
