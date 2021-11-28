package router

import (
	"github.com/gin-gonic/gin"
	repository "github.com/jcasanella/k8s_dashboard/repository"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl")

	k8s := router.Group("/k8s")
	{
		v1 := k8s.Group("/v1")
		{
			v1.GET("/pods/count", repository.CountPods)
			v1.GET("/pods", repository.ListPods)

			v1.GET("/configmaps", repository.ListConfigMaps)
			v1.GET("/configmaps/count", repository.CountConfigMaps)

			v1.GET("namespaces", repository.ListNamespaces)
		}
	}

	return router
}
