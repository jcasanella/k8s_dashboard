package main

import (
	"log"

	"github.com/jcasanella/k8s_dashboard/configcontext"
	"github.com/jcasanella/k8s_dashboard/repository"
	"github.com/jcasanella/k8s_dashboard/router"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client := newClient()
	router := router.Router(client)

	error := router.Run("localhost:8085")
	log.Fatal(error)
}

func getK8sClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *configcontext.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

func newClient() *repository.Client {
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	return &repository.Client{
		Clientset: clientSet,
		Pod:       clientSet.CoreV1().Pods("dma"),
	}
}
