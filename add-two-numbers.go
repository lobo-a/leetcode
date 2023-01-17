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

func add(a, b *inode)(head *inode){
	var c, cur *inode
	for{
		ai := 0
		bi := 0
		ci := 0

		if a == nil && b == nil && c == nil{
			break
		}

		if a != nil{
			ai = a.V
			a = a.Next
		}

		if b != nil {
			bi = b.V
			b = b.Next
		}

		if c != nil {
			ci = c.V
			c = nil
		}
		sum := ai + bi + ci
		if sum > 9 {
			c = &inode{V: int(sum / 10) }
		}
		if head == nil {
			head = &inode{V: sum % 10 }
			cur = head
		}else {
			cur.Next = &inode{V: sum % 10}
			cur = cur.Next
		}

	}
	return
}

func main() {
	// https://leetcode.cn/problems/add-two-numbers/
	// 
	a := intToList(123)
	b := intToList(888)
	l := add(a, b)
	for l != nil {
		fmt.Println(l.V)
		l = l.Next
	}
	fmt.Println("xxxx")
}
