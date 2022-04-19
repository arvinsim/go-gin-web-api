package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-web-api/models"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	// Setup Router
	router := gin.Default()
	router.GET("/*hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
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

		// Get specific item
		itemsGroup.GET("/:id", func(c *gin.Context) {
			var item models.Item
			id := c.Param("id")

			result := models.DB.Model(&models.Item{}).Where("id = ?", id).Take(&item)

			httpCode := http.StatusOK
			if result.Error != nil {
				httpCode = http.StatusBadRequest
				message := fmt.Sprintf("There was an error retrieving item with id %s", id)
				c.String(httpCode, message)
				return
			}

			c.JSON(httpCode, gin.H{
				"data": item,
			})
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

		// Delete item
		itemsGroup.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")

			result := models.DB.Delete(&models.Item{}, id)

			httpCode := http.StatusOK
			message := fmt.Sprintf("The item with id %s was deleted", id)
			if result.Error != nil {
				httpCode = http.StatusBadRequest
				message = fmt.Sprintf("There was an error deleting item with id %s", id)
			}

			c.String(httpCode, message)
		})
	}

	router.Run(":8000")
}
