package main

import (
	"log"

	"github.com/jcasanella/k8s_dashboard/config"
	"github.com/jcasanella/k8s_dashboard/models"
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
	config, err := clientcmd.BuildConfigFromFlags("", *config.Kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return kubernetes.NewForConfig(config)
}

func newClient() *models.K8s {
	clientSet, err := getK8sClient()
	if err != nil {
		panic(err.Error())
	}

	return &models.K8s{
		Clientset: clientSet,
	}
}
