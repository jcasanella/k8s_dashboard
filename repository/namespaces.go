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
	k8s := c.MustGet("client").(*models.K8s)
	clientNamespace := newClientNamespace(k8s)
	names := clientNamespace.List()

	c.IndentedJSON(http.StatusOK, names)
}

func newClientNamespace(k8s *models.K8s) *ClientNamespace {
	return &ClientNamespace{
		Clientset: k8s.Clientset,
		Namespace: k8s.Clientset.CoreV1().Namespaces(),
	}
}
