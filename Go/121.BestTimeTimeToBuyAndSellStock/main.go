package main

import "fmt"

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for i := range prices {
		if prices[i] <= buy {
			buy = prices[i]
		}
		if (prices[i] - buy) > profit {
			profit = (prices[i] - buy)
		}
	}
	return profit
}

func main() {
	prices := []int{2, 8, 1, 5}
	fmt.Println(maxProfit(prices))
}
