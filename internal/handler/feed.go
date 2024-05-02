package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
