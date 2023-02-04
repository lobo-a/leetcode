package main
import (
	"fmt"
)

func maxProfit(prices []int) int {
	ret := 0
	length := len(prices)
	leftProfit := make([]int, length)
	rightProfit := make([]int, length)
	for i, buy := 1, prices[0]; i < length; i ++ {
		profit := prices[i] - buy
		if profit > leftProfit[i - 1] {
			leftProfit[i] = profit
		}else {
			leftProfit[i] = leftProfit[i - 1]
		}
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	for i, sale := length - 2, prices[length - 1]; i > 0; i -- {
		profit := sale - prices[i]
		if profit > rightProfit[i + 1] {
			rightProfit[i] = profit
		}else {
			rightProfit[i] = rightProfit[i + 1]
		}
		if prices[i] > sale {
			sale = prices[i]
		}
		maxProfit := leftProfit[i] + rightProfit[i]
		if maxProfit > ret{
			ret = maxProfit
		}
	}
	return ret
}

func main() {
	// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
	// https://blog.csdn.net/MIC10086/article/details/113872698
	prices := []int{3,3,5,0,0,3,2,5}
	fmt.Println(maxProfit(prices))
	fmt.Println("xxxx")
}
