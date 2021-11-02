package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	var ms2 myStruct
	ms2 = *ms
	fmt.Println(ms2)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}

type myStruct struct {
	foo int
}

func one(xPtr *int) {
	*xPtr = 1
}
