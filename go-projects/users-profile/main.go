package main

import "fmt"

type User struct {
	name, country string
	age           int
}

type UserList struct {
	User
}

func main() {
	userOne := User{
		name:    "John",
		country: "England",
		age:     23,
	}

	ne := newEmployee(userOne)
	ul := ne

	ul.showUsers()

	modUserOne := User{
		name:    "Sam",
		country: "Germany",
		age:     30,
	}

	ul.modifyUser(modUserOne)

	ul.showUsers()

}

func newEmployee(u User) UserList {
	user := User{
		name:    u.name,
		country: u.country,
		age:     u.age,
	}

	el := UserList{
		User: user,
	}

	return el
}

func (ul UserList) showUsers() {
	fmt.Println("Name:", ul.name, "Country:", ul.country, "Age:", ul.age)

}

func (user *User) modifyUser(newUser User) {
	*user = newUser
}
