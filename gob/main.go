package main

import (
	"encoding/gob"
	"os"
)

type User struct{
	Name string
	Password string
}

func main()  {
	user := User{
		Name: "adam",
		Password: "1234",
	}
	file, _ := os.Create("user.gob")
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(user)
}