package main

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")

	n := len(s)
	if n == 0 {
		return 0
	}

	var start int

	signMultiplier := 1
	if s[0] == '-' {
		signMultiplier = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	var res int
	for i := start; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			return res * signMultiplier
		}

		res = res*10 + int(s[i]-'0')

		if res*signMultiplier <= math.MinInt32 {
			return math.MinInt32
		} else if res*signMultiplier >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res * signMultiplier
}

func main() {
	myAtoi("   -42")
}
