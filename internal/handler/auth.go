package handler

import (
	"net/http"

	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/gin-gonic/gin"
)

type loginInput struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) verify(c *gin.Context) {
	userId := c.MustGet("userId").(string)

	c.JSON(http.StatusOK, map[string]interface{}{"user_id": userId})
}

func (h *Handler) login(c *gin.Context) {
	var input loginInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Id, input.Password)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *Handler) register(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"user_id": id})
}
