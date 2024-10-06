package handlers

import (
	"net/http"
	"xanny-tree-api/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) AddLike(c *gin.Context) {
	data, err := h.service.AddLike()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: data})
}