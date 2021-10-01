package router

import (
	"github.com/gin-gonic/gin"

	"github.com/jcasanella/k8s_dashboard/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", middleware.GetAlbums)
	router.GET("/albums/:id", middleware.GetAlbumByID)
	router.POST("/albums", middleware.PostAlbums)

	return router
}
