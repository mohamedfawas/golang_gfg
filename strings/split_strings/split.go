package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Hai, I am Tony, I am from, Kannur!"
	fmt.Println("created string is : ", str1, "having length ", len(str1))
	fmt.Println(strings.SplitAfter(str1, ","))
	str_slice := strings.Split(str1, ",")
	fmt.Println(str_slice)
	// fmt.Println("length after split is : ", len(str_slice))

	// fmt.Printf("type of the variable is %T ", str_slice)
	fmt.Println(len(strings.SplitAfter(str1, ",")))
	fmt.Println(strings.SplitAfterN(str1, ",", 2), "have a length of : ", len(strings.SplitAfterN(str1, ",", 2)))

}
