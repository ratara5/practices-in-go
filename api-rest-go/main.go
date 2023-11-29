/*
API REST En Go CON Gin
https://www.youtube.com/watch?v=1EN0DHJTJZY
*/
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//CREAR ESTRUCTURA
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Year int `json:"year"`
}

//CREAR DATOS (NORMALMENTE VIENEN DESDE UNA BASE DE DATOS): Crear slicing
var albums =[]album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
	{ID: "3", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
}

//CREAR HANDLERS (MANEJADORES) DE PETICIONES PARA MOSTRAR DATOS O GUARDAR DATOS ANTE LA PETICIÓN DEL CLIENTE
func getAlbums(c *gin.Context){ //función que recibe como parámetro la petición del cliente
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")
	
	for _, a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album not found"})
}

func main() {
	//CREAR RUTAS PARA OBTENER LA PETICIÓN DEL CLIENTE
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run(":8080")
}
