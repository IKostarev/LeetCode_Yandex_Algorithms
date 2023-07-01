package leetcode

func main() {
	climbStairs(5)
}

func climbStairs(n int) int {
	if n == 1 {
		return n
	}

	n1 := 1
	n2 := 2

	for i := 3; i <= n; i++ {
		next := n1
		n1 = n2
		n2 += next
	}

	return n2
}
