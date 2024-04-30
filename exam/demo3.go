// count freq of each word in a sentence
package main

import (
	"fmt"
	"strings"
)

func main() {
	sent := "hai hai hai I am from Kerala nice, to meet you hai"

	arr_str := strings.Split(sent, " ")

	frq_map := make(map[string]int)

	for _, v := range arr_str {
		frq_map[v]++
	}
	fmt.Println(frq_map)

}
