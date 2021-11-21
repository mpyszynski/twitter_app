package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mpyszynski/twitter_app/internal/pkg/twitter"
)

// Client for echo API
type Client struct {
	echo *echo.Echo
}

// New creates new echo client
func New(twitterClient twitter.Client) *Client {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	handler := handler{
		twitterClient: &twitterClient,
	}

	e.GET("/stream/:keyWord", handler.getMessagesFromStream)

	return &Client{
		echo: e,
	}
}

// Start echo REST API
func (c Client) Start(port string) error {
	if err := c.echo.Start(port); err != nil {
		return err
	}
	return nil
}

