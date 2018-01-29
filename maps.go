package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	val, prs := m["k1"]
	fmt.Println(val, prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}
