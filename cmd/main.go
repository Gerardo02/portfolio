package main

import (
	"github.com/gerardo02/porfolio/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./public")

	handler := handler.Handler{
		Home: handler.HomeHandler{},
	}

	r.GET("/", handler.HandleRenderIndex)

	r.Run()
}
