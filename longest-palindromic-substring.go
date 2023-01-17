package main
import (
	"fmt"
)

func longestPalindromicSubstring(str string) string{
	ret := str[0:1]
	length := len(str)
	arr := make([][]int, length)
	for i := 0; i < length; i++ {
		arr[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		arr[i][i] = 1
		if (i+1) < length && str[i] == str[i+1] {
			arr[i][i+1] = 1
			ret = str[i:i+2]
		}
	} 

	for i := 0; i < length; i++ {
		for j := i+2; j < length; j++ {
			if str[i] == str[j] && arr[i+1][j-1] == 1 {
				arr[i][j] = 1
				tmps := str[i:j+1]
				if len(tmps) > len(ret){
					ret = tmps
				}
			} 
		}	
	}
	return ret
}


func main() {
	// https://leetcode.cn/problems/longest-palindromic-substring/
	//
	str := "cbbd"
	fmt.Println(longestPalindromicSubstring(str))
	fmt.Println("xxxx")
}
