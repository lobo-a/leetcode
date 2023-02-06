package main
import (
	"fmt"
)

func singleNumberII(nums []int) int {
	length := len(nums)
	n1,n2 := 0, 0
	for i := 0; i < length; i ++ {
		fmt.Println(n1, n2, nums[i])
		n1 = ( n1 ^ nums[i] ) & (^n2)
		n2 = ( n2 ^ nums[i] ) & (^n1)
	}
	return n1
}

func singleNumberII2(nums []int) int {
	length := len(nums)
	k := 3
	bitMap := make([]int, 32)
	for i := 0; i < length; i ++ {
		ni := nums[i]
		for j := 0; j < 32; j ++ {
			if (1 << j & ni) > 0 {
				bitMap[j]++
			}
		}
	}
	ret := 0
	for j := 0; j < 32; j ++ {
		if bitMap[j] % k != 0 {
			ret = (1 << j) | ret
		}
	}
	return ret
}

func main() {
	// https://leetcode.cn/problems/single-number-ii/
	// https://blog.csdn.net/qq_17550379/article/details/83926804
	nums := []int{3,3,2,3}
	fmt.Println(singleNumberII2(nums))
	fmt.Println("xxxx")
}
