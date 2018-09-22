package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//Fileds
	fmt.Println(strings.Fields("我 看 football")) //[我 看 football]

	//split
	fmt.Println(strings.Split("我-们-不-一-样", "-")) //[我 们 不 一 样]

	s := "你好"           // UTF-8
	fmt.Println(len(s)) //结果：6

	fmt.Println(len([]rune(s))) // 2

	slice := make([]string, 10)
	for i := 0; i < 5; i++ {
		slice[i] = "t"
	}
	fmt.Println(slice)
	//join
	ss := strings.Join(slice, "棒") //t棒t棒t棒t棒t棒棒棒棒棒
	fmt.Println(ss, len(ss), len([]rune(ss)))

	//Contains
	fmt.Println(strings.Contains("IamXiaoming", "X"))   //true
	fmt.Println(strings.Contains("IamXiaoming", "x m")) //false
	fmt.Println(strings.Contains("IamXiaoming", "am"))  //true
	fmt.Println(strings.Contains("IamXiaoming", ""))    //true

	//Index
	fmt.Println(strings.Index("abc", "e")) //-1
	fmt.Println(strings.Index("abc", "c")) //2

	/*
		通过buffer进行组装拼接，使用buffer是优先创建一个缓冲区，然后向缓冲区中写入数据，类似Java中的StringBuffer
	*/
	var b bytes.Buffer
	b.Grow(22)

	for i := 0; i < 22; i++ {
		b.WriteString("c")
	}
	fmt.Println(b.String())

	fmt.Println(nonRepeatingSubStr("abcbacae"))
}

func nonRepeatingSubStr(s string) string {
	last := make(map[rune]int)
	startTemp := 0
	start := 0
	maxLength := 0
	var res []rune
	st := []rune(s)
	for i, ch := range st {

		if lastI, ok := last[ch]; ok && lastI >= startTemp {
			startTemp = i
		}

		if i+1-startTemp > maxLength {
			maxLength = i + 1 - startTemp
			start = startTemp
		}

		last[ch] = i
	}

	for j := start; j < start+maxLength; j++ {
		temp := st[j]
		res = append(res, temp)
	}
	result := string(res)
	return result
}
