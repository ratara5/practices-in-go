/*
CRUD API Rest con Go (GIN+GORN) y Mysql
https://www.youtube.com/watch?v=9b7KDUmnttM
*/
package main

import (
	"net/http"

	"raynet/go-backend/configs"
	"raynet/go-backend/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
