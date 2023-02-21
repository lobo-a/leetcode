package main
  
import "fmt"

func findKthLargest(nums []int, k int) int {
    length := len(nums)
    if length == 0{
        return 0
    }
    return qSort(nums, 0, length-1, k)
}

func qSort(nums []int, st, end, k int) int {
    if st < end {
        pos := qItemPos(nums, st, end)
        if (pos + 1) == k {
            return nums[pos]
        }else if (pos + 1) > k {
            return qSort(nums, st, pos - 1, k)
        }else if (pos + 1) < k {
            return qSort(nums, pos + 1, end, k)
        }
    }else if st == end {
        if (st + 1) == k {
            return nums[st]
        }
    }
    return -1
}

func qItemPos(nums []int, st, end int) int {
    for st < end {
        for st < end && nums[st] >= nums[end] {
            end--
        }
        nums[st], nums[end] = nums[end], nums[st]
        for st < end && nums[st] > nums[end] {
            st++
        }
        nums[st], nums[end] = nums[end], nums[st]
    }
    return st
}

func main() {
    // https://leetcode.cn/problems/kth-largest-element-in-an-array/
    fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
    fmt.Println("xxx")
  
}
  
