package main

import "fmt"

func Variables(){
	var a = "initial" //infers type explicitly
	fmt.Println(a)

	var b,c int = 1, 2 //initialize more than two variables at once and infer types
	fmt.Println(b, c)

	var d =true
	fmt.Println(d)

	var e int //Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	fmt.Println(e)

	f := "apple" //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
	fmt.Println(f)
}

func main(){
	Variables()
}