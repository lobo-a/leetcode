package main
import (
	"fmt"
)

func gys(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		}else if a < b {
			b = b - a
		}
	}
	return a
}

func gys2(a, b int) int {
    if a < b{
        a, b = b, a
    }
     temp := a % b
    for temp != 0 {
        a = b
        b = temp
        temp = a % b
    }
    return b
}

func main() {
	fmt.Println(gys(8, 4))
	fmt.Println("xxxx")
}
