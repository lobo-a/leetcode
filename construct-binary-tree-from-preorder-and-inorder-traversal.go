package main
import (
	"fmt"
)


type inode struct {
	Value int
	Left *inode
	Right *inode
}

func buildTree(preorder, inorder []int) (head *inode){
	length := len(preorder)
	tmpInorder := inorder
	items := make([][]int, 0)
	posIdx := 0
	var cur *inode
	for i := 0; i < length; i++ {
		if posIdx == 0 {
			head = &inode{Value: preorder[i]}
			cur = head
		} else if posIdx == 1 {
			cur.Left = &inode{Value: preorder[i]}
			cur = cur.Left
		} else if posIdx == 2 {
			cur.Right = &inode{Value: preorder[i]}
			cur = cur.Right
		}
		leftItems, rightItems := findNodeItems(preorder[i], tmpInorder)
		if len(rightItems) > 0 {
			items = append(items, rightItems)
		}
		if len(leftItems) > 0 {
			posIdx = 1
			tmpInorder = leftItems
		}else if len(items) > 0 {
			posIdx = 2
			itemLength := len(items)
			tmpInorder = items[itemLength - 1]
			items = items[:itemLength - 1]
		}else{
			break
		}
	}
	
	return
}

func findNodeItems(root int, inorder []int)(leftItems, rightItems []int){
	if len(inorder) == 0 {
		return
	}
	mid := -1
	length := len(inorder)
	for i := 0; i < length; i++ {
		if root == inorder[i] {
			mid = i
			break
		}
	}
	leftItems = inorder[0:mid]
	rightItems = inorder[mid+1:]
	return
}

func main() {
	// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
	// 3,4,8,9,1,
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	head := buildTree(preorder, inorder)
	items := []*inode{head}
	var cur *inode
	for {
		if cur == nil && len(items) == 0 {
			break
		}
		if cur == nil {
			cur = items[0]
			items = items[1:]
		}
		if cur == nil {
			fmt.Println("null")
			continue
		}
		fmt.Println(cur.Value)
		if cur.Right != nil {
			items = append(items, cur.Right)
		}
		if cur.Left != nil {
			cur = cur.Left
		}else{
			cur = nil
		}

		
	}
	fmt.Println("xxxx")
}
