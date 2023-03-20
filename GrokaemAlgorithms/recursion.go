package GrokaemAlgorithms

import "fmt"

func look_for_key(box []int) {
	for item := 0; item <= len(box); item++ {
		if item.is_a_box {
			look_for_key(item)
		} else if item.is_a_key {
			fmt.Println("You found key!")
		}
	}
}
