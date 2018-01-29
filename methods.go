package main

import "fmt"

type rect struct {
	width, height int
}

// Pointer type method
func (r *rect) area() int {
	r.width = r.width * 2
	return r.width * r.height
}

// Value receiver type method
func (r rect) perim() int {
	r.width = r.width * 2
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Println(r.area())
	fmt.Println(r)
	fmt.Println(r.perim())
	fmt.Println(r)
	
	r = rect{10, 5}
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(r)
	fmt.Println(rp.perim())
	fmt.Println(r)
}
