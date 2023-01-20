package main

import "fmt"

type User struct {
	name     string
	ocupação string
	country  string
}

func main() {

	users := []User{

		{"John Doe", "gardener", "USA"},
		{"Roger Roe", "driver", "UK"},
		{"Paul Smith", "programmer", "Canada"},
		{"Lucia Mala", "teacher", "Slovakia"},
		{"Patrick Connor", "lojista", " EUA"},
		{"Tim Welson", "programmer", "Canadá"},
		{"Tomas Smutny", "programmer", "Slovakia"},
	}

	var programmers []User

	fmt.Println(programmers)
	for _, user := range users {

		if isProgrammer(user) {
			programmers = append(programmers, user)
		}
	}

	fmt.Println("Programadores:")
	for _, u := range programmers {

		fmt.Println(u)
	}
}

func isProgrammer(usuário User) bool {

	return usuário.ocupação == "programmer"
}
