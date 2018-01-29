package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	fmt.Println(s[2:5])
	fmt.Println(s[:5])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := 1 + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
