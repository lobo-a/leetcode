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

func cutNode(head *inode, n int) *inode {
	var del, delPre, cur *inode
	cur = head
	i := 0
	for cur != nil {
		if i+1 == n {
			del = head
		}else if i+1 > n {
			delPre = del
			del = del.Next
		}
		i += 1
		cur = cur.Next
	}
	if delPre != nil {
		tmp := delPre.Next
		if tmp != nil {
			delPre.Next = tmp.Next
		}
	}else if del != nil {
		head = del.Next
	}
	return head
}

func main() {
	// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
	//
	l := intToList(54321)
	l = cutNode(l, 2)
	for l != nil {
		fmt.Println(l.V)
		l = l.Next
	}
	fmt.Println("xxxx")
}
