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

type snode struct{
	V string
	Next *snode
}
func strToList(str string) (head *snode){
	length := len(str)
	var cur *snode
	for i := 0; i < length; i++{
		// fmt.Println( int(str[i]) - int('0'))
		if head == nil {
			head = &snode{V: string(str[i])}
			cur = head
		}else{
			cur.Next = &snode{V: string(str[i])}
			cur = cur.Next
		}
	}
	return
}


func swapPairs(h *snode)(head *snode){
	i := 0
	pre := h
	cur := h
	for{
		i += 1
		if cur == nil {
			break
		}
		second := cur
		cur = cur.Next
		if i % 2 == 0{
			if head == nil {
				second.Next = pre
				pre.Next = cur
				head = second
			}else{
				first := pre.Next
				first.Next = cur
				second.Next = first
				pre.Next = second 
				pre = first
			}
		}
	}
	return
}

func main() {
	// https://leetcode.cn/problems/swap-nodes-in-pairs/
	//
	l := strToList("123456789")
	l = swapPairs(l)
	for l != nil {
		fmt.Println(l.V)
		l = l.Next
	}
	
	fmt.Println("xxxx")
}
