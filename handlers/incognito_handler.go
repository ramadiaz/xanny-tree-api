package handlers

import (
	"net/http"
	"xanny-tree-api/dto"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) UploadIncognitoMessage(c *gin.Context) {
	message := c.Request.FormValue("message")

	if message == "" {
		c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Error: "Message empty"})
		return
	}

	err := h.service.UploadIncognitoMessage(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Status: http.StatusOK, Message: "Incognito message successfully sent"})
}
