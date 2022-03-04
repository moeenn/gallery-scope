package server

import (
	"gallery-scope/pkg/server/handlers"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	// server static files
	{
		r.Static("/public", "./public")
		r.StaticFile("/", "./public/index.html")
		r.StaticFile("/favicon.ico", "./public/img/favicon.ico")
	}

	// resource api endpoints
	{
		r.POST("/api/urls", handlers.GenerateURLs)
		r.POST("/api/download", handlers.DownloadGallery)
	}

	return r
}
