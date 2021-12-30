package repository

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ClientConfigMap struct {
	Clientset kubernetes.Interface
	Configmap v1.ConfigMapInterface
}

type ConfigMapOpers interface {
	List() []models.ConfigMap
	Count() int
	Create()
}

func (c ClientConfigMap) List() []models.ConfigMap {
	configmaps, err := c.Configmap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.ConfigMap
	for _, x := range configmaps.Items {
		names = append(names, models.ConfigMap{Name: x.Name})
	}

	return names
}

func ListConfigMaps(c *gin.Context) {
	k8s := c.MustGet("Clientset").(*models.K8s)
	clientConfigMap := newClientConfigMap(k8s)
	names := clientConfigMap.List()
	c.IndentedJSON(http.StatusOK, names)
}

func (c ClientConfigMap) Count() int {
	configmaps, err := c.Configmap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return len(configmaps.Items)
}

func CountConfigMaps(c *gin.Context) {
	k8s := c.MustGet("Clientset").(*models.K8s)
	clientConfigMap := newClientConfigMap(k8s)
	numConfigMaps := clientConfigMap.Count()
	c.IndentedJSON(http.StatusOK, numConfigMaps)
}

func newClientConfigMap(k8s *models.K8s) *ClientConfigMap {
	return &ClientConfigMap{
		Clientset: k8s.Clientset,
		Configmap: k8s.Clientset.CoreV1().ConfigMaps("dma"),
	}
}
