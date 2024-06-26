package main

import (
	"fmt"

	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

type TokenProvider interface {
	Token() string
	SetUserID(userID string)
}

type Request[T any] struct {
	JWT    string
	UserID string
	Data   T
}

func (r *Request[T]) Token() string {
	return r.JWT
}

func (r *Request[T]) SetUserID(userID string) {
	r.UserID = userID
}

func main() {
	req := transfer.Request[string]{
		JWT:  "test",
		Data: "test",
	}
	var _ TokenProvider = (*Request[string])(nil)

	check(&req)
	check2(&req)
}

func check(args interface{}) {
	_, ok := args.(TokenProvider)
	if !ok {
		fmt.Println("Аргумент не реализует интерфейс TokenProvider")
		return
	}
}

func check2(args interface{}) {
	_, ok := args.(*transfer.Request[string])

	if !ok {
		fmt.Println("Аргумент не реализует  transfer.Request[string]")
		return
	}
}
