package handlers

import (
	"net/http"
	"xanny-tree-api/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterUrl(c *gin.Context) {
	var data dto.Tree

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Data registered successfully"})
}