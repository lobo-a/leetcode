package main
import (
	"fmt"
)

func recv(num int) int {
	ret := 0
	for {
		if num == 0{
			break
		}
		tmpi := num % 10
		num = int(num / 10)
		ret = ret * 10 + tmpi
	}
	return ret
}

func recvEqual(a, b int) bool{
	return (recv(a) + b) == (a + recv(b))
}

func recvEqualList(num []int)[][2]int{
	ret := make([][2]int, 0)
	l := len(num)

	for x:=0; x < l; x++ {
		for y:=x+1; y< l; y++{
			if recvEqual(num[x], num[y]){
				ret = append(ret, [2]int{x, y})
			}
		}	
	}

	return ret
}

func main() {
  // https://leetcode.cn/problems/count-nice-pairs-in-an-array/
	// nums := []int{42,11,1,97}
	nums := []int{13,10,35,24,76}
	recvEqualList := recvEqualList(nums)
	fmt.Println(recvEqualList)
}
