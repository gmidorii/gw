package gw

import (
	"errors"

	sl "github.com/nlopes/slack"
)

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
	if s.token == "" {
		return errors.New("failed send message: token is empty")
	}
	client := sl.New(s.token)
	at := sl.Attachment{
		//Color: "",
		Title: title,
		Text:  body,
	}
	_, _, err := client.PostMessage(dest, "", sl.PostMessageParameters{
		Attachments: []sl.Attachment{at},
	})
	return err
}
