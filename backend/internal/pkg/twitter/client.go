package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mpyszynski/twitter_app/internal/config/env"
	"github.com/mpyszynski/twitter_app/internal/models"
)

type TweetsRetriever interface {
	StartStream(hashTag string, msgChannel chan<- models.Message) error
}

type Client struct {
	stream      *twitter.StreamService
	tweetParser twitter.SwitchDemux
}

func New(auth *env.TwitterAuth) Client {
	config := oauth1.NewConfig(auth.ApiKey, auth.ApiSecret)
	token := oauth1.NewToken(auth.ApiToken, auth.ApiTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	demux := twitter.NewSwitchDemux()
	return Client{
		stream:      client.Streams,
		tweetParser: demux,
	}
}

func getTweet(tweet interface{}, demux twitter.SwitchDemux) models.Message {
	var message models.Message
	demux.Tweet = func(tweet *twitter.Tweet) {
		message.Platform = "Twitter"
		message.Text = tweet.Text
		message.User = tweet.User.Name
		message.Nickname = tweet.User.ScreenName
		message.Url = tweet.User.URL
	}
	demux.Handle(tweet)
	return message
}

func (c *Client) StartStream(hashTag string, msgChannel chan<- models.Message) error {
	params := &twitter.StreamFilterParams{
		Track:         []string{hashTag},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := c.stream.Filter(params)
	if err != nil {
		return err
	}
	for message := range stream.Messages {
		tweet := getTweet(message, c.tweetParser)
		msgChannel <- tweet
	}
	close(msgChannel)
	return nil
}
