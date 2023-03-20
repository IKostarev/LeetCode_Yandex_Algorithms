package leetcode

func mySqrt(x int) int {
	var del = 0

	for del*del < x {
		del++
	}

	if del*del > x {
		del--
	}

	return del
}
