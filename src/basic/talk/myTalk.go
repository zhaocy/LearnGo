package talk

import "fmt"

type myTalk string

func (myTalk *myTalk) Hello(userName string)string{
	return "hello"
}

func(talk *myTalk) Talk(heard string, end bool, err error ){
	fmt.Println("")
}
