package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *Handler) getFeed(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	feed, err := h.services.Feed.Get(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, feed)
}

func (h *Handler) getRealTimeFeed(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer wsConnection.Close()

	deliveries, err := h.services.Feed.GetRealTime(userId)

	if err == nil {
		for delivery := range deliveries {
			err := wsConnection.WriteMessage(websocket.TextMessage, delivery.Body)
			if err != nil {
				return
			}
		}
	}
}
