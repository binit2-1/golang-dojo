package main

import "fmt"


//slice -> dynamic array
//most used construct in go
//+ useful methods in array
//type can be string, int, etc..
func main(){
	var s[] string
	fmt.Println("uniunit: ", s, s == nil, len(s) == 0) //slcies by default are created with values = nil and len = 0

	s = make([]string, 3, 4) //third parameter to initialize capacity too, keep len zero as a proper way 
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s)) //len = capacity by default

	s[0] = "hi"
	s[1] = "hello"
	s[2] = "bye"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "ok")
	s = append(s, "haan", "oye")
	fmt.Println("apd: ", s)
	fmt.Println("len: ", len(s))

	//slice copying
	c := make([]string, 3)
	copy(c,s)
	fmt.Println("copy: ", c)

	d := []string{"hi", "hello"}
	fmt.Println("dcl: ", d)

	//slicing
	fmt.Println(s[1:])// including 1st index till the last index
	fmt.Println(s[:])//all
	fmt.Println(s[:4])// excluding the 4th index and further, rest all print


	twoDim := make([][]int, 3)
	for i:=range 3{
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := range innerLen{
			twoDim[i][j] = i+j
		}
	}
	fmt.Println(twoDim)
}