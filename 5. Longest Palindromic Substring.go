package leetcode

func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}

	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i+1)
		len := 0

		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}

		if len > end-start {
			start = i - ((len - 1) / 2)
			end = i + (len / 2)
		}

	}

	return s[start : end+1]

}

func expandFromCenter(s string, left, right int) int {

	if s == "" || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
