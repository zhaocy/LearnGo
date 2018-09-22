package main

import "fmt"

type Integer int

type Person struct {
	name    string
	age     int
	country string
}

func (i Integer) less(a Integer) bool {
	return i < a
}

func NewPerson(name string, age int, country string) *Person {
	return &Person{name: name, age: age, country: country}
}

func main() {

	var a Integer = 1

	if a.less(2) {
		fmt.Println("Less than 2")
	}

	p := NewPerson("gene", 35, "China")
	fmt.Println(p.name)
}
