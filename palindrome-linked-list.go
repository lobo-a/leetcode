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

func palindromeLinkedList(head *Inode) bool {
	if head == nil {
		return false
	}
	fastCur := head
	slowCur := head
	nodeList := make([]*Inode, 0)
	var secondCur *Inode
	
	for {
		// 偶数个数据
		if fastCur == nil {
			secondCur = slowCur
			break
		}
		// 奇数个数据
		if fastCur.Next == nil {
			secondCur = slowCur.Next
			break
		}
		nodeList = append(nodeList, slowCur)
		fastCur = fastCur.Next.Next
		slowCur = slowCur.Next
	}
	length := len(nodeList)
	idx := length
	for {
		idx = idx - 1
		if idx < 0 || secondCur == nil {
			break
		}
		v1 := nodeList[idx].Val
		v2 := secondCur.Val
		if v1 != v2 {
			return false
		}
		secondCur = secondCur.Next
	}
	return true
}

func main() {
	// https://leetcode.cn/problems/palindrome-linked-list/
	l := bList([]int{1,2,3,3,2,0})
	fmt.Println(palindromeLinkedList(l))
	fmt.Println("xxxx")
}




