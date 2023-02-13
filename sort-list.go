package main
import (
	"fmt"
)
type Inode struct {
    Val int
    Next *Inode
}

func bList(nums []int) (head *Inode) {
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

func sortList(h *Inode) (head *Inode){
	if h == nil {
		return nil
	}
	tmpHeads := []*Inode{h}
	var retCur *Inode
	for {
		length := len(tmpHeads)
		if length == 0 {
			break
		}
		cur := tmpHeads[length - 1]
		tmpHeads = tmpHeads[:length-1]
		if cur.Next == nil {
			if head == nil {
				head = cur
				retCur = cur
			}else{
				retCur.Next = cur
				retCur = retCur.Next
			}
			continue
		}
		leftH, midNode, rightH := sortOneNode(cur)
		if rightH != nil {
			tmpHeads = append(tmpHeads, rightH)
		}
		if midNode != nil {
			tmpHeads = append(tmpHeads, midNode)
		}
		if leftH != nil {
			tmpHeads = append(tmpHeads, leftH)
		}
	}
	return
}

func sortOneNode(h *Inode) (leftH, midNode, rightH *Inode){
	if h == nil{
		return nil, h, nil
	}
	cur := h.Next
	pre := h
	var LeftCur *Inode
	for cur != nil {
		tmp := cur
		if tmp.Val > h.Val {
			pre = cur
			cur = cur.Next
		}else{
			cur = cur.Next
			pre.Next = cur
			if LeftCur == nil {
				LeftCur = tmp
				LeftCur.Next = nil
				leftH = LeftCur
			} else {
				tmp.Next = LeftCur.Next
				LeftCur.Next = tmp
			}
		}
	}
	rightH = h.Next
	h.Next = nil
	midNode = h
	return
}

func sortNums(nums []int, stIdx, endIdx int) []int {
	length := len(nums)
	if length == 0 || length == 1 || endIdx >= length {
		return nums
	}
	if stIdx < endIdx {
		pos := sortOneNums(nums, stIdx, endIdx)
		sortNums(nums, stIdx, pos - 1)
		sortNums(nums, pos + 1, endIdx)
	}

	return nums
}

func sortOneNums(nums []int, stIdx, endIdx int) int{
	for endIdx > stIdx {
		for endIdx > stIdx && endIdx >= 0 && nums[endIdx] >= nums[stIdx] {
			endIdx = endIdx -1
		}
		nums[stIdx], nums[endIdx] = nums[endIdx], nums[stIdx]
		for endIdx > stIdx && endIdx >= 0 && nums[endIdx] > nums[stIdx] {
			stIdx = stIdx +1
		}
		nums[stIdx], nums[endIdx] = nums[endIdx], nums[stIdx]	
	}
	return stIdx
}

func main() {
	// https://leetcode.cn/problems/sort-list/ 参考快速排序
	nums := []int{5,1,3,9,2,6,0,4}
	l := bList(nums)
	l = sortList(l)
	for l != nil {
		fmt.Println("v = ",l.Val)
		l = l.Next
	}
	fmt.Println("xxxx")
}

