package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/mpyszynski/twitter_app/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTweet(t *testing.T){
	testTweet := twitter.Tweet{
		Text: "foo",
		User: &twitter.User{
			Name: "bar",
			ScreenName: "baz",
			URL: "fizz",
		},
	}
	demux := twitter.NewSwitchDemux()

	tweet := getTweet(testTweet, demux)
	assert.IsType(t, models.Message{}, tweet)
}
