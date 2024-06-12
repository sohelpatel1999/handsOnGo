package controller

import (
	"context"
	"fmt"
	"handsOnGO/dto"
	"handsOnGO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlllPersonDetails(c *gin.Context) {
	persons, err := service.GetAllPersons(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, persons)
}

func CreatePerson(c *gin.Context) {
	var person dto.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print(person)

	personresponse, _, err := service.CreatePerson(context.Background(), person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personresponse)

}

func GetPersonDetailsById(c *gin.Context) {
	ID := c.Query("id")
	fmt.Println(ID)

	personresponse, err := service.GetPersonDetailsById(context.Background(), ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personresponse)

}

func UpdatePersonDetailsById(c *gin.Context) {
	ID := c.Query("id")
	if ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}
	fmt.Println("Received ID:", ID)
	var person dto.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personresponse, err := service.UpdatePersonDetailsById(context.Background(),person, ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personresponse)

}
