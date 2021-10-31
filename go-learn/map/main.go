package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["0"] = "Oxygen"
	elements["Ca"] = "Calcium"
	elements["C"] = "Carbon"

	fmt.Println(elements)

	alpha := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	website := map[string]map[string]string{
		"Google": map[string]string{
			"name": "Google",
			"type": "Search",
		},
		"Youtube": map[string]string{
			"name": "Youtube",
			"type": "Video",
		},
	}

	fmt.Println(alpha)
	fmt.Println(website)

}
