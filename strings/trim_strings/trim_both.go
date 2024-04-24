package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "   Hai I am Tony           "
	fmt.Println(str1)
	fmt.Println("after trim ")
	fmt.Println(strings.Trim(str1, " "))
	fmt.Println(strings.TrimRight(str1, " "))
	fmt.Println(strings.TrimSpace(str1))
}
