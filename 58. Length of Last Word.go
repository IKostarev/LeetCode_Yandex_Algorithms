package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	var count int

	switch {
	case len(s) == 0:
		return 0
	case len(s) == 1:
		return 1
	case len(s) == 2 && s[1] == 32:
		return 1
	case len(s) == 2 && s[0] == 32:
		return 1
	}

	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 32 {
			continue
		}
		if s[i] != 32 {
			count++
		}
		if s[i-1] == 32 {
			return count
		}
	}

	s = strings.TrimSpace(s)

	return len(s)
}

func main() {
	lengthOfLastWord("   fly me   to   the moon  ")
}
