package main

import "fmt"

type User struct {
	name, country string
	age           int
}

type UserList []string

func main() {
	emOne := User{
		name:    "John",
		country: "England",
		age:     23,
	}

	ne := newEmployee(emOne)
	fmt.Println(ne)
	ul := UserList{}

	ul.showUsers()

}

func newEmployee(u User) UserList {
	user := User{
		name:    u.name,
		country: u.country,
		age:     u.age,
	}

	el := UserList{}

	el = append(el, "Name: "+user.name+"Country: "+user.country+"Age: "+string(user.age))

	return el
}

func (ul UserList) showUsers() {
	for _, user := range ul {
		fmt.Println(user)
	}
}
