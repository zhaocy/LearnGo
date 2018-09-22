package inter

import "fmt"

//Human实现SayHi方法
func (h Human)SayHi() string{
	say:="Hi, I am "+h.Name+" you can call me on "+h.Phone+" \n"
	return say
}

func (s Human)Sing(lyrics string){
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的Sing方法
func (e Employee)Sing(lyrics string){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s . My Sing Now!! %s \n", e.Name,
		e.Company, e.Phone, lyrics)
}
