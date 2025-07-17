package handler

import (
	"github.com/gerardo02/porfolio/view/home"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func (h *HomeHandler) HandleRenderHome(c *gin.Context) {
	err := render(c, home.Home())
	if err != nil {
		c.String(500, "template rendering error: %v", err)
	}
}
