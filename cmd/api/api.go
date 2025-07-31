package api

import (
	"github.com/gerardo02/porfolio/handler"
	"github.com/gin-gonic/gin"
)

func Start() error {
	r := gin.Default()

	r.Static("/static", "./public")

	handler := handler.Handler{
		Home: handler.HomeHandler{},
	}

	NewRouter(r, handler)

	if err := r.Run(); err != nil {
		return err
	}

	return nil
}
