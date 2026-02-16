package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 21
	return &p
}

func main() {
	fmt.Println(person{"Bob", 21})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"}) //not mentioned = zero valued
	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age) //mutable structs

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
