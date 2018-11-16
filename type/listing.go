package main

import (
	"example/type/entities"
	"fmt"
)

func main() {
	a := entities.Admin{
		Right: 10,
	}

	a.Name = "adam"
	a.Email = "adam@email.com"

	fmt.Printf("User: %v", a)
}