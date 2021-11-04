package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	y := s
	y[0] = 5
	fmt.Println(s, y)

	mapstring := map[string]int{
		"hello": 3,
	}

	m := mapstring
	m["y"] = 100
	fmt.Println(mapstring, m)
}
