package main

import (
	"fmt"
	"math"
)

var (
	aa = 2
	ss = "kk"
	bb = true
)

func variableZeroValue() {
	//声明变量
	var a int
	var s string
	var m map[int]int
	//未赋值的变量会有默认值 数值类型是0 字符串类型是空串 引用类型是nil
	fmt.Println(a, s, m)         //0
	fmt.Printf("%d %s \n", a, s) //0
	fmt.Printf("%d %q \n", a, s) //0 ""
}

//变量赋初始值
func variableInitialValue() {
	var a, b = 3, 4 //等价var a, b int = 3,4
	var s = "abc"   //等价var s string = "abc"
	var j, k, l = 2, "zz", true
	fmt.Println(a, b, s) //3 4 abc
	fmt.Println(j, k, l) //2 "zz" true

}

func variableShorter() {
	a, b, c, s := 3, 4, true, "defer"
	fmt.Println(a, b, c, s)
}

func consts() {
	const (
		j, k = 5, 6
	)
	const filename = "ab.txt"
	const a, b = 3, 4
	c := math.Sqrt(a) //常量a没有指定类型，在这里就是float64类型
	fmt.Println(c)
}

func enums() {
	const (
		app = iota + 1
		_
		java
		python
		golang
	)

	fmt.Println(app, java, python, golang) //1 3 4 5
}

func main() {

	variableZeroValue()
	variableInitialValue()
	variableShorter()

	consts()
	enums()
}
