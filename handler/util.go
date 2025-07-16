package handler

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, component templ.Component) error {
	return component.Render(c.Request.Context(), c.Writer)
}
