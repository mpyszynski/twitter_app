package main

import (
	"fmt"
	"github.com/mpyszynski/twitter_app/internal/config/env"
	"log"
)

func main() {
	config, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}
	fmt.Println(config.TwitterAuth.AppID)
}
