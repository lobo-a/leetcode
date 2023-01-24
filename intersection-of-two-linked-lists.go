package main
import (
	"fmt"
)

type inode struct {
	Value int
	Next *inode
}

func getIntersectionNode(headA, headB *inode) *inode {
	var fNode *inode
	curA := headA
	curB := headB
	for {
		if curA == curB {
			fNode = curA
			break
		}
		if curA == nil {
			curA = headB
		}
		if curB == nil {
			curB = headA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return fNode
}

func main() {
  // https://leetcode.cn/problems/intersection-of-two-linked-lists/
	tt := &inode{Value: 9, Next: &inode{Value: 10}}
	headA := &inode{Next: &inode{Next: tt}}
	headB := &inode{Next: tt}
	fmt.Println(getIntersectionNode(headA, headB))
	fmt.Println("xxxx")
}
