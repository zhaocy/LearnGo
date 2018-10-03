package main

import "fmt"

func adder() func(i int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(v int) iAdder {
	return func(i int) (int, iAdder) {
		return v + i, adder2(v + i)
	}
}

func main() {

	/*a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+...%d = %d\n", i, a(i))
	}*/


	a:= adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s,a = a(i)
		fmt.Printf("0+...%d = %d\n", i, s)
	}

}
