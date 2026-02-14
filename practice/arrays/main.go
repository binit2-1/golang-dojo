package main

import "fmt"

func main(){
	var a[5] int;
	fmt.Println("emp:", a) //by default zero valued [0 0 0 0 0]

	a[4] = 100 
	fmt.Println("set:", a) //set: [0 0 0 0 100]
    fmt.Println("get:", a[4]) //get: 100

	b := [5]int{1, 2, 3, 4, 5} //declare and initialize 
	fmt.Println("dcl: ", b)

	c := [...]int{1,2,3,4,5} //counts elements by itself
	fmt.Println("dcl: ", c)

	d := [...]int{1, 2:200, 3:400, 5} //puts 400 in 2nd and 3rd index
	fmt.Println("dcl: ", d)

	var twoDim[2][3] int
	for i:= range 2{
		for j:= range 3{
			twoDim[i][j] =i+j
		}
	}
	fmt.Println("2d: ", twoDim)

	twoD2 := [2][3] int {
		{1,2,3},
		{2,3,4},
	} 
	fmt.Println("ini 2d: ", twoD2)
}