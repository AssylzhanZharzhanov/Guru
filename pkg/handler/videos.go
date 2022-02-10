package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getVideos(c *gin.Context) {
	courseId, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, courseId)
}
