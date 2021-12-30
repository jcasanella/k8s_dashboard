package models

import "k8s.io/client-go/kubernetes"

type K8s struct {
	Clientset kubernetes.Interface
}
