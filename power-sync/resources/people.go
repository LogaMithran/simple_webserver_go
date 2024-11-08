package resources

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-sync/connectors"
	"power-sync/entities"
)

func GetStarWarsPeople(c *gin.Context) {
	var people []entities.People

	results := connectors.DB.Scopes(connectors.Paginate(c)).Find(&people)
	count := results.RowsAffected

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  people,
		"count": count,
	})
}

func GetStarWarsPeopleById(c *gin.Context) {
	var people *entities.People

	result := connectors.DB.Where("name = ?", c.Param("name")).First(&people)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"people": people,
	})
}

func CreatePeople(c *gin.Context) {
	var people entities.People
	if err := c.BindJSON(&people); err != nil {
		fmt.Printf("Error in binding the JSON %V", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	if result := connectors.DB.Create(&people).Error; result != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Cannot create people",
			"error":   result.Error(),
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   people,
	})
}

func UpdatePeople(c *gin.Context) {
	var people entities.People

	if err := c.BindJSON(&people); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	if response := connectors.DB.Save(&people); response != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   response.Error,
			"Message": "Cannot update people",
		})
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

func DeletePeople(c *gin.Context) {

	if response := connectors.DB.Where("name = ?", c.Param("name")).Delete(&entities.People{}); response.Error != nil {
		fmt.Println(response.Error.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": response.Error.Error(),
		})
	}
	c.JSON(http.StatusNoContent, "")
}
