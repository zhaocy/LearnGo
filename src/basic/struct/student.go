package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	slie := []student{
		{Name: "张飞", Age: 30},
		{Name: "关羽", Age: 35},
		{Name: "刘备", Age: 45},
	}

	m := make(map[string]*student)

	for i, stu := range slie {
		//m[stu.Name] = &stu
		m[stu.Name] = &slie[i]
	}

	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}

func main() {
	pase_student()
}
