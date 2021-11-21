package router

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mpyszynski/twitter_app/internal/mocks"
	"github.com/mpyszynski/twitter_app/internal/models"
	"github.com/stretchr/testify/assert"
	"io"
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
	err := handler.getMessagesFromStream(c)
	assert.Nil(t, err)
	reader := bufio.NewReader(rec.Body)
	var msgList []models.Message
	for {
		var msg models.Message
		data, err := reader.ReadBytes('\n')
		fmt.Println(string(data))
		if err == io.EOF {
			break
		}
		if err != nil {
			assert.Fail(t, err.Error())
		}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		msgList = append(msgList, msg)
	}
	assert.Equal(t, msgList[0].Text, "foobar")
	//assert.Equal(t, msgList[1].Text, "foo")

}