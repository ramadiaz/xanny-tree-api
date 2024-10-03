package handlers

import (
	"database/sql"
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

	err = h.service.RegisterUrl(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Data registered successfully"})
}

func (h *compHandlers) GetUrl(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "url required"})
		return
	}

	origin, err := h.service.GetUrl(url)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, dto.Response{Status: http.StatusNotFound, Error: "Original url not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Body: origin, Message: "Data retrieved successfully"})
}
