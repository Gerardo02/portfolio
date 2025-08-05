package home

import (
	"github.com/gerardo02/porfolio/internal/handler/utils"
	"github.com/gerardo02/porfolio/internal/view/home"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func (h *HomeHandler) HandleRenderHome(c *gin.Context) {
	err := utils.Render(c, home.Home())
	if err != nil {
		c.String(500, "template rendering error: %v", err)
	}
}

func (h *HomeHandler) HandleRenderIndex(c *gin.Context) {
	err := utils.Render(c, home.Index())
	if err != nil {
		c.String(500, "template rendering error: %v", err)
	}
}
