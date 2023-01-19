package main
import (
	"fmt"
)

func strStr(haystack, needle string) int {
	needleLength := len(needle)
	haystackLength := len(haystack)
	for  i, j := 0, needleLength; j < haystackLength; i, j = i+1, j+1 {
		if haystack[i:j] == needle {
			return i
		}
	}
	return -1
}

func main() {
	// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
	//
	fmt.Println( strStr("1234567890", "6789") )
	fmt.Println("xxxx")
}
