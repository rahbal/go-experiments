package main

import "fmt"

func main() {

	fmt.Println(Solution("werty"))

}

func Solution(word string) string {

	wo := ""
	for pos, char := range word {

		pos += pos

		wo = string(char) + wo

	}

	return wo

}
