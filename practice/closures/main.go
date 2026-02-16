package main

import "fmt"

func Closure() func() int{
	i := 0;

	return func() int{
		i++
		return i
	}
}

func main(){
	increment := Closure()
	fmt.Println(increment())
	fmt.Println(increment())
	
}

//closure closes a another function into it and nit keeps the modifications applied to a variable that is inside its scope (i) retain it even after it is removed from call stack