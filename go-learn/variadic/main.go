package main

import "fmt"

func main() {
	printValues(1, 2, 3, 4, 5, 6)

	x := 0
	increment := func() int {
		x++
		return x
	}

	printNum := func() string {
		return string(increment())
	}

	fmt.Println(printNum(), "test")
	fmt.Println(increment())

}

func printValues(num ...int) {

	for _, val := range num {
		fmt.Println(val)
	}
}
