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

func gbs(a, b int) int {
	gys := gys(a, b)
	if gys == 0 {
		gys = 1
	}
	return a * b / gys 
}

func main() {
	fmt.Println(gys(8, 4))
	fmt.Println("xxxx")
}
