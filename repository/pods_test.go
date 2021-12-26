package repository

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// import (
// 	v1 "k8s.io/api/core/v1"
// 	"k8s.io/client-go/kubernetes/fake"
// )

// func createDummyPods() {
// 	fakeKubeClient := fake.NewSimpleClientset()
// 	pod := &v1.Pod{}

// 	fakeKubeClient.
// }

func TestListPods(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ListPods(c)
}
