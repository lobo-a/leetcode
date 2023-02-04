package main
import (
	"fmt"
)

func maxProfit(prices []int) int {
	ret := 0
	//every day profit
	profits := make([]int, 0)
	maxProfits := make([]int, 2)
	length := len(prices)
	for i := 1; i < length; i ++ {
		profit := prices[i] - prices[i-1]
		profitLen := len(profits)
		isAdd := true
		if profitLen > 0 {
			if (profit * profits[profitLen - 1]) >= 0 {
				profits[profitLen - 1] = profits[profitLen - 1] + profit
				isAdd = false
			}
		}
		if isAdd == true {
			profits = append(profits, profit)
		}
	}
	profitsLen := len(profits)
	for i := 0; i < profitsLen; i ++ {
		if profits[i] > maxProfits[0] || profits[i] > maxProfits[1] {
			if maxProfits[0] > maxProfits[1] {
				maxProfits[1] = profits[i]
			}else {
				maxProfits[0] = profits[i]
			}
		}
	}
	// fmt.Println(profits)
	ret = maxProfits[0] + maxProfits[1]
	return ret
}

func main() {
	// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))
	fmt.Println("xxxx")
}
