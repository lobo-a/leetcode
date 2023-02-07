package main
import (
	"fmt"
)

func gys(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a != b {
		if a > b {
			a = a - b
		}else {
			b = b - a
		}
	}
	return a
}

func maxPoints(nums [][2]int) int {
	ret := 0
	length := len(nums)
	if length == 0 {
		return ret
	}
	for i := 0; i < length - 1; i ++ {
		rateMap := make(map[string]int)
		for j := i + 1; j < length; j ++ {
			xdiff := nums[j][0] - nums[i][0]
			ydiff := nums[j][1] - nums[i][1]
			if xdiff < 0 {
				xdiff = -xdiff
			}
			if ydiff < 0 {
				ydiff = -ydiff
			}
			if xdiff == 0 {
				rateMap["x0"] = rateMap["x0"] + 1
			}else if ydiff == 0 {
				rateMap["y0"] = rateMap["y0"] + 1
			}else{
				xyGys := gys(ydiff, xdiff)
				key := fmt.Sprintf("%d/%d", ydiff/xyGys, xdiff/xyGys)
				rateMap[key] = rateMap[key] + 1
			}
		}
		for _, v := range rateMap{
			if v > ret {
				ret = v
			}
		}
	}
	return ret + 1
}

func main() {
	// https://leetcode.cn/problems/max-points-on-a-line/
	points := [][2]int{[2]int{1,1},[2]int{3,2},[2]int{5,3},[2]int{4,1},[2]int{2,3},[2]int{1,4}}
	fmt.Println(maxPoints(points))
	fmt.Println("xxxx")
}
