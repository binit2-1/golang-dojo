package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	n := map[string]int{"k1": 1} //map without make
	fmt.Println("map2:", n)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"] //if k2 not there in map it returns false
	fmt.Println("prs:", prs)

	if maps.Equal(m, n) {
		fmt.Println("m == n")
	} else {
		fmt.Println("not equal")
	}
}
