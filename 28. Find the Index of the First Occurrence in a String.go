package leetcode

import (
	"fmt"
)

func main() {
	res := strStr("hello", "ll")

	fmt.Println(res)
}
func strStr(haystack string, needle string) int {
	flag := false
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < len(needle); j++ {
				if i+j < len(haystack) && haystack[i+j] == needle[j] {
					flag = true
				} else {
					flag = false
					break
				}
			}
		}
		if flag {
			return i
		}
	}
	return -1
}
