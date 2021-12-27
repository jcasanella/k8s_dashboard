package repository

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ClientNamespace struct {
	Clientset kubernetes.Interface
	Namespace v1.NamespaceInterface
}

type NamespaceOpers interface {
	List() []models.Namespace
	Count() int
	Create()
}

func (c ClientNamespace) List() []models.Namespace {
	namespaces, err := c.Namespace.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.Namespace
	for _, x := range namespaces.Items {
		names = append(names, models.Namespace{
			Name:   x.Name,
			Status: string(x.Status.Phase),
			Age:    x.CreationTimestamp.Format(time.RFC3339)})
	}

	return names
}

func ListNamespaces(c *gin.Context) {
	clientset := c.MustGet("client").(*kubernetes.Clientset)
	client := newClientNamespace(clientset)
	names := client.List()

	c.IndentedJSON(http.StatusOK, names)
}

func newClientNamespace(clientset *kubernetes.Clientset) *ClientNamespace {
	return &ClientNamespace{
		Clientset: clientset,
		Namespace: clientset.CoreV1().Namespaces(),
	}
}
