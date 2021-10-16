package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcasanella/k8s_dashboard/models"
)

var albums []models.Album

func init() {
	albums = append(albums, models.Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	albums = append(albums, models.Album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	albums = append(albums, models.Album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99})
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
