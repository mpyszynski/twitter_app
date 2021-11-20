package router

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mpyszynski/twitter_app/internal/mocks"
	"github.com/mpyszynski/twitter_app/internal/models"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	logger = log.Logger{}
)
func TestStreamHandler(t *testing.T){
	mockTwitterStreamer := mocks.MockTwitterClient{}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/stream/test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := handler{twitterClient: &mockTwitterStreamer}
	res := handler.getMessagesFromStream(c)
	assert.Nil(t, res)
	msg := models.Message{}
	json.NewDecoder(rec.Body).Decode(&msg)

}