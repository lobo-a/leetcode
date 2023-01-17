package main
import (
	"fmt"
)

func mergeArray(a, b []int)[]int{
	ret := make([]int, 0)
	ai := 0
	bi := 0
	aLen := len(a)
	bLen := len(b)
	for{
		if ai < aLen && bi < bLen{
			if a[ai] < b[bi] {
				ret = append(ret, a[ai])
				ai += 1
			}else{
				ret = append(ret, b[bi])
				bi += 1
			}
		}else if ai < aLen{
			ret = append(ret, a[ai])
			ai += 1
		}else if bi < bLen{
			ret = append(ret, b[bi])
			bi += 1
		}else{
			break
		}
	}
	return ret
}

func main() {
	// https://leetcode.cn/problems/median-of-two-sorted-arrays/
	//
	arr := mergeArray([]int{1,2}, []int{3,4})
	length := len(arr)
	if length % 2 == 0 {
		median := float64(arr[length / 2] + arr[length / 2 - 1]) / 2.0
		fmt.Println(median)
	}else{
		median := arr[int(length / 2)]
		fmt.Println(median)
	}
	fmt.Println()
	fmt.Println("xxxx")
}
