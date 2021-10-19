package main

import "fmt"

func main() {
	printValues(1, 2, 3, 4, 5, 6)
}

func printValues(num ...int) {

	for _, val := range num {
		fmt.Println(val)
	}
}
