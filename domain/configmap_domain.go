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

type ClientConfigMap struct {
	Clientset kubernetes.Interface
	ConfigMap v1.ConfigMapInterface
}

type ConfigMapOpers interface {
	List() ([]models.ConfigMap, error)
	Count() (int, error)
	Create(configmap *v1core.ConfigMap) (*v1core.ConfigMap, error)
}

func (c ClientConfigMap) Count() (int, error) {
	lcm, err := c.ConfigMap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		log.Panicf("Error conting configmaps %s", err.Error())
		return -1, err
	}

	return len(lcm.Items), nil
}

func (c ClientConfigMap) List() ([]models.ConfigMap, error) {
	lcm, err := c.ConfigMap.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		log.Panicf("Error listing configmaps %s", err.Error())
		return nil, err
	}

	var cm []models.ConfigMap
	for _, x := range lcm.Items {
		cm = append(cm, models.ConfigMap{Name: x.Name})
	}

	return cm, nil
}

func (c ClientConfigMap) Create(configmap *v1core.ConfigMap) (*v1core.ConfigMap, error) {
	cm, err := c.ConfigMap.Create(context.TODO(), configmap, meta.CreateOptions{})
	if err != nil {
		log.Panicf("Error cresting configmap %s:%s", cm.Name, err.Error())
		return nil, err
	}

	return cm, nil
}
