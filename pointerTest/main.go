package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	persons := []*Person{new(Person), new(Person), new(Person)}

	setFunc(persons)
	fmt.Println(persons[0]) // &{fool 1}
	fmt.Println(persons[1]) // &{fool 2}
	fmt.Println(persons[2]) // &{fool 3}
}

func setFunc(persons []*Person) {
	for _, p := range persons {
		p.name = "fool"
		p.age = 1
	}
	// そもそもスライスには反映しない
	fmt.Println(persons) // [{ 0} { 0} { 0}]
}
