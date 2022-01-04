package repository

import (
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/domain"
	"github.com/jcasanella/k8s_dashboard/models"
)

func ListPods(c *gin.Context) {
	// limitQuery := c.DefaultQuery("limit", "10")
	// limitValue, err := strconv.ParseInt(limitQuery, 10, 64)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// listOptions := &meta.ListOptions{
	// 	Limit: limitValue,
	// }

	// pods, err := newPod().Pod.List(context.TODO(), *listOptions)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var names []models.Pod
	// for _, pod := range pods.Items {
	// 	names = append(names, models.Pod{Name: pod.Name, Continue: pods.ListMeta.Continue, RemainingItemCount: *pods.ListMeta.RemainingItemCount})
	// }

	// client := newClient()
	k8s := c.MustGet("client").(*models.K8s)
	clientPod := newClientPod(k8s)
	names := clientPod.List()

	c.IndentedJSON(http.StatusOK, names)
}

func CountPods(c *gin.Context) {
	k8s := c.MustGet("client").(*models.K8s)
	clientPod := newClientPod(k8s)
	numPods := clientPod.Count()

	c.IndentedJSON(http.StatusOK, numPods)
}

func newClientPod(k8s *models.K8s) *domain.ClientPod {
	return &domain.ClientPod{
		Clientset: k8s.Clientset,
		Pod:       k8s.Clientset.CoreV1().Pods("dma"),
	}
}
