package router

import "github.com/gin-gonic/gin"

type IRoute interface {
	Do(route *gin.Engine)
}
