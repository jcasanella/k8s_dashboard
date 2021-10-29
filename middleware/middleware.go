package middleware

import (
	"context"
	"flag"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var albums []models.Album

// var kubeconfig *string
var clientSet *kubernetes.Clientset

func init() {
	albums = append(albums, models.Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	albums = append(albums, models.Album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	albums = append(albums, models.Album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99})

	var err error
	clientSet, err = getK8sClient()
	if err != nil {
		panic(err.Error())
	}
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "album not found"})
}

func GetTitle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"albums": albums})
}

func CountPods(c *gin.Context) {
	pods, err := clientSet.CoreV1().Pods("dma").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, len(pods.Items))
}

func getK8sClient() (*kubernetes.Clientset, error) {
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

	return kubernetes.NewForConfig(config)
}
