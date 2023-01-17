package main
import (
	"fmt"
)

func findLongestSubStr(str string) int {
	num := 0
	characterMap := make(map[int32]struct{})
	for _, s := range str{
		if _, ok := characterMap[s]; ok {
			length := len(characterMap)
			if length > num{
				num = length
			}
			characterMap = make(map[int32]struct{})
		}else{
			characterMap[s] = struct{}{}
		}
	}
	length := len(characterMap)
	if length > num{
		num = length
	}
	return num
}

func main() {
	// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
	//
	str := "pwwkew"
	fmt.Println(findLongestSubStr(str))
	fmt.Println("xxxx")
}
