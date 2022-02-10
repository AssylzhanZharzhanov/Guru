package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getCourses(c *gin.Context) {
	lang := c.Query("lang")
	
	courses, err := h.service.GetCourses(lang)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getTrendingCourses(c *gin.Context) {
	lang := c.Query("lang")

	courses, err := h.service.GetTrendingCourses(lang)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getComingSoonCourses(c *gin.Context) {
	lang := c.Query("lang")

	courses, err := h.service.GetComingSoonCourses(lang)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) getCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	course, err := h.service.GetCourseByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, course)
}
