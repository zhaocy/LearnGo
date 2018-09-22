package main

import (
	"basic/inter"
	"fmt"
	"reflect"
)

func main() {

	student:=inter.Student{inter.Human{"gene",35,"16623120527"},"EastSouth",30.50}
	employee:=inter.Employee{inter.Human{"Tome",40,"13872370909"},"QQ",50.02}
	//
	//
	var iMen inter.Men

	//Student
	iMen = student
	fmt.Print(iMen.SayHi())
	iMen.Sing("我是一只小🐝")

	//Employee
	iMen = employee
	iMen.Sing("一二三四 一二三四")
	fmt.Println(iMen.SayHi())


	//slice Men
	x:= make([]inter.Men,2)
	x[0] ,x[1] = student,employee
	for _, value:= range x  {
		value.Sing("黄河在咆哮...")
	}



	//空接口空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），
	// 但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。
	var a interface{}

	i:=6
	//s:="hello"
	a = i
	//a = s

	//if v,ok:= a.(int);ok{
	//	fmt.Println(v)
	//}
	fmt.Println(a, reflect.TypeOf(a),a.(int))

}
