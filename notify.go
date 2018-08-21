package gw

import (
	"errors"

	sl "github.com/nlopes/slack"
)

type Notifier interface {
	Send(title, dest, body string, ok bool) error
}

type slack struct {
	token    string
	okColor  string
	errColor string
}

func NewSlack(token, okColor, errColor string) Notifier {
	return slack{
		token:    token,
		okColor:  okColor,
		errColor: errColor,
	}
}

func (s slack) Send(title, dest, body string, ok bool) error {
	if s.token == "" {
		return errors.New("failed send message: token is empty")
	}
	client := sl.New(s.token)
	at := sl.Attachment{
		Color: "#006400",
		Title: title,
		Text:  body,
	}

	if ok {
		at.Color = s.okColor
	} else {
		at.Color = s.errColor
	}

	_, _, err := client.PostMessage(dest, "", sl.PostMessageParameters{
		Attachments: []sl.Attachment{at},
	})
	return err
}
