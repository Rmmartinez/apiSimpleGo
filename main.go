package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string `json: "id""`
	Title  string `json="title"`
	Artist string `json="artist"`
	Year   int    `json="year"`
}

// esto podría ser información de una BD
var albums = []album{
	{ID: "1", Title: "Cupido", Artist: "Tini", Year: 2023},
	{ID: "2", Title: "Fuerza Natural", Artist: "Gustavo Cerati", Year: 2009},
	{ID: "3", Title: "Canción Animal", Artist: "Soda Stereo", Year: 1990},
	{ID: "4", Title: "Ruli", Artist: "El Kuelgue", Year: 2013},
	{ID: "5", Title: "+", Artist: "Ed Sheeran", Year: 2011},
}

// handler-capturamos la petición del cliente con *gin.Context
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// handler-se agrega el álbum a la lista de álbumes
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

// handler-recibe id y lo retorna si lo encontró, sino, not found
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no existe"})
}

// handler-recibe id y lo elimina si lo encontró, sino, not found
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album eliminado"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no existe"})
}

func main() {
	router := gin.Default()
	//hacemos una petición con el método get
	//localhost:8080/albums
	router.GET("/albums", getAlbums)

	//hacemos una petición con el método post
	//localhost:8080/albums y en el body, el json
	router.POST("/albums", postAlbums)

	//hacemos una petición con el método get
	//localhost:8080/albums/1
	router.GET("/albums/:id", getAlbumByID)

	//hacemos una petición con el método delete
	//localhost:8080/albums/1
	router.DELETE("/albums/:id", deleteAlbum)

	//ejecutamos un servidor
	router.Run("localhost:8080")
}
