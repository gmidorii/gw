package gw

import (
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
