package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Struct basico
type album struct {
	ID string `json:"id"`
	Title string `json:"titulo"`
	Artist string `json:"artist"`
	Year  int `json:"year"`
}

//Pequeña base de datos
var slnums =[]album{
	{ID:"1", Title:"Familia", Artist:"Camila CAbello", Year: 2022},
	{ID:"2", Title:"21", Artist:"Adele", Year: 2011},
	{ID:"3", Title:"The Eminem show", Artist:"Eminem", Year: 2022},
}
// Se utiliza para que el cliente de respuesta
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	fmt.Print("Hello World")
	routter := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}