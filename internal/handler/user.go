package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.services.User.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) searchUser(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")

	users, err := h.services.User.Search(firstName, lastName)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
