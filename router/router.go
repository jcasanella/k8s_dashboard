package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	repository "github.com/jcasanella/k8s_dashboard/repository"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "")
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./favicon.ico")
	})

	k8s := router.Group("/v1")
	{
		v1 := k8s.Group("/k8s")
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
