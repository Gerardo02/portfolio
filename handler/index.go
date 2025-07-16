package handler

import (
	"github.com/gerardo02/porfolio/view/index"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Home HomeHandler
}

func (Handler) HandleRenderIndex(c *gin.Context) {
	err := render(c, index.Index())
	if err != nil {
		c.String(500, "template rendering error: %v", err)
	}
}
