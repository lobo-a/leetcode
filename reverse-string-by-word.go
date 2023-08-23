package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(reverseStr([]byte("cat a is this"))))
}

// reverseStr 字符串按单词反转
// this is a cat
// cat a is this
func reverseStr(str []byte) []byte {
	length := len(str)
	if length <= 1 {
		return str
	}
	// 先全部反转
	str = reverseSubStr(str, 0, length - 1)
	// 按某个单词局部在反转
	start := 0
	end := 0
	for end < length {
		if str[end] == ' '{
			str = reverseSubStr(str, start, end - 1)
			start = end + 1
		}else if end == length-1 {
			str = reverseSubStr(str, start, end)
			break
		}
		end++
	}
	return str
}
func reverseSubStr(str []byte, start, end int) []byte {
	fmt.Println(string(str), start, end)
	for start < end {
		str[end], str[start] = str[start], str[end]
		start++
		end--
	}
	return str
}
