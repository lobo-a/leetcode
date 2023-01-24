package main
import (
	"fmt"
)

func longestConsecutive(nums []int) int{
	numsMap := make(map[int]int)
	ret := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		v := nums[i]
		if _, ok := numsMap[v]; ok {
			continue
		}
		// 连续的首位和末尾必须记录最大连续值		
		preMax := 0
		nextMax := 0

		if v1, ok := numsMap[v-1]; ok {
			preMax = v1
		}
		if v1, ok := numsMap[v+1]; ok {
			nextMax = v1
		}
		// 新的最大连续值
		numsMap[v] = preMax + nextMax + 1
		if numsMap[v] > ret {
			ret = numsMap[v]
		}
		// 更新连续首位的值
		if preMax > 0 {
			numsMap[v-preMax] = numsMap[v]
		}
		// 更新连续末位的值
		if nextMax > 0 {
			numsMap[v+nextMax] = numsMap[v]
		}
	}
	return  ret
}

func main() {
	// https://leetcode.cn/problems/longest-consecutive-sequence/
	fmt.Println(longestConsecutive([]int{6,100,4,200,1,5,3,2}))
	fmt.Println("xxxx")
}
