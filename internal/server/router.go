package server

import (
	"github.com/gerardo02/porfolio/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h handler.Handler) {
	r.GET("/", h.Home.HandleRenderHome)
	r.GET("/home", h.Home.HandleRenderIndex)
}
