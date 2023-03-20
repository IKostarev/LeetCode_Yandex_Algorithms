package GrokaemAlgorithms

func binary_search(list []int, item int) (result int, searchCount int) {
	mid := len(list) / 2

	switch {
	case len(list) == 0:
		return
	case list[mid] > item:
		result, searchCount = binary_search(list[:mid], item)
	case list[mid] < item:
		result, searchCount = binary_search(list[mid+1:], item)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}
	searchCount++
	return
}
