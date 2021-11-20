package router

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mpyszynski/twitter_app/internal/models"
	"github.com/mpyszynski/twitter_app/internal/pkg/twitter"
	"net/http"
)

type handler struct {
	twitterClient twitter.TweetsRetriever
}

func (h *handler) getMessagesFromStream(ctx echo.Context) error {
	errorChan := make(chan error)
	messagesChan := make(chan models.Message)

	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx.Response().WriteHeader(http.StatusOK)

	keyWord := ctx.Param("keyWord")

	enc := json.NewEncoder(ctx.Response())

	go func() {
		err := h.twitterClient.StartStream(keyWord, messagesChan)
		if err != nil {
			errorChan <- err
		}
	}()

	for msg := range messagesChan {
		select {
		case err := <-errorChan:
			return err
		case <-messagesChan:
			if err := enc.Encode(msg); err != nil {
				return err
			}
			ctx.Response().Flush()
		}
	}
	close(errorChan)
	return nil
}
