package main

import (
	"github.com/mpyszynski/twitter_app/internal/config/env"
	"github.com/mpyszynski/twitter_app/internal/pkg/twitter"
	"github.com/mpyszynski/twitter_app/internal/pkg/router"
	"log"
)

func main() {
	config, err := env.LoadConfig(".", "secrets")
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}
	twitterClient := twitter.New(config.TwitterAuth)
	echo := router.New(twitterClient)
	err = echo.Start(":3001")
	if err != nil {
		log.Fatalf("failed to start router server %v", err)
	}
}
