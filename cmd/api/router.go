package api

import (
	"github.com/gerardo02/porfolio/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h handler.Handler) {
	r.GET("/", h.Home.HandleRenderHome)
}
