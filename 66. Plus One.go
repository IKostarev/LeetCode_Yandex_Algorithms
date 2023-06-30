package leetcode

func main() {
	//plusOne([]int{1, 2, 3})
	plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6})
}
func plusOne(digits []int) []int {
	return increment(digits)
}

func increment(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	last_digit := digits[len(digits)-1]

	last_digit += 1
	if last_digit > 9 {
		last_digit = 0
		result := increment(digits[:len(digits)-1])
		return append(result, last_digit)
	} else {
		digits[len(digits)-1] = last_digit
		return digits
	}
}
