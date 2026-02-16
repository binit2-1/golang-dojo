package main

	

import "fmt"


func add(a int, b int)int{
	sum:=a+b
	return sum
}

func addAdd(a, b, c int)int{
	return a+b+c
}

func vals()(int, int){
	return 3, 4 //Multiple returns
}

func variadicSum(nums ...int)int{
	total:=0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)

	return total
}

func main(){
	fmt.Println(add(3, 4))
	fmt.Println(addAdd(3, 4, 5))

	a, b:= vals()
	fmt.Println(a, b)
	
	nums := []int{1,2,3,4,5}

	variadicSum(nums...)
}