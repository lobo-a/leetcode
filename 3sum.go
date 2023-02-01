package main
import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	numsMap := make(map[int]int)
	for i := 0; i < length; i ++ {
		numsMap[nums[i]] = 0
	}
	for i := 0; i < length; i ++ {
		for j := i + 1; j < length; j++ {
			tmp := nums[i] + nums[j]
			tmp = 0 - tmp
			if v, ok := numsMap[tmp]; ok && v == 0 {
				ret = append(ret, []int{nums[i] , nums[j] , tmp})
				numsMap[tmp] = 1
				numsMap[nums[i]] = 1
				numsMap[nums[j]] = 1
			}
		}
	}
	return ret
}

func main() {
	// https://leetcode.cn/problems/3sum/
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
	fmt.Println("xxxx")
}
