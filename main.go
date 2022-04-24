package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-web-api/models"
	"log"
	"net/http"
	"os"
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

	// Returns what the build environment is
	router.GET("/environment", func(c *gin.Context) {
		environment := os.Getenv("BUILD_ENVIRONMENT")
		if environment == "" {
			c.String(http.StatusOK, "The build environment is not set")
		} else {
			c.String(http.StatusOK, fmt.Sprintf("The build environment is set to %s", environment))
		}
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Upload file
	router.POST("/upload", func(c *gin.Context) {
		// Single file
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}
		log.Println(file.Filename)

		// Upload the file to specific dst.
		result := c.SaveUploadedFile(file, "./uploaded_files/")
		log.Println(result)
		if result.Error != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("There was a problem trying to upload '%s'", file.Filename))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Items endpoints
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
