package router

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mpyszynski/twitter_app/internal/models"
	"net/http"

	"github.com/mpyszynski/twitter_app/internal/pkg/twitter"
)

type Client struct {
	echo *echo.Echo
}

func New(twitterClient twitter.Client) *Client {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	streamHandler := func(ctx echo.Context) error {
		return streamMessages(ctx, &twitterClient)
	}
	e.GET("/stream/:keyWord", streamHandler)

	return &Client{
		echo: e,
	}
}

func (c Client) Start(port string) error {
	if err := c.echo.Start(port); err != nil {
		return err
	}
	return nil
}

func streamMessages(ctx echo.Context, retriever twitter.TweetsRetriever) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().WriteHeader(http.StatusOK)
	keyWord := ctx.Param("keyWord")
	messagesChan := make(chan models.Message)
	enc := json.NewEncoder(ctx.Response())
	var err error
	go func() {
		strmError := retriever.StartStream(keyWord, messagesChan)
		if strmError != nil {
			err = strmError
		}
	}()

	for msg := range messagesChan {
		if err = enc.Encode(msg); err != nil {
			return err
		}
		ctx.Response().Flush()
	}
	close(messagesChan)
	return nil
}
