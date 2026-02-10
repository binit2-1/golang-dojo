package main

import (
	"fmt"
	"math"
)

const s string = "constant" //can be used outside function body

func Cosnt() {
	fmt.Println(s)

	const n = 50000000000
	
	const d = 3e20/n //calculates with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) //numeric constant has no type until explicit conversion

	fmt.Println(math.Sin(n)) //A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.

}

func main(){
	Cosnt()
}