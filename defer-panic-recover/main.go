package main

import "fmt"

func main() {
	first()
	// It will run after all the functions
	defer second()
	recover()
	third()
	panic("Stop!")

}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}
