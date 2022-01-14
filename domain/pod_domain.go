package domain

import (
	"context"
	"log"

	"github.com/jcasanella/k8s_dashboard/models"
	v1core "k8s.io/api/core/v1"
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
	Create(pod *v1core.Pod) (*v1core.Pod, error)
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

func (c ClientPod) Create(pod *v1core.Pod) (*v1core.Pod, error) {
	p, err := c.Pod.Create(context.TODO(), pod, meta.CreateOptions{})
	if err != nil {
		log.Panicf("Error creating pod %s:%s", pod.Name, err.Error())
		return nil, err
	}

	return p, nil
}
