package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Android struct {
	Person
	Model string
}

func (a Android) talk() {
	fmt.Println("Hi, My name is : ", a.Person.Name)
	fmt.Println("I call and talk using the phone model : ", a.Model)
}

func main() {
	p := Person{"fawas", 26}
	a := Android{p, "a123"}
	fmt.Println(a)
	a.talk()
}
