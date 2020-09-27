package api

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/bst-api/interfaces"
)

type GinWrapper struct {
	*gin.Context
}

func NewGinContextWrapper(c *gin.Context) interfaces.GinContext {
	return &GinWrapper{c}
}
