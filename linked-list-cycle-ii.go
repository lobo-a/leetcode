package main
import (
	"fmt"
)

type inode struct {
	Value int
	Next *inode
}

func detectCycle(head *inode) *inode{
	if head == nil {
		return nil
	}
	fastNode := head
	slowNode := head
	for {
		slowNode = slowNode.Next
		for i := 0; i < 2; i++ {
			fastNode = fastNode.Next
			if fastNode == nil {
				return nil
			}
		}
		if slowNode == fastNode {
			break
		}
	}
	cur := head
	for {
		if cur == nil || slowNode == nil {
			return nil
		}
		if cur == slowNode {
			break
		}
		slowNode = slowNode.Next
		cur = cur.Next
	}
	return cur
}

func main() {
	// https://leetcode.cn/problems/linked-list-cycle-ii/
	// 方案讲解：https://zhuanlan.zhihu.com/p/105473806
	// 0 11 1 2 9 10 -> 11
	tt := &inode{Value: 11}
	tt.Next = &inode{Value: 1, Next: &inode{Value: 2, Next: &inode{Value: 9, Next: &inode{Value: 10, Next: tt}}}}
	head := &inode{Value: 0, Next: tt}
	fmt.Println(detectCycle(head))
	fmt.Println("xxxx")
}
