package main
  
import "fmt"
//回旋数组
func printArr(n int){
    if n <= 0 {
        return
    }
    // 初始化数组
    max := n * n
    arr := make([][]int, n)
    for i := 0; i < n ; i ++ {
        arr[i] = make([]int, n)
    }
    x, y := 0, 0
    turn := 0
    for i:=1; i <= max; i++ {
        arr[x][y] = i
        mod := turn % 4
        if mod == 0 { // x->, y++
            if y == (n-1) || arr[x][y+1] != 0{
                x++
                turn++
            }else{
                y++
            }
        } else if mod == 1 { // x++, y->
            if x == (n-1) || arr[x+1][y] != 0 {
                y--
                turn++
            }else{
                x++
            }
        } else if mod == 2 { // x->, y--
            if y == 0 || arr[x][y-1] != 0{
                x--
                turn++
            }else{
                y--
            }
        } else if mod == 3 { // x--, y->
            if x == (n-1) || arr[x-1][y] != 0{
                y++
                turn++
            }else{
                x--
            }
        }
        
    }
    fmt.Println(arr)
}

func main() {
    printArr(5)
    fmt.Println("xxx")
}
