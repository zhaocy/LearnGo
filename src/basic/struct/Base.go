package main

import "fmt"

/*
Go语言提供了继承，但是采用了组合的语法，我们将其称为匿名组合
*/

type Base struct {
	name string
}

func (b *Base) set(name string) {
	b.name = name
}

func (b Base) get() string {
	return b.name
}

//集成Base
type Derived struct {
	Base
	age int
}

func (d Derived) get() (name string, age int) {
	return d.name, d.age
}

func main() {
	d := new(Derived)
	d.set("gene")
	fmt.Println(d.get())

	b := Base{"张飞"}
	fmt.Println(b.get())
	b.set("关羽")
	fmt.Println(b.get())
}
