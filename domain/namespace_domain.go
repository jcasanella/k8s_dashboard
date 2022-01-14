package domain

import (
	"context"
	"log"
	"time"

	"github.com/jcasanella/k8s_dashboard/models"
	v1core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type ClientNamespace struct {
	Clientset kubernetes.Interface
	Namespace v1.NamespaceInterface
}

type NamespaceOpers interface {
	List() ([]models.Namespace, error)
	Count() (int, error)
	Create(ns *v1core.Namespace) (*v1core.Namespace, error)
}

func (c ClientNamespace) List() ([]models.Namespace, error) {
	ns, err := c.Namespace.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		log.Panicf("Error listing namespaces %s", err.Error())
		return nil, err
	}

	var mns []models.Namespace
	for _, x := range ns.Items {
		mns = append(mns, models.Namespace{
			Name:   x.Name,
			Status: string(x.Status.Phase),
			Age:    x.CreationTimestamp.Format(time.RFC3339)})
	}

	return mns, nil
}

func (c ClientNamespace) Count() (int, error) {
	ns, err := c.Namespace.List(context.TODO(), meta.ListOptions{})
	if err != nil {
		log.Panicf("Error counting namespace %s", err.Error())
		return -1, err
	}

	return len(ns.Items), nil
}

func (c ClientNamespace) Create(namespace *v1core.Namespace) (*v1core.Namespace, error) {
	ns, err := c.Namespace.Create(context.TODO(), namespace, meta.CreateOptions{})
	if err != nil {
		log.Panicf("Error creating namespace %s:%s", ns.Name, err.Error())
		return nil, err
	}

	return ns, nil
}
