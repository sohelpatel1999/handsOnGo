package controller

import (
	"context"
	"fmt"
	"handsOnGO/dto"
	"handsOnGO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GerPersonDetails(c *gin.Context) {
	fmt.Println("sohel")
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
