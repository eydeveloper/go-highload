package handler

import (
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPost(c *gin.Context) {
	var input entity.Post
	userId := c.MustGet("userId").(string)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postId, err := h.services.Post.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"post_id": postId})
}

func (h *Handler) updatePost(c *gin.Context) {
	var post entity.Post
	userId := c.MustGet("userId").(string)
	postId := c.Param("id")

	if err := c.BindJSON(&post); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Post.Update(userId, postId, post); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "You have updated the post."})
}

func (h *Handler) getPost(c *gin.Context) {
	postId := c.Param("id")
	post, err := h.services.Post.Get(postId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) deletePost(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	postId := c.Param("id")

	if err := h.services.Post.Delete(userId, postId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "You have deleted the post."})
}
