package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/albums", middleware.GetAlbums)
		v1.GET("/albums/:id", middleware.GetAlbumByID)
		v1.POST("/albums", middleware.PostAlbums)
	}

	router.Static("/js", "./js")
	router.LoadHTMLGlob("templates/*.tmpl")
	v2 := router.Group("/v2")
	{
		v2.GET("/albums", middleware.GetTitle)
	}

	return router
}
