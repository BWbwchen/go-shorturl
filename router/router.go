package router

import (
	"shorturl/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	r.GET("/:shortname", api.HandleRedirect)
	r.POST("/api", api.AddShortName)

	return r
}
