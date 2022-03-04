package handlers

import (
	"gallery-scope/pkg/actions"
	"gallery-scope/pkg/record"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateURLs(c *gin.Context) {
	var body record.Record

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	urls, err := actions.GeneratePreviewURLs(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process request",
			"details": err,
		})
		return
	}

	c.JSON(http.StatusOK, urls)
}

func DownloadGallery(c *gin.Context) {
	var body record.Record

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err := actions.DownloadGallery(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to download gallery",
			"details": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "The gallery will be downloaded",
	})
}
