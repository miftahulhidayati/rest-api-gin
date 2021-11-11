package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miftahulhidayati/rest-api-gin/database"
)

func main() {
	router := gin.Default()

	db := database.InitMysqlDB()

	router.GET("/hello", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"code" : http.StatusOK,
		})
	})
	router.POST("/hello")

	router.Run(":8080")
}