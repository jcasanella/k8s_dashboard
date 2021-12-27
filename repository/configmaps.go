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

func (c ClientConfigMap) Count() int {
	configmaps, err := c.Configmap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return len(configmaps.Items)
}

func ListConfigMaps(c *gin.Context) {
	clientset := c.MustGet("Clientset").(*kubernetes.Clientset)
	client := newClientConfigMap(clientset)
	names := client.List()
	c.IndentedJSON(http.StatusOK, names)
}

func CountConfigMaps(c *gin.Context) {
	Clientset := c.MustGet("Clientset").(*kubernetes.Clientset)
	client := newClientConfigMap(Clientset)
	numConfigMaps := client.Count()
	c.IndentedJSON(http.StatusOK, numConfigMaps)
}

func newClientConfigMap(clientset *kubernetes.Clientset) *ClientConfigMap {
	return &ClientConfigMap{
		Clientset: clientset,
		Configmap: clientset.CoreV1().ConfigMaps("dma"),
	}
}
