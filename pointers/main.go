package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(x *int) {
	*x = 0
}
