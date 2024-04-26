package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) follow(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id := c.Param("id")

	if err := h.services.Following.Follow(id, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "You have followed the user."})
}

func (h *Handler) unfollow(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	id := c.Param("id")

	if err := h.services.Following.Unfollow(id, userId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{"message": "You have unfollowed the user."})
}
