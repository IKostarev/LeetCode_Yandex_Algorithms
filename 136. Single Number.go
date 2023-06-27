package leetcode

import "fmt"

func mySort(items []int) []int {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	return items
}

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	nums = mySort(nums)

	var sum int

	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				sum++
			}

		}

		if sum == 1 {
			fmt.Println(nums[i])
			return nums[i]
		}
	}

	return 0
}

func main() {
	num := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(num))
}
