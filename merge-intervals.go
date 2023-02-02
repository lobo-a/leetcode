package main
import (
	"fmt"
)

func mergeIntervals(intervals [][2]int)[][2]int{
	ret := make([][2]int, 0)
	isMerge := make(map[int]int)
	tmpIntervals := intervals
	length := len(tmpIntervals)
	for i := 0; i < length; i++ {
		isMerge[i] = 0
	}
	
	for i := 0; i < len(tmpIntervals)-1; i ++ {
		if isMerge[i] == 1 {
			continue
		}
		tmpi := tmpIntervals[i]
		isAdd := false
		
		for j := i + 1; j < len(tmpIntervals); j++ {
			if isMerge[j] == 1 {
				continue
			}
			tmpj := tmpIntervals[j]
			// merge
			if tmpi[0] <= tmpj[1] && tmpj[0] <= tmpi[1] {
				isAdd = true
				isMerge[i] = 1
				isMerge[j] = 1
				if tmpi[0] > tmpj[0] {
					tmpi[0] = tmpj[0]
				}
				if tmpi[1] < tmpj[1] {
					tmpi[1] = tmpj[1]
				}
			}
		}
		if isAdd == true {
			tmpIntervals = append(tmpIntervals, tmpi)
			isMerge[length] = 0
			length ++
		}
	}
	// fmt.Println("xx",tmpIntervals)
	for i := 0; i < len(tmpIntervals); i ++ {
		if isMerge[i] == 0 {
			ret = append(ret, tmpIntervals[i])
		}
	}
	return ret
}

func main() {
	// https://leetcode.cn/problems/merge-intervals/
	intervals := [][2]int{
		[2]int{1,3},
		[2]int{2,6},
		[2]int{8,10},
		[2]int{9,15},
		[2]int{15,18},

	}
	fmt.Println(mergeIntervals(intervals))
	fmt.Println("xxxx")
}
