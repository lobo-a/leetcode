package main
  
import "fmt"

func subArrLen(nums []int) int {
    ret := 0
    length := len(nums)
    tmpLen := make([]int, length)

    for i := 0; i < length; i++ {
        tmpLenMax := 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                tt := tmpLen[j] + 1
                if tt > tmpLenMax {
                    tmpLenMax = tt
                }
            }
        }
        tmpLen[i] = tmpLenMax
        if tmpLen[i] > ret {
            ret = tmpLen[i]
        }
    }
    return ret
}

func main() {
    // https://leetcode.cn/problems/longest-increasing-subsequence/
    fmt.Println(subArrLen([]int{10,9,2,5,3,7,101,18}))
    fmt.Println(subArrLen([]int{7,7,7,7,7,7,7}))
    fmt.Println(subArrLen([]int{0,1,0,3,2,3}))
    fmt.Println("xxx")
  
}
  
