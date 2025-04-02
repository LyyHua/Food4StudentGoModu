package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *application) writeJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func (app *application) readJSON(c *gin.Context, data interface{}) error {
	// Set max body size (optional)
	maxBytes := 1024 * 1024 // one megabyte
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxBytes))

	if err := c.ShouldBindJSON(&data); err != nil {
		return err
	}

	return nil
}

func (app *application) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}

	c.JSON(statusCode, payload)
}
