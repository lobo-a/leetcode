package main
import (
	"fmt"
)

type inode struct{
	V int
	Next *inode
}

func intToList(num int) (head *inode) {
	if num == 0 {
		head = &inode{V: 0}
		return 
	}
	var cur *inode
	for{
		if num == 0 {
			break
		}
		tmpi := num % 10
		num = int(num / 10)
		if head == nil {
			head = &inode{V: tmpi}
			cur = head
		}else {
			cur.Next = &inode{V: tmpi}
			cur = cur.Next
		}
	}
	return
}

func reverseKGroup(head *inode, k int)(h *inode){
	i := 0
	cur := head
	preHead := head
	var preEnd *inode
	for {
		if cur == nil{
			break
		}
		i += 1
		tmp := cur
		cur = cur.Next
		if i % k == 0 {
			tmp.Next = nil
			tmpH, tmpE := reverseList(preHead)
			if h == nil {
				h = tmpH
			}
			if preEnd == nil {
				preEnd = tmpE
			}else{
				preEnd.Next = tmpH
				preEnd = tmpE
			}
			preHead = cur
		}
	}
	if i % k != 0 {
		preEnd.Next = preHead
	}
	return
}

func reverseList(h *inode)(head, end *inode){
	tmp := h
	var cur *inode 
	for{
		if tmp == nil {
			break
		}
		if cur == nil {
			cur = tmp
			tmp = tmp.Next
			cur.Next = nil
			end = cur
		}else{
			tt := tmp.Next
			tmp.Next = cur
			cur = tmp
			tmp = tt
		}
	}
	head = cur
	return
} 

func main() {
	// https://leetcode.cn/problems/reverse-nodes-in-k-group/
	//
	l := intToList(987654321)

	l = reverseKGroup(l, 4)
	for l != nil {
		fmt.Println(l.V)
		l = l.Next
	}
	fmt.Println("xxxx")
}
