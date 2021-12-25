package repository

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type configmapObject struct {
	Configmap v1.ConfigMapInterface
}

func newConfigmap() *configmapObject {
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	return &configmapObject{Configmap: clientSet.CoreV1().ConfigMaps("dma")}
}

func ListConfigMaps(c *gin.Context) {
	configMaps, err := newConfigmap().Configmap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.ConfigMap
	for _, configMap := range configMaps.Items {
		names = append(names, models.ConfigMap{Name: configMap.Name})
	}

	c.IndentedJSON(http.StatusOK, names)
}

func CountConfigMaps(c *gin.Context) {
	configMaps, err := newConfigmap().Configmap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, len(configMaps.Items))
}
