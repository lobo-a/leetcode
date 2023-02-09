package main
import (
	"fmt"
)

type Inode struct {
	Val int
	Next *Inode
}

func bList(nums []int) (head *Inode){
	length := len(nums)
	var cur *Inode
	for i := 0; i < length; i++ {
		if head == nil {
			head = &Inode{Val: nums[i]}
			cur = head
		}else{
			cur.Next = &Inode{Val: nums[i]}
			cur = cur.Next
		}
	}
	return
}

func deleteDuplicates(h *Inode) (head *Inode) {
	valMap := make(map[int]int)
	cur := h
	for cur != nil {
		valMap[cur.Val] = valMap[cur.Val] + 1
		cur = cur.Next
	}
	cur = h
	pre := h
	for cur != nil {
		tmp := cur
		val := cur.Val
		cur = cur.Next
		if v, ok := valMap[val]; ok && v > 1 {
			if tmp == h {
				pre = cur
				h = cur
			}else{
				pre.Next = cur
				tmp.Next = nil
			}	
		}else{
			pre = tmp
		}
	}
	return h
}

func main() {
	// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
	l := bList([]int{1,1,1,2,3})
	l = deleteDuplicates(l)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println("xxxx")
}
