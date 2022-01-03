package repository

import (
	"testing"

	"github.com/jcasanella/k8s_dashboard/models"
	v1 "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

var podsToTest = []*v1.Pod{
	buildPodDefinition("dma", "nginx1", "nginx", "nexus/nginx"),
	buildPodDefinition("dma", "nginx2", "nginx", "nexus/nginx"),
	buildPodDefinition("dma", "mongodb1", "mongodb", "nexus/mongodb"),
	buildPodDefinition("dma", "mongodb2", "mongodb", "nexus/mongodb"),
	buildPodDefinition("dma", "cassandra1", "cassandra", "nexus/cassandra"),
}

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

func isPodNameInSlice(pods []models.Pod, name string) bool {
	for _, elem := range pods {
		if elem.Name == name {
			return true
		}
	}

	return false
}

func TestListPods(t *testing.T) {
	numExpected := 0
	client := newTestClientPods()
	pods := client.List()
	if len(pods) != numExpected {
		t.Fatal("Should not exist any pod")
	}

	for _, pod := range podsToTest {
		t.Run(pod.Name, func(t *testing.T) {
			if _, err := client.Create(pod); err != nil {
				t.Fatal("Error creating Pod")
			}

			listPods := client.List()
			numExpected++
			if len(listPods) != numExpected {
				t.Fatalf("The expected number of pods is %d numExpected, the actual number is %d", numExpected, len(listPods))
			}

			if isPodNameInSlice(listPods, pod.Name) != true {
				t.Fatalf("The expected pod name %s but does not exist", pod.Name)
			}
		})
	}
}

func TestCountPods(t *testing.T) {
	numExpected := 0
	client := newTestClientPods()
	pods := client.List()
	if len(pods) != numExpected {
		t.Fatal("Should not exist any pod")
	}

	for _, pod := range podsToTest {
		t.Run(pod.Name, func(t *testing.T) {
			if _, err := client.Create(pod); err != nil {
				t.Fatal("Error creating Pod")
			}

			numPods := client.Count()
			numExpected++
			if numPods != numExpected {
				t.Fatalf("The expected number of pods is %d numExpected, the actual number is %d", numExpected, numPods)
			}
		})
	}
}
