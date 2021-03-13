package router

import (
	"shorturl/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))

	r.GET("/:shortname", api.HandleRedirect)

	r.POST("/api/login", api.LoginHandler)

	// need login to generate short name
	authorized := r.Group("/api")
	authorized.Use(api.AuthRequired)
	{
		authorized.POST("/", api.AddShortName)
	}

	return r
}
