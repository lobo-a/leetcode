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

func mergeList(heads... *inode)(head *inode){
	length := len(heads)
	curs := make([]*inode, length)
	for i := 0; i < length; i++ {
		if heads[i] != nil {
			curs[i] = heads[i] 
		}
	}
	var tmp *inode
	tmpList := heads
	for {
		ii, minNode := findMinNode(curs)
		if minNode == nil {
			break
		}
		if head == nil {
			head = minNode
			tmp = head
		}else{
			tmp.Next = minNode
			tmp = tmp.Next
		}
		if tmpList[ii] != nil {
			curs[ii] = tmpList[ii].Next
			tmpList[ii] = tmpList[ii].Next
		}
	}	
	return
}

func findMinNode(nodes []*inode)(ii int, node *inode){
	length := len(nodes)
	for i := 0; i < length; i++ {
		if nodes[i] == nil {
			continue
		}
		if node == nil || node.V > nodes[i].V {
			ii = i
			node = nodes[i]
		}
	}
	return
}

func main() {
	// https://leetcode.cn/problems/merge-k-sorted-lists/
	//
	l1 := intToList(541)
	l2 := intToList(431)
	l3 := intToList(62)

	l := mergeList(l1,l2,l3)
	for l != nil {
		fmt.Println(l.V)
		l = l.Next
	}
	fmt.Println("xxxx")
}
