package main

import "fmt"

type author struct {
	name   string
	branch string
}

func (a *author) show(abranch string) {
	(*a).branch = abranch
}

func main() {
	auth1 := author{
		name:   "virat kohli",
		branch: "cse",
	}

	fmt.Println("before : ", auth1)

	auth1.show("electrical")
	fmt.Println("after : ", auth1)
}
