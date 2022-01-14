package repository

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jcasanella/k8s_dashboard/domain"
	"github.com/jcasanella/k8s_dashboard/models"
)

func ListNamespaces(c *gin.Context) {
	k8s := c.MustGet("client").(*models.K8s)
	client := newClientNamespace(k8s)
	ns, err := client.List()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, map[string]interface{}{"error": strings.Join([]string{err.Error()}, "; ")})
	}

	c.IndentedJSON(http.StatusOK, ns)
}

func CountNamespaces(c *gin.Context) {
	k8s := c.MustGet("client").(*models.K8s)
	client := newClientNamespace(k8s)
	count, err := client.Count()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, map[string]interface{}{"error": strings.Join([]string{err.Error()}, "; ")})
	}

	c.IndentedJSON(http.StatusOK, count)
}

func newClientNamespace(k8s *models.K8s) *domain.ClientNamespace {
	return &domain.ClientNamespace{
		Clientset: k8s.Clientset,
		Namespace: k8s.Clientset.CoreV1().Namespaces(),
	}
}
