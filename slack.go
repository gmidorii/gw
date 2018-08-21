package main

import "fmt"

type Notifier interface {
	Send(title, dest, body string) error
}

type slack struct {
	token string
}

func NewSlack(token string) Notifier {
	return slack{
		token: token,
	}
}

func (s slack) Send(title, dest, body string) error {
	fmt.Printf("%v, %v, %v, %v", s.token, title, dest, body)
	return nil
}
