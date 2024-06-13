package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func (h *Handler) sendMessage(c *gin.Context) {
	token := c.GetHeader("Authorization")
	requestId := c.MustGet("X-Request-ID").(string)

	logrus.Info("Handling send message request with ID: " + requestId)

	client := &http.Client{}
	request, err := http.NewRequest("POST", "http://localhost:8001/api/messages", c.Request.Body)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request.Header.Set("X-Request-ID", requestId)
	request.Header.Set("Authorization", token)
	response, err := client.Do(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var jsonResponse interface{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(response.StatusCode, jsonResponse)
}

func (h *Handler) getMessages(c *gin.Context) {
	token := c.GetHeader("Authorization")
	id := c.Param("id")
	requestId := c.MustGet("X-Request-ID").(string)

	logrus.Info("Handling get messages request with ID: " + requestId)

	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://localhost:8001/api/messages/"+id, nil)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request.Header.Set("Authorization", token)
	request.Header.Set("X-Request-ID", requestId)
	response, err := client.Do(request)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var jsonResponse interface{}

	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(response.StatusCode, jsonResponse)
}
