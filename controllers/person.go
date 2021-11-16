package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/miftahulhidayati/rest-api-gin/models"
)


func (conn *DBConn) CreatePerson(c *gin.Context) {
	var person models.Person
	var result gin.H

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	person.FirstName = firstName
	person.LastName = lastName

	conn.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

func (idb *DBConn) GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *DBConn) GetPersons(c *gin.Context) {
	var (
		persons []models.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *DBConn) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")
	var (
		person    models.Person
		newPerson models.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newPerson.FirstName = firstName
	newPerson.LastName = lastName
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *DBConn) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
