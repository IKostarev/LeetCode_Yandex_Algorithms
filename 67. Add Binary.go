package leetcode

func main() {
	bin("1111", "1111")
}

func num(min string, max int) string {
	var newMin string

	dop := max - len(min)

	for i := 0; i < dop; i++ {
		newMin += "0"
	}

	newMin += min

	return newMin
}

func bin(a, b string) string {
	var max int
	var res string
	var next = 0

	if len(a) < len(b) {
		max = len(b) + 1
	} else {
		max = len(a) + 1
	}

	a = num(a, max)
	b = num(b, max)

	for i := max - 1; i >= 0; i-- {
		if a[i] == 48 && b[i] == 48 {
			if next == 1 {
				res = "1" + res
				next = 0
			} else {
				res = "0" + res
				next = 0
			}
		}

		if a[i] == 49 && b[i] == 49 {
			if next == 1 {
				res = "1" + res
				next = 1
			} else {
				res = "0" + res
				next = 1
			}
		}

		if a[i] == 48 && b[i] == 49 || a[i] == 49 && b[i] == 48 {
			if next == 1 {
				res = "0" + res
				next = 1
			} else {
				res = "1" + res
				next = 0
			}
		}

	}

	if res[0] == 48 {
		return res[1:]
	}

	return res
}
