package GrokaemAlgorithms

import "fmt"

func hash_map() {
	var m map[string]int

	n := make(map[string]string)

	m["route"] = 66
	n["name"] = "Ilya"

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
