package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_base_project/app/base"
	instant2 "go_base_project/app/controllers/instant"
)

type web struct{}

func (web) Do(router *base.Router) {
	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello world!")
	})

	instant := router.Group("/instant")
	instant.POST("pricing", instant2.PricingController{}.Index)
}

func Init() web {
	return web{}
}
