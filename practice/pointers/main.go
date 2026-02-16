package main

import "fmt"

func copyVal(ival int){
	ival = 0
}

func adressVal(pval *int){
	*pval = 0;
}

func main() {
	i := 1
    fmt.Println("initial:", i)

	copyVal(i)
    fmt.Println("zeroval:", i)

	adressVal(&i)
    fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
} 