package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jcasanella/k8s_dashboard/router"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	// Check kubernetes configs
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, _ := c.CoreV1().Pods("dma").List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, v := range pods.Items {
		fmt.Printf("%s\n", v.Name)
	}

	cm, _ := c.CoreV1().ConfigMaps("dma").List(context.TODO(), v1.ListOptions{})
	fmt.Printf("There are %d cm in the cluster\n", len(cm.Items))

	for _, v := range cm.Items {
		fmt.Printf("%s\n", v.Name)
	}

	router := router.Router()

	error := router.Run("localhost:8085")
	log.Fatal(error)
}
