package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCategories(c *gin.Context){
	lang := c.Query("lang")

	categories, err := h.service.GetCategories(lang)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}
