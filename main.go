package main

import (
	"net/http"

	"github.com/miftahulhidayati/rest-api-gin/controllers"
	"github.com/miftahulhidayati/rest-api-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.InitMysqlDB()

	DBConn := &controllers.DBConn{DB: db}

	router.GET("/hello", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"code" : http.StatusOK,
		})
	})
	router.GET("/person", DBConn.GetPersons)
	router.GET("/person/:id", DBConn.GetPerson)
	router.POST("/person", DBConn.CreatePerson)
	router.PUT("/person", DBConn.UpdatePerson)
	router.DELETE("/person/:id", DBConn.DeletePerson)

	router.Run(":8080")
}