package server

import (
	"gallery-scope/pkg/server/handlers"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// server static files
	{
		r.Static("/public", "./public")
		r.StaticFile("/", "./public/index.html")
		r.StaticFile("/favicon.ico", "./public/img/favicon.ico")
	}

	// resource api endpoints
	api := r.Group("/api")
	{
		api.POST("/urls", handlers.GenerateURLs)
		api.POST("/download", handlers.DownloadGallery)
	}

	return r
}
