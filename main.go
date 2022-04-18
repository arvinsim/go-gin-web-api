package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-web-api/models"
	"net/http"
	"runtime"
)

func main() {
	models.ConnectDatabase()

	// Setup Router
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(http.StatusOK, runtime.GOOS)
	})

	router.GET("/echo/:val", func(c *gin.Context) {
		val := c.Param("val")
		c.String(http.StatusOK, val)
	})

	// Get all items
	router.GET("/items/all", func(c *gin.Context) {
		var items []models.Item
		result := models.DB.Find(&items)

		httpCode := http.StatusOK
		if result.Error != nil {
			httpCode = http.StatusBadRequest
		}

		c.JSON(httpCode, gin.H{
			"data":         items,
			"rowsAffected": result.RowsAffected,
		})
	})

	// Add item
	router.POST("/items/add", func(c *gin.Context) {
		name := c.PostForm("name")
		item := models.Item{Name: name}
		result := models.DB.Select("Name").Create(&item)

		httpCode := http.StatusOK
		message := fmt.Sprintf("The item %s was created", name)

		if result.Error != nil {
			httpCode = http.StatusBadRequest
			message = fmt.Sprintf("There was an error creating item %s", name)
		}

		c.String(httpCode, message)
	})

	router.Run(":8000")
}
