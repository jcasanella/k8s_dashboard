package repository

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func newTestClientPods() *ClientPod {
	cp := ClientPod{}
	cp.Clientset = fake.NewSimpleClientset()
	cp.Pod = cp.Clientset.CoreV1().Pods("dma")

	return &cp
}

func buildPodDefinition(namespace string, podName string, imageName string, image string) *v1.Pod {
	pod := &v1.Pod{
		TypeMeta: meta.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: meta.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            imageName,
					Image:           image,
					ImagePullPolicy: "Always",
				},
			},
		},
	}

	return pod
}

func TestListPods(t *testing.T) {
	client := newTestClientPods()
	pods := client.List()
	if len(pods) != 0 {
		t.Fatal("Should not exist any pod")
	}

	pod := buildPodDefinition("dma", "pod1", "nginx", "nginx")
	if _, err := client.Create(pod); err != nil {
		t.Fatal("Error creating Pod")
	}

	pods = client.List()
	if len(pods) != 1 {
		t.Fatal("Should not exist any pod")
	}
}
