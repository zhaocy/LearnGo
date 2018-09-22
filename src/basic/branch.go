package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	switch {
	case score < 60:

	}
}

func main() {
	const filename = "a.txt"

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}
