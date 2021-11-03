package repsitory

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/configcontext"
	"github.com/jcasanella/k8s_dashboard/models"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type PodObject struct {
	Pod v1.PodInterface
}

var podObject *PodObject

func init() {
	var err error
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	podObject = &PodObject{Pod: clientSet.CoreV1().Pods("dma")}
}

func ListPods(c *gin.Context) {
	pods, err := podObject.Pod.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	var names []models.Pod
	for _, pod := range pods.Items {
		names = append(names, models.Pod{Name: pod.Name})
	}

	c.IndentedJSON(http.StatusOK, names)
}

func CountPods(c *gin.Context) {
	pods, err := podObject.Pod.List(context.TODO(), meta.ListOptions{})
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
