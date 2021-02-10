package main

import "fmt"

func main() {
	num := []int{2, 4, 6, 6, 4}

	res := 0

	for _, s := range num {
		fmt.Println(res)
		res = res ^ s
	}
	fmt.Println(res)

	// THIS GIVES SURPRISING RESULT
	// SHOWS MY LIMITED KNOWLEDGE OF BITWISE OPERATORS

}
