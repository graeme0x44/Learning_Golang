package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	person_1 := person{"Fred", 100}

	struct_pointer := &person{"Bob", 20}
	fmt.Println(struct_pointer.age)

	fmt.Println(person_1.name)

	sp := &person_1
	fmt.Println(sp.age)

	person_3 := person{age: 32, name: "Test"}
	fmt.Println(person_3)
}

