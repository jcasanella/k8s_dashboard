package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl")

	v1 := router.Group("/v1")
	{
		v1.GET("/albums", middleware.GetAlbums)
		v1.GET("/albums/:id", middleware.GetAlbumByID)
		v1.POST("/albums", middleware.PostAlbums)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/albums", middleware.GetTitle)
	}

	k8sV1 := router.Group("/k8s/v1")
	{
		k8sV1.GET("/count", middleware.CountPods)
	}

	return router
}
