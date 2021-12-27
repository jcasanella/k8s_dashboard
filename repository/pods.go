package repository

import (
	"context"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ClientPod struct {
	Clientset kubernetes.Interface
	Pod       v1.PodInterface
}

type PodOpers interface {
	List() []models.Pod
	Count() int
	Create()
}

func (c ClientPod) List() []models.Pod {
	pods, err := c.Pod.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.Pod
	for _, pod := range pods.Items {
		names = append(names, models.Pod{Name: pod.Name})
	}

	return names
}

func (c ClientPod) Count() int {
	pods, err := c.Pod.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return len(pods.Items)
}

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
	clientset := c.MustGet("client").(*kubernetes.Clientset)
	client := newClientPod(clientset)
	names := client.List()

	c.IndentedJSON(http.StatusOK, names)
}

func CountPods(c *gin.Context) {
	clientset := c.MustGet("client").(*kubernetes.Clientset)
	client := newClientPod(clientset)
	numPods := client.Count()

	c.IndentedJSON(http.StatusOK, numPods)
}

func newClientPod(clientset *kubernetes.Clientset) *ClientPod {
	return &ClientPod{
		Clientset: clientset,
		Pod:       clientset.CoreV1().Pods("dma"),
	}
}
