package main

import (
	"fmt"
)

type rect struct{
	width, height int
}

//constructor
func newRectangle(width, height int) *rect{
	rectangle := rect{
		width: width,
		height: height,
	}
	return &rectangle
}

//methods
//pointer reciever
func(r *rect) area()int{
	return r.width * r.height
}

//value reciever
func(r rect) perimeter()int{
	return 2*(r.width + r.height)
}

func main(){
	r:= rect{10, 12}
	fmt.Println("dimensions:", r.width, "x", r.height)

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())

	r2 := newRectangle(20, 12)
	fmt.Println("dimensions:", r2.width, "x", r2.height)

	fmt.Println("Area:", r2.area())
	fmt.Println("Perimeter:", r2.perimeter())

	square:= struct{
		width, height int
	}{12, 12}
	fmt.Println("Sqaure Dimensions:", square.width, "x", square.height)
}
