package main

import "fmt"

//Only for loop in golang
func main(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1 
	}

	for j:=0; j <= 3; j++ {
		fmt.Println(i)
		j += 1 
	}

	for i:= range 3{
		fmt.Println("range",i)
		i++
	}

	for{
		fmt.Println("loop") //runs till break
		break
	}

	for n:= range 6{
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}
}