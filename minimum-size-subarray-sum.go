package main
  
import "fmt"
  
func minSubArrayLen2(nums []int, target int) int {
    ret := 0
    length := len(nums)
    if length == 0 {
        return ret
    }
    i := 0
    for  i < length {
        tmp := nums[i]
        if tmp >= target && ret == 0{
            ret = 1
            break
        }
        isExist := false
        j := i + 1
        i ++
        for j < length {
            tmp += nums[j]
            if tmp >= target {
                subLength := j - i + 1
                if ret == 0 {
                    ret = subLength
                }else if subLength < ret{
                    ret = subLength
                }
                isExist = true
                break
            }
            j++
        }
        if isExist == true {
            i = j 
        }
    }

    return ret
}

func minSubArrayLen(nums []int, target int) int {
    ret := 0
    length := len(nums)
    i, j := 0, 0
    tmpSum := 0
    for i <= j && j < length {
        tmpSum = tmpSum + nums[j]
        j++
        if tmpSum >= target {
            tmpLen := j - i
            if ret == 0 || tmpLen < ret{
                ret = tmpLen
            }
            i = j
            tmpSum = 0
        }
    }
    return ret
}

func main() {
    // https://leetcode.cn/problems/minimum-size-subarray-sum/ 
    fmt.Println(minSubArrayLen([]int{2,3,1,2,4,3}, 7))
    fmt.Println(minSubArrayLen([]int{1,4,4}, 4))
    fmt.Println(minSubArrayLen([]int{1,1,1,1,1,1,1}, 11))
    fmt.Println("xxx")
  
}
  
