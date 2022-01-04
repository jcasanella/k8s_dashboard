package domain

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

	for _, expected := range podsToTest {
		t.Run(expected.Name, func(t *testing.T) {
			if _, err := client.Create(expected); err != nil {
				t.Fatalf("Error creating Pod %s", expected.Name)
			}

			actual := client.List()
			numExpected++
			if len(actual) != numExpected {
				t.Fatalf("The expected number of pods is %d numExpected, the actual number is %d", numExpected, len(actual))
			}

			if isPodNameInSlice(actual, expected.Name) != true {
				t.Fatalf("The expected pod name %s but does not exist", expected.Name)
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

	for _, expected := range podsToTest {
		t.Run(expected.Name, func(t *testing.T) {
			if _, err := client.Create(expected); err != nil {
				t.Fatalf("Error creating Pod %s", expected.Name)
			}

			actual := client.Count()
			numExpected++
			if actual != numExpected {
				t.Fatalf("The expected number of pods is %d numExpected, the actual number is %d", numExpected, actual)
			}
		})
	}
}

func TestCreatePods(t *testing.T) {
	client := newTestClientPods()

	for _, expected := range podsToTest {
		t.Run(expected.Name, func(t *testing.T) {
			actual, err := client.Create(expected)
			if err != nil {
				t.Fatalf("Error creating pod: %s", expected.Name)
			}

			if expected != actual {
				t.Fatal("Pod created is different from expected")
			}
		})
	}
}
