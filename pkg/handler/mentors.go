package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getMentors(c *gin.Context) {
	mentors, err := h.service.GetMentors()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, mentors)
}

func (h *Handler) getMentorByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	mentor, err := h.service.GetMentorByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, mentor)
}
