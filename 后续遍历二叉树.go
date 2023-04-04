package main
  
import "fmt"
// 非递归的方式实现二叉树的后序遍历


type Inode struct {
    Val int
    Left *Inode
    Right *Inode
  }
  
  func scanTree(root *Inode) {
    var cur *Inode
    tmpNodes := []*Inode{ root }
    isScanSubTree := make(map[int]bool, 0)
    for {
      length := len(tmpNodes)
      if length == 0 {
        break
      }
      cur = tmpNodes[length - 1]
      tmpNodes = tmpNodes[0: length - 1]
      //fmt.Println(cur, tmpNodes)
      tmpLeft := cur.Left
      tmpRight := cur.Right
      if v, ok := isScanSubTree[length - 1]; ok && v {
        fmt.Println(cur.Val)
      }else if tmpLeft == nil && tmpRight == nil {
        fmt.Println(cur.Val)
      }else {
        isScanSubTree[length - 1] = true
        tmpNodes = append(tmpNodes, cur)
        if tmpRight != nil {
            tmpNodes = append(tmpNodes, tmpRight)
        }
        if tmpLeft != nil {
            tmpNodes = append(tmpNodes, tmpLeft)
        }
      }
      
    }
  }

func main() {
    n1 := &Inode{Val: 1}
    n2 := &Inode{Val: 2}
    n3 := &Inode{Val: 3}
    n4 := &Inode{Val: 4}
    n5 := &Inode{Val: 5}
    n6 := &Inode{Val: 6}
    n7 := &Inode{Val: 7}
    n8 := &Inode{Val: 8}
    n9 := &Inode{Val: 9}
    n10 := &Inode{Val: 10}
    n1.Left = n2
    n1.Right = n3
    n2.Left = n4
    n2.Right = n5
    n3.Left = n6
    n3.Right = n7
    n4.Left = n8
    n4.Right = n9
    n5.Left = n10
    scanTree(n1)
    fmt.Println("xxx")
}
