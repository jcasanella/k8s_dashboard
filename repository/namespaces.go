package repository

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type namespaceObject struct {
	Namespace v1.NamespaceInterface
}

func newNamespace() *namespaceObject {
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	return &namespaceObject{Namespace: clientSet.CoreV1().Namespaces()}
}

func ListNamespaces(c *gin.Context) {
	namespaces, err := newNamespace().Namespace.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.Namespace
	for _, namespace := range namespaces.Items {
		names = append(names, models.Namespace{
			Name:   namespace.Name,
			Status: string(namespace.Status.Phase),
			Age:    namespace.CreationTimestamp.Format(time.RFC3339)})
	}

	c.IndentedJSON(http.StatusOK, names)
}
