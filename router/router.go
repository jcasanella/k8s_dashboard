package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/middleware"
	repository "github.com/jcasanella/k8s_dashboard/repository"
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

	k8s := router.Group("/k8s")
	{
		v1 := k8s.Group("/v1")
		{
			v1.GET("/pods/count", repository.CountPods)
			v1.GET("/pods", repository.ListPods)

			v1.GET("/configmaps", middleware.ListConfigMaps)
			v1.GET("/configmaps/count", middleware.CountConfigMaps)

			v1.GET("namespaces", repository.ListNamespaces)
		}
	}

	return router
}
