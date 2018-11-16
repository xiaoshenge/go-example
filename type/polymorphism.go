package main

import "fmt"

type notifyer interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user)notify()  {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	name string
	email string
}

func (a *admin)notify()  {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name,a.email)
}

func main() {
	bill := user{
		name: "bill",
		email: "bill@email.com",
	}

	sendNotify(&bill)


	lisa := admin{"lisa", "lisa@email.com"}
	sendNotify(&lisa)
}

func sendNotify(n notifyer)  {
	n.notify()
}