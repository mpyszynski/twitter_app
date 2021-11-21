package mocks

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/mpyszynski/twitter_app/internal/models"
)

type MockStream struct {

}

func (ms *MockStream) Filter(params *twitter.StreamFilterParams) (*twitter.Stream, error) {
	messages := []models.Message{
		{
			Platform: "foo",
			User: "bar",
			Text: "foobar",
			Url: "baz",
			Nickname: "fizz",
		},
		{
			Platform: "fizz",
			User: "baz",
			Text: "bar",
			Url: "foobar",
			Nickname: "foo",
		},
	}
	stream := twitter.Stream{
		Messages: make(chan interface{}),
	}
	go func() {
		for msg := range messages {
			stream.Messages <- msg
		}
		close(stream.Messages)
	} ()
	return &stream, nil
}