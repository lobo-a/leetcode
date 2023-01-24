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
	// 《九章算术》原文:  可半者半之，不可半者，副置分母、子之数，以少减多，更相减损，求其等也。以等数约之.
	fmt.Println(gys(8, 4))
	fmt.Println("xxxx")
}
