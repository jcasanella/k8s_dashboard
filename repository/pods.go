package repository

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/configcontext"
	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type podObject struct {
	Pod v1.PodInterface
}

func newPod() *podObject {
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	return &podObject{Pod: clientSet.CoreV1().Pods("dma")}
}

func ListPods(c *gin.Context) {
	limitQuery := c.DefaultQuery("limit", "10")
	limitValue, err := strconv.ParseInt(limitQuery, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	listOptions := &meta.ListOptions{
		Limit: limitValue,
	}

	pods, err := newPod().Pod.List(context.TODO(), *listOptions)
	if err != nil {
		panic(err.Error())
	}

	var names []models.Pod
	for _, pod := range pods.Items {
		names = append(names, models.Pod{Name: pod.Name, Continue: pods.ListMeta.Continue, RemainingItemCount: *pods.ListMeta.RemainingItemCount})
	}

	c.IndentedJSON(http.StatusOK, names)
}

func CountPods(c *gin.Context) {
	pods, err := newPod().Pod.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, len(pods.Items))
}

func getK8sClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *configcontext.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}
