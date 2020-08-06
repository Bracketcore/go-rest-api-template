package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundResponse return
func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found."})
}

// ForbiddenResponse return
func ForbiddenResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view this resource."})
}

// BadRequestResponse return
func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

// InputBadRequestResponse return
func InputBadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

// InternalServerErrorResponse return
func InternalServerErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

// DuplicateEntryErrorResponse return
func DuplicateEntryErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, gin.H{"error": message})
}
