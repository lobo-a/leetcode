package main
  
import "fmt"

// 第n步可以向前或向后跳越n米，给定dist距离，需要多少步能到达位置
func main() {
    fmt.Println(minStep(2))
    fmt.Println("xxx")
}
  

func minStep(dist int) int{
    if dist == 0 {
        return 0
    }
    posArr := []int{0}
    stIdx := 0
    for step := 1; ; step++{
        posLength := len(posArr)
        for i := stIdx; i < posLength; i++ {
            t1 := posArr[i] + step
            if t1 == dist {
                return step
            }
            t2 := posArr[i] - step
            if t2 == dist {
                return step
            }
            posArr = append(posArr, t1, t2)
        }
        stIdx = posLength 
    }
    return -1
}
