package mocks

import (
	"fmt"
	"github.com/mpyszynski/twitter_app/internal/models"
	"time"
)

type MockTwitterClient struct {

}

func (mt *MockTwitterClient) StartStream(hashTag string, msgChannel chan<- models.Message) error {
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
	for _, msg := range messages {
		msgChannel <- msg
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
	close(msgChannel)
	return nil
}