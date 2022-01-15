package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/domain"
	"github.com/jcasanella/k8s_dashboard/models"
)

func ListConfigMaps(c *gin.Context) {
	k8s := c.MustGet("Clientset").(*models.K8s)
	client := newClientConfigMap(k8s)
	names, err := client.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, names)
}

func CountConfigMaps(c *gin.Context) {
	k8s := c.MustGet("Clientset").(*models.K8s)
	client := newClientConfigMap(k8s)
	count, err := client.Count()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, count)
}

func newClientConfigMap(k8s *models.K8s) *domain.ClientConfigMap {
	return &domain.ClientConfigMap{
		Clientset: k8s.Clientset,
		ConfigMap: k8s.Clientset.CoreV1().ConfigMaps("dma"),
	}
}
