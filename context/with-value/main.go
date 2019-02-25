package main

import (
	"fmt"
	"context"
)

type User struct{
	Name string
	IsLoggedIn bool
}

type userKeyType int

var userKey userKeyType

func contextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func getUserFromContext(ctx context.Context) *User {
	fmt.Printf("%#v\n", userKey)
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil
	}
	return user
}

func printUser(ctx context.Context)  {
	user := getUserFromContext(ctx)
	fmt.Printf("%#v\n", user)
}

func main()  {
	user := &User{
		Name: "john",
		IsLoggedIn: false,
	}
	ctx := contextWithUser(context.Background(), user)
	printUser(ctx)
}