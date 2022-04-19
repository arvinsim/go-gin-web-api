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

	itemsGroup := router.Group("/items")
	{
		// Get all items
		itemsGroup.GET("/", func(c *gin.Context) {
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

		// TODO: Get specific item
		itemsGroup.GET("/:id", func(context *gin.Context) {
		})

		// Add item
		itemsGroup.POST("/", func(c *gin.Context) {
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

		// Update item
		itemsGroup.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			name := c.PostForm("name")

			result := models.DB.Model(&models.Item{}).Where("id = ?", id).Update("name", name)

			httpCode := http.StatusOK
			message := fmt.Sprintf("The item %s with id %s was updated", name, id)
			if result.Error != nil {
				httpCode = http.StatusBadRequest
				message = fmt.Sprintf("There was an error saving item %s with id %s", name, id)
			}

			c.String(httpCode, message)
		})

		// TODO: Delete item
		itemsGroup.DELETE("/:id", func(context *gin.Context) {})
	}

	router.Run(":8000")
}
