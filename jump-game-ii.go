package main
import (
	"fmt"
)

func jump(nums []int) int {
	step := 0
	length := len(nums)
	maxReach := 0
	end := 0
	for i := 0; i < length - 1; i++ {
		max := nums[i] + i
		if max > maxReach {
			maxReach = max
		}
		if i == end {
			step++
			end = maxReach
		}
		
	}
	return step
}

func main() {
	// https://leetcode.cn/problems/jump-game-ii/
	// 方案讲解：https://zhuanlan.zhihu.com/p/82831669
	nums := []int{2,3,1,1,4}
	fmt.Println(jump(nums))
	fmt.Println("xxxx")
}
