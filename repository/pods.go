package repository

import (
	"context"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/configcontext"
	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	Clientset kubernetes.Interface
	Pod       v1.PodInterface
}

type Pod interface {
	List() []models.Pod
	Count() int
	Create()
}

func (c Client) List() []models.Pod {
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

func (c Client) Count() int {
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
	client := c.MustGet("client").(*Client)
	names := client.List()

	c.IndentedJSON(http.StatusOK, names)
}

func CountPods(c *gin.Context) {
	client := c.MustGet("client").(*Client)
	numPods := client.Count()

	c.IndentedJSON(http.StatusOK, numPods)
}

func getK8sClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *configcontext.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}
